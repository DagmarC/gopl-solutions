// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 139.

// Findlinks3 crawls the web, starting with the URLs on the command line.
package main

import (
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"

	"github.com/DagmarC/gopl-solutions/utils"

	"github.com/DagmarC/gopl-solutions/utils/html"
)

func breadthFirst(f func(item, host string) []string, worklist []string, host string) {

	seen := make(map[string]bool)

	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item, host)...)
			}
		}
	}
}

func crawl(path, host string) []string {

	url, err := html.GetURL(path)
	if err != nil {
		log.Fatal(err)
	}

	savePage := func() {

		fpath, dirPath := getFileDirPath(url, host)

		os.MkdirAll(dirPath, os.ModePerm)

		f, err := os.Create(fpath)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		resp, err := http.Get(path)
		if err != nil {
			log.Fatal(err)
		}
		io.Copy(f, resp.Body)
	}

	if url.Host == host {
		savePage() // Save only pages that come from the host page.
	}

	list, err := Extract(path)
	if err != nil {
		log.Print(err)
	}
	return list
}

func getFileDirPath(url *url.URL, host string) (string, string) {
	wd, err := utils.GetCurrentWD()
	if err != nil {
		log.Fatal(err)
	}

	var fpath string
	if filepath.Ext(url.Path) == "" {
		fpath = filepath.Join(wd, host, url.Path, "index.html")
	} else {
		fpath = filepath.Join(wd, host, url.Path)
	}
	dirPath := filepath.Dir(fpath)

	return fpath, dirPath
}

func main() {
	// Crawl the web breadth-first,
	// starting from the command-line arguments.
	breadthFirst(crawl, os.Args[1:], html.GetHost(os.Args[1]))
}
