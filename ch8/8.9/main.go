// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 250.

// The du3 command computes the disk usage of the files in a directory.
package main

// The du3 variant traverses all directories in parallel.
// It uses a concurrency-limiting counting semaphore
// to avoid opening too many files at once.

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

// go run . -v /Users /Library /opt
var vFlag = flag.Bool("v", false, "show verbose progress messages")

// Info about the file that will be coming through the chan
type file struct {
	root string
	name string
	size int64 // size of file in bytes
}

type rootInfo struct {
	name  string
	size  int64 // total size of files in bytes
	count int64 // total number of files
}

type rootsMap map[string]*rootInfo // key is the name of the root folder, for better indexing

//!+
func main() {
	//!-
	flag.Parse()

	// Determine the initial directories.
	rootsArgs := flag.Args()
	if len(rootsArgs) == 0 {
		rootsArgs = []string{"."}
	}

	fileChan := make(chan file, len(rootsArgs))
	roots := make(rootsMap, len(rootsArgs))

	//!+
	// Traverse each root of the file tree in parallel.

	var n sync.WaitGroup
	for _, root := range rootsArgs {
		roots[root] = &rootInfo{name: root} // init map of each root file
		n.Add(1)
		go walkDir(root, &n, fileChan, root)
	}
	go func() {
		n.Wait()
		close(fileChan)
	}()
	//!-

	// Print the results periodically.
	var tick <-chan time.Time
	if *vFlag {
		tick = time.Tick(500 * time.Millisecond)
	}

loop:
	for {
		select {
		case f, ok := <-fileChan:
			if !ok {
				break loop // fileChan was closed
			}
			if _, ok := roots[f.root]; ok { // check if root exists in map (should exist)
				roots[f.root].count++ // new file has arrived
				roots[f.root].size += f.size
			}
		case <-tick:
			printDiskUsage(roots)
		}
	}

	printDiskUsage(roots) // final totals
	//!+
}

//!-

func printDiskUsage(roots rootsMap) {
	var ttlCnt, ttlSize int64
	for r, i := range roots {
		ttlCnt += i.count
		ttlSize += i.size
		fmt.Printf("Root:\t%-8s %8d files\t%8.1f GB\n", r, i.count, float64(i.size)/1e9)
	}
	fmt.Printf("Total:\t\t%9d files\t%8.1f GB\n--------------------------------------------\n", ttlCnt, float64(ttlSize)/1e9)

}

// walkDir recursively walks the file tree rooted at dir
// and sends the fileInfo (name, root and size) of each found file on fileChan.
//!+walkDir
func walkDir(dir string, n *sync.WaitGroup, fileChan chan<- file, root string) {
	defer n.Done()
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(subdir, n, fileChan, root)
		} else {
			fileChan <- file{root: root, name: entry.Name(), size: entry.Size()}
		}
	}

}

//!-walkDir

//!+sema
// sema is a counting semaphore for limiting concurrency in dirents.
var sema = make(chan struct{}, 20)

// dirents returns the entries of directory dir.
func dirents(dir string) []os.FileInfo {
	sema <- struct{}{}        // acquire token
	defer func() { <-sema }() // release token
	// ...
	//!-sema

	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du: %v\n", err)
		return nil
	}
	return entries
}
