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

	"github.com/DagmarC/gopl-solutions/ch8/8.10/links"
)

// CANCEL EVENT STEPS:
// 1. Utility function cancelled
// 2. done channel and close done channel in its own GOR (eg os.Stdin event listener and closer then)
// 3. make other GORS to respond to steps above (VIA select case <-done event OR cancelled() function)
// 4. To make sure (testing purposes) that all gors has finished, you can call panic
//    at the place of returning after done is closed to see the stacktrace
//    NOTE: panic: Help .. goroutine 1 [running]: main.main()

// cancelled is the Utility function that checks or polls the cancellation state
// at the instant it is called
func cancelled(done chan struct{}) bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}

//!+semaphore
// tokens is a counting semaphore used to
// enforce a limit of 20 concurrent requests.
var tokens = make(chan struct{}, 20)

// crawl the URL nodes reachable by at most maxDepth links, node.d represents the current depth node, if maxDepth=0 -> no link printed

func crawl(node URLnode, maxDepth int, done chan struct{}) []URLnode {
	if cancelled(done) {
		return []URLnode{}
	}

	if maxDepth == 0 {
		return []URLnode{} // no depth --> no URL to return
	}
	fmt.Println(node.url)
	if node.d >= maxDepth {
		return []URLnode{} // max depth reached, do not extraxt more URLs
	}

	tokens <- struct{}{} // acquire a token
	list, err := links.ExtractTimeout(node.url)
	<-tokens // release the token

	if err != nil {
		log.Printf("CUSTOM ERROR: Node %s, Depth %d, %v", node.url, node.d, err)
		go func() {
			if !cancelled(done) {
				close(done) //  Via closing done channel - Broadcast to receivers, that sending will not occur.
			}
		}()
		return []URLnode{}
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

	done := make(chan struct{}) // Signal the cancellation
	nodesChan := make(chan []URLnode)

	var n int // number of pending sends to worklist
	var listNodes []URLnode

	// Start with the command-line arguments.
	n++
	go func() {
		nodesChan <- createURLNodes(flag.Args(), 1) // Sending, root URL nodes at depth 1
	}()

	// Crawl the web concurrently.
	seen := make(map[string]bool)
	for ; n > 0; n-- {

		// To be sure signal has come <-done is done twice in a row.
		select {
		case <-done:
			fmt.Println("Done channel closed.")
			return
		default:
		}

		select {
		case listNodes = <-nodesChan: // Receiving
			for _, node := range listNodes {
				if !seen[node.url] {
					seen[node.url] = true
					n++
					go func(node URLnode) {
						nodesChan <- crawl(node, *depthPtr, done) // Sending side, if crawl ends with error, done signal will be received.
					}(node)
				}
			}
		case <-done:
			fmt.Println("Done channel closed..")
			return

		}

	}
	// panic("Help") // Debugging stacktrace for how many goroutines are running = only one main should be running
}

//!-
