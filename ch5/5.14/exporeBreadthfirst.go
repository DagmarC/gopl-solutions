package main

import (
	"fmt"
	"os"
)

var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},

	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},

	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func breadthFirst(m map[string][]string) {

	seen := make(map[string]bool)
	worklist := make([]string, 0, 2*len(m))

	for key := range m {
		worklist = append(worklist, key)
	}
	fmt.Println("KEYS BEGIN", worklist)

	for len(worklist) > 0 {
		items := worklist
		worklist = nil

		for _, item := range items {
			if !seen[item] {
				fmt.Println("NOW SEEN", item)
				seen[item] = true
				worklist = append(worklist, m[item]...)
			}
		}
	}
}

func main() {
	breadthFirst(prereqs)

	// OR use original breadthFirstOrig with function as an argument:
	for corse, prereq := range prereqs {
		fmt.Fprintf(os.Stdout, "==========\n\"%s\" depends on following corses:\n", corse)
		breadthFirstOrig(getPrereqs, prereq)
		fmt.Fprintf(os.Stdout, "==========\n\n")
	}
}
