// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package memo_test

import (
	"testing"

	"github.com/DagmarC/gopl-solutions/ch9/9.3"
	"github.com/DagmarC/gopl-solutions/ch9/9.3/memotest"
)

var httpGetBody = memotest.HTTPGetBody

func Test(t *testing.T) {
	m := memo.New(httpGetBody)
	memotest.Sequential(t, m)
}

func TestConcurrent(t *testing.T) {
	m := memo.New(httpGetBody)
	memotest.Concurrent(t, m)
}

func TestCConcClose1(t *testing.T) {
	m := memo.New(httpGetBody)
	memotest.ConcurrentClose1(t, m) // closes done channel after some sleep
}

func TestSequentialClose(t *testing.T) {
	m := memo.New(httpGetBody)
	memotest.SequentialClose(t, m) // closes done channel after some sleep
}