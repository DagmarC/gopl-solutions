package main

import (
	"fmt"
	"os"
	"reflect"
	"sort"
	"strings"
)

func main() {
	if len(os.Args) <= 2 {
		fmt.Println("Give at least two words as arguments:", os.Args[1:])
	}
	for i := 1; i < len(os.Args)-1; i++ {
		fmt.Printf("Are they anagrams? %s %s %t\n", os.Args[i], os.Args[i+1], anagram(os.Args[i], os.Args[i+1]))
	}
}

func anagram(s1, s2 string) bool {

	if len(s1) != len(s2) {
		return false
	}
	if s1 == s2 {
		return false
	}

	for _, r := range s1 {
		var ok bool
		s2, ok = runePresent(s2, r)
		if !ok {
			return false
		}
	}
	return true
}

func runePresent(s2 string, r rune) (string, bool) {

	if len(s2) == 1 {
		return "", rune(s2[0]) == r
	}
	if i := strings.LastIndex(s2, string(r)); i != -1 {

		if i == 0 {
			return s2[i+1:], true // 1st el

		} else if i == len(s2)-1 {
			return s2[:i], true // last el

		} else {
			return s2[:i] + s2[i+1:], true // middle el
		}
	}
	return "", false
}

// ray-g
func isAnagramMap(s1, s2 string) bool {
	if s1 == s2 {
		return false
	}

	if len(s1) != len(s2) {
		return false
	}

	m1 := make(map[rune]int)
	m2 := make(map[rune]int)

	for _, c := range s1 {
		m1[c]++
	}

	for _, c := range s2 {
		m2[c]++
	}

	if len(m1) != len(m2) {
		return false
	}

	for k, v := range m1 {
		if m2[k] != v {
			return false
		}
	}
	return true
}

func isAnagramReflect(s1, s2 string) bool {
	if s1 == s2 {
		return false
	}

	if len(s1) != len(s2) {
		return false
	}

	a1 := strings.Split(s1, "")
	a2 := strings.Split(s2, "")

	sort.Strings(a1)
	sort.Strings(a2)

	return reflect.DeepEqual(a1, a2)
}