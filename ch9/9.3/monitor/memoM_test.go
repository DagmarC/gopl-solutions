// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package monitor_test

import (
	"testing"

	"github.com/DagmarC/gopl-solutions/ch9/9.3/memotest"
	"github.com/DagmarC/gopl-solutions/ch9/9.3/monitor"
)

var httpGetBody = memotest.HTTPGetBody

func Test(t *testing.T) {
	m := monitor.New(httpGetBody)
	defer m.Close() // monitor version
	memotest.Sequential(t, m)
}

func TestConcurrent(t *testing.T) {
	m := monitor.New(httpGetBody)
	defer m.Close() // monitor version
	memotest.Concurrent(t, m)
}

func TestSequentialCloseMM(t *testing.T) {
	m := monitor.New(httpGetBody)
	defer m.Close()                // monitor version
	memotest.SequentialClose(t, m) // closes done channel after some sleep
}
func TestConcurrentCloseM(t *testing.T) {
	m := monitor.New(httpGetBody)
	defer m.Close()                // monitor version
	memotest.ConcurrentClose(t, m) // closes done channel after some sleep
}
