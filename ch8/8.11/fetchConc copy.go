// Fetchall fetches URLs in parallel and reports their times and sizes.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func cancelled(done <-chan struct{}) bool {
	select {
	case <-done:
		return true
	default:
	}
	return false
}

// FetchCpy shows how to fetch  multiple URLs of uknown length concurrently and when the first response arrives, cancel the rest of them
// done channel is used only when I dont know the exact number of ongoing adresses to be fetched, however here I know it, since len(os.Args[1:]) os exact.
func mainCpy() {
	if len(os.Args[1:]) == 0 {
		return
	}
	fmt.Println("Number of addresses to be fetched: ", len(os.Args[1:]))

	ch := make(chan string, len(os.Args[1:])) // bufferred chan to prevent the goroutine leak
	done := make(chan struct{})               // to broadcast and stop gors

	for _, url := range os.Args[1:] {
		go fetchCpy(url, ch, done) // start a goroutine
	}

	fmt.Println(<-ch) // receive the 1st response from channel ch, rest of responses are being discarded
	close(done)       // broadcast the end to all active gors
	panic("DEBUG")    // debug if all gors has finished
}

func fetchCpy(url string, ch chan<- string, done <-chan struct{}) {
	if cancelled(done) {
		fmt.Println("End fetching url before GET: ", url)
		return
	}
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}

	if cancelled(done) {
		fmt.Println("End fetching url before COPY resp: ", url)
		return
	}

	nbytes, err := io.Copy(io.Discard, resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	select {
	case ch <- fmt.Sprintf("INFO: Response: %7d bytes url=%s", nbytes, url):
	case <-done:
		fmt.Println("End fetching url before sending to resp.body to channel ch.", url)
	}
}

//!-
