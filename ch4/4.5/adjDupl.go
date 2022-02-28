package main

import (
	"fmt"
)

func main() {
	t := []string{"a", "a", "a", "b", "c", "c"}
	t = removeAdjDuplicates(t)
	fmt.Println(t)
}

// removeDuplicates eliminates adjacent duplicates in-place.
func removeAdjDuplicates(s []string) []string {
	for i := 0; i < len(s); i++ {

		var j = i + 1
		for ; j < len(s) && s[i] == s[j]; j++ {
		}

		if j == i+1 {
			continue
		}

		if j == len(s) {
			s = s[:i+1]
			break
		}

		copy(s[i+1:], s[j:])
		s = s[:len(s)-(j-i)+1]
	}
	return s
}
