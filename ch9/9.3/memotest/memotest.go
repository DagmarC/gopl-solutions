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

/*
//!+seq
	m := memo.New(httpGetBody)
//!-seq
*/

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

/*
//!+conc
	m := memo.New(httpGetBody)
//!-conc
*/

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

func ConcurrentClose1(t *testing.T, m M) {
	//!+conc
	var n sync.WaitGroup
	var i int
	for url := range incomingURLs() {
		n.Add(1)
		go func(url string, i int) {
			done := make(chan struct{}, 1)

			go func(url string, i int) {
				rand.Seed(time.Now().UnixNano())
				time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
				// fmt.Printf("INFO: Closing done channel %s number %d.\n", url, i)
				close(done)
			}(url, i)

			// defer close(done)
			defer n.Done()
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
	n.Wait()
	//!-conc
}

func SequentialClose(t *testing.T, m M) {
	//!+seq
	var i int
	for url := range incomingURLs() {
		start := time.Now()
		done := make(chan struct{}, 1)

		go func() {
			rand.Seed(time.Now().UnixNano())
			time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
			// fmt.Println("\n----INFO: Closing done channel.")
			close(done)
		}()

		value, err := m.Get(url, done, i)
		if err != nil {
			log.Print(err)
			i++
			continue
		}
		fmt.Printf("\n========%s number %d, %s, %d bytes=========\n",
			url, i, time.Since(start), len(value.([]byte)))

		i++
	}
	//!-seq
}
