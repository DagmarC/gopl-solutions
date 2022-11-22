// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 241.

// Crawl2 crawls web links starting with the command-line arguments.
//
// This version uses a buffered channel as a counting semaphore
// to limit the number of concurrent calls to links.Extract.
package main

import (
	"log"
	"net/url"
	"os"
	"sync"

	"github.com/DagmarC/gopl-solutions/utils/html"
)

//!+sema
// tokens is a counting semaphore used to
// enforce a limit of 20 concurrent requests.
var tokens = make(chan struct{}, 10)

func crawl(path string, save bool) []string {

	tokens <- struct{}{} // acquire a token
	list, err := html.Extract(path)
	<-tokens // release the token

	if err != nil {
		log.Print(err)
	}
	return list
}

//!-sema

const dst string = "pages" // destination folder for html pages in root dir

//!+
func main() {
	var wg sync.WaitGroup // WG for saving page
	var n int             // number of pending sends to worklist

	worklist := make(chan []string)
	hosts := make([]string, 0, len(os.Args)-1)
	// Start with the command-line arguments.

	for _, arg := range os.Args[1:] {
		cmdURL, err := url.Parse(arg)
		if err != nil {
			log.Fatal(err)
		}
		hosts = append(hosts, cmdURL.Host)
	}

	n++
	go func() {
		worklist <- os.Args[1:]
	}()

	// Compare URL hosts of given cmd links and param link (links from crawler)
	hostLink := func(link string) bool {
		url, err := url.Parse(link)
		if err != nil {
			log.Fatal(err)
		}
		for _, host := range hosts {
			if url.Host == host {
				return true
			}
		}
		return false
	}

	// Crawl the web concurrently.
	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <-worklist
		for _, link := range list {
			if !seen[link] && hostLink(link) {
				seen[link] = true
				n++
				wg.Add(1)
				go func(link string) {
					worklist <- crawl(link, true)
					err := html.SavePage(link, dst)
					if err != nil {
						log.Fatal(err)
					}
					defer wg.Done()
				}(link)
			}
		}
	}
	wg.Wait() // wait for all gors of saving
}

//!-
