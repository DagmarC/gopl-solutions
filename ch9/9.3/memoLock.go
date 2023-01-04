// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 276.

// Package memo provides a concurrency-safe memoization a function of
// a function.  Requests for different keys proceed in parallel.
// Concurrent requests for the same key block until the first completes.
// This implementation uses a Mutex.
package memo

import (
	"errors"
	"fmt"
	"sync"
)

// Func is the type of the function to memoize.
type Func func(string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

//!+
type entry struct {
	res    result
	ready  chan struct{} // closed when res is ready
	cancel chan struct{}
}

func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]*entry)}
}

type Memo struct {
	f     Func
	mu    sync.Mutex // guards cache
	cache map[string]*entry
}

func (memo *Memo) Get(key string, done chan struct{}, i int) (value interface{}, err error) {
	// fmt.Println("\n===REQ===", key)
	if cancelled(done) {
		return nil, errors.New("INFO: request cancelled, before being even made")// Do not make any request at all
	}

	memo.mu.Lock()
	e := memo.cache[key]

	if e == nil {
		// This is the first request for this key.
		// This goroutine becomes responsible for computing
		// the value and broadcasting the ready condition.
		e = &entry{ready: make(chan struct{})}
		memo.cache[key] = e
		memo.mu.Unlock()
		e.res.value, e.res.err = memo.f(key)

		close(e.ready) // broadcast ready condition
	} else {
		// This is a repeat request for this key.
		memo.mu.Unlock()
		<-e.ready // wait for ready condition
	}
	// 9.3: If client cancel the request delete the key from cache
	if cancelled(done) {
		memo.mu.Lock()
		defer memo.mu.Unlock()

		delete(memo.cache, key)
		fmt.Printf("\n===KEY %s number %d==== deleted from cache %v \n", key,i, memo.cache)
		return nil, errors.New("INFO: request cancelled, after request being made")// Do not make any request at all
	}
	return e.res.value, e.res.err
}

func cancelled(done chan struct{}) bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}

//!-
