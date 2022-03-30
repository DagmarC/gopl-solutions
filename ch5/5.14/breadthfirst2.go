package main

// Ecample from https://github.com/ray-g/gopl/blob/8cd4330890081305af7eff2ac09f1a821f99c9de/ch05/ex5.14/bsd.go

import (
	"fmt"
	"os"
)

func breadthFirstOrig(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

func getPrereqs(course string) (order []string) {
	fmt.Fprintf(os.Stdout, "%s\n", course)
	order = append(order, prereqs[course]...)
	return
}