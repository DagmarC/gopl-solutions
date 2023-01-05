// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 272.

// Package memotest provides common functions for
// testing various designs of the memo package.
package memotest

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"sync"
	"testing"
	"time"
)

//!+httpRequestBody
func httpGetBody(url string) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

//!-httpRequestBody

var HTTPGetBody = httpGetBody

func incomingURLs() <-chan string {
	ch := make(chan string)
	go func() {
		for _, url := range []string{
			"https://golang.org",
			"https://godoc.org",
			"https://play.golang.org",
			"https://golang.org",
			"http://gopl.io",
			"https://golang.org",
			"https://godoc.org",
			"https://play.golang.org",
			"http://gopl.io",
			"https://godoc.org",
		} {
			ch <- url
		}
		close(ch)
	}()
	return ch
}

type M interface {
	Get(key string, done chan struct{}, i int) (interface{}, error)
}

func Sequential(t *testing.T, m M) {
	//!+seq
	for url := range incomingURLs() {
		start := time.Now()
		value, err := m.Get(url, nil, 0)
		if err != nil {
			log.Print(err)
			continue
		}
		fmt.Printf("%s, %s, %d bytes\n",
			url, time.Since(start), len(value.([]byte)))
	}
	//!-seq
}

func Concurrent(t *testing.T, m M) {
	//!+conc
	var n sync.WaitGroup
	done := make(chan struct{}, 1)
	for url := range incomingURLs() {
		n.Add(1)
		go func(url string) {
			defer n.Done()
			start := time.Now()
			value, err := m.Get(url, done, 0)
			if err != nil {
				log.Print(err)
				return
			}
			fmt.Printf("%s, %s, %d bytes\n",
				url, time.Since(start), len(value.([]byte)))
		}(url)
	}
	n.Wait()
	//!-conc
}

////////////
func SequentialClose(t *testing.T, m M) {
	//!+seq
	var wg sync.WaitGroup // Wait for closing all done channel, to prevent gor leak

	var i int // request number being processed

	for url := range incomingURLs() {

		// Testing purposes: closes done channel, either it will close when the request is being processed - request will be cancelled or after - which changes nothing
		done := make(chan struct{}, 1)
		wg.Add(1)
		go func() {
			rand.Seed(time.Now().UnixNano())
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			close(done)
			wg.Done()
		}()

		start := time.Now()
		value, err := m.Get(url, done, i)
		if err != nil {
			log.Print(err)
			i++ // increment the request number even if the error has occured
			continue
		}
		fmt.Printf("========%s number %d, %s, %d bytes=========\n",
			url, i, time.Since(start), len(value.([]byte)))
		fmt.Println()
		i++
	}
	wg.Wait() // Wait until all done channels are closed.
	//!-seq
}

func ConcurrentClose(t *testing.T, m M) {
	//!+conc
	var wq sync.WaitGroup
	var i int
	for url := range incomingURLs() {
		wq.Add(2) // 1 gor for url and 1 gor for closing done channel
		go func(url string, i int) {
			defer wq.Done()

			done := make(chan struct{}, 1)
			go func(url string, i int) {
				rand.Seed(time.Now().UnixNano())
				time.Sleep(time.Duration(rand.Intn(3000)) * time.Millisecond)
				close(done)
				defer wq.Done()
			}(url, i)

			start := time.Now()
			value, err := m.Get(url, done, i)
			if err != nil {
				log.Print(err)
				return
			}
			fmt.Printf("\n=========%s number %d, %s, %d bytes=======\n",
				url, i, time.Since(start), len(value.([]byte)))
		}(url, i)
		i++
	}
	wq.Wait()
	//!-conc
}
