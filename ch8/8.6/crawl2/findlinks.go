// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 241.

// Crawl2 crawls web links starting with the command-line arguments.
//
// This version uses a buffered channel as a counting semaphore
// to limit the number of concurrent calls to links.Extract.
package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/DagmarC/gopl-solutions/utils/links"
)

//!+semaphore
// tokens is a counting semaphore used to
// enforce a limit of 20 concurrent requests.
var tokens = make(chan struct{}, 20)

// crawl the URL nodes reachable by at most maxDepth links, node.d represents the current depth node, if maxDepth=0 -> no link printed

func crawl(node URLnode, maxDepth int) []URLnode {
	if maxDepth == 0 {
		return []URLnode{} // no depth --> no URL to return
	}
	fmt.Println(node.url)
	if node.d >= maxDepth {
		return []URLnode{} // max depth reached, do not extraxt more URLs
	}

	tokens <- struct{}{} // acquire a token
	list, err := links.Extract(node.url)
	<-tokens // release the token

	if err != nil {
		log.Print(err)
	}

	return createURLNodes(list, node.d+1) // depth is increased by one from parent node
}

//!-semaphore

//!+URLnode structre
type URLnode struct {
	url string
	d   int
}

func newURLnode(url string, depth int) URLnode {
	return URLnode{url: url, d: depth}
}

func createURLNodes(links []string, d int) []URLnode {
	nodes := make([]URLnode, 0, len(links))
	for _, l := range links {
		nodes = append(nodes, newURLnode(l, d))
	}
	return nodes
}

//!-URLnode structre

//!+
var depthPtr = flag.Int("depth", 3, "depth: URLs reachable by at most n-depth links will be fetched")

func main() {
	flag.Parse()
	nodesChan := make(chan []URLnode)
	var n int // number of pending sends to worklist

	// Start with the command-line arguments.
	n++
	go func() {
		nodesChan <- createURLNodes(flag.Args(), 1) // root URL nodes at depth 1
	}()

	// Crawl the web concurrently.
	seen := make(map[string]bool)
	for ; n > 0; n-- {
		listNodes := <-nodesChan
		for _, node := range listNodes {
			if !seen[node.url] {
				seen[node.url] = true
				n++
				go func(node URLnode) {
					nodesChan <- crawl(node, *depthPtr)
				}(node)
			}
		}
	}
}

//!-
