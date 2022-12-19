// Fetchall fetches URLs in parallel and reports their times and sizes.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

// Fetch multiple URLs concurrently and when the first response arrives, cancel the rest of them - bufferred channel to prevent gors leak.
func main() {
	if len(os.Args[1:]) == 0 {
		return
	}
	fmt.Println("Number of addresses to be fetched: ", len(os.Args[1:]))
	ch := make(chan string, len(os.Args[1:])) // bufferred chan to prevent the goroutine leak
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // start a goroutine
	}

	fmt.Println(<-ch) // receive the 1st response from channel ch, rest of responses are being discarded
	// panic("DEBUG") // debug if all gors has finished
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}

	nbytes, err := io.Copy(io.Discard, resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}

//!-
