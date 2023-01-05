// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 278.

// Package memo provides a concurrency-safe non-blocking memoization
// of a function.  Requests for different keys proceed in parallel.
// Concurrent requests for the same key block until the first completes.
// This implementation uses a monitor goroutine.
package monitor

import (
	"fmt"
)

//!+Func

// Func is the type of the function to memoize.
type Func func(key string) (interface{}, error) // e.g.  memotest.HTTPGetBody

// A result is the result of calling a Func.
type result struct {
	value interface{}
	err   error
}

type entry struct {
	res   result
	ready chan struct{} // closed when res is ready
}

//!-Func

//!+get

// A request is a message requesting that the Func be applied to key.
type request struct {
	key      string
	response chan<- result // the client wants a single result
	cancel   chan struct{} // done channel from client
	number   int           // used for better tracking/debugging
}

type Memo struct{ requests chan request }

// New returns a memoization of f.  Clients must subsequently call Close.
func New(f Func) *Memo {
	memo := &Memo{requests: make(chan request)}
	go memo.server(f)
	return memo
}

// Get 9.3 add an optional done channel, if the client has cancelled the operation response will be nil and error msg will be fulfilled
func (memo *Memo) Get(key string, done chan struct{}, i int) (interface{}, error) {
	response := make(chan result)

	memo.requests <- request{key, response, done, i}
	res := <-response

	return res.value, res.err
}

// Close: Client is responsible for closing the connection
func (memo *Memo) Close() {
	fmt.Println("INFO: Closing memo.requestes channel")
	close(memo.requests)
}

//!-get

//!+monitor gor

func (memo *Memo) server(f Func) {
	cache := make(map[string]*entry)
	deleteKey := make(chan string) // signal with key that will delete that key from cache

outer:
	for req := range memo.requests {
		// Skip the request asap, before the entry is added to a map
		if cancelled(req.cancel) {
			req.response <- result{nil, fmt.Errorf("request number %d cancelled before being processed", req.number)}
			continue outer // wait for the next request recievement
		}
		e := cache[req.key]
		if e == nil {
			// This is the first request for this key.
			e = &entry{ready: make(chan struct{})}
			cache[req.key] = e // If the req is cancelled during await of ready state, we simply delete the entry from map

			go e.call(f, req.key) // call f(key)
		}
		go e.deliver(req, deleteKey)
		<-removeKey(deleteKey, cache) // blocked until removeKey has decided whether to remove key from cache or just continue normally
	}

}

func removeKey(deleteS <-chan string, cache map[string]*entry) <-chan struct{} {
	done := make(chan struct{}) // output channel would block until recieve operation is performed on calling side (server) to prevent multiple deletes at the same time
	go func() {
		k := <-deleteS // Wait for the deleteSignal, "" means that nothing should to be deleted
		if k != "" {
			delete(cache, k)
			fmt.Printf("INFO: entry %s deleted, map=%v\n", k, cache)
		}
		done <- struct{}{} // operation successfull
	}()
	return done
}

func (e *entry) call(f Func, key string) {
	// Evaluate the function.
	e.res.value, e.res.err = f(key)
	// Broadcast the ready condition.
	close(e.ready)
}

// key string, response chan<- result, deleteS chan<- string, cancel <-chan struct{}, n int
func (e *entry) deliver(req request, deleteKey chan<- string) {

	// Wait for the ready condition (evalualtion of the memoization func from call()).
	<-e.ready

	// Cancel needs double check for sure for better timing.
	select {
	case <-req.cancel:
		deleteKey <- req.key
		req.response <- result{nil, fmt.Errorf("request %s number %d cancelled after response was cached 1", req.key, req.number)} // 9.3: Do not chache the result of a cancelled func
		return
	default: // non-blocking
	}
	// Send the result to the client or discard.
	select {
	case <-req.cancel:
		deleteKey <- req.key
		req.response <- result{nil, fmt.Errorf("request %s number %d cancelled after response was cached 2", req.key, req.number)}
	case req.response <- e.res:
		deleteKey <- "" // signal that nothing should be deleted
	}
}

//!-monitor
func cancelled(done chan struct{}) bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}
