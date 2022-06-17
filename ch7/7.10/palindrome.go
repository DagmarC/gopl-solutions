package main

import (
	"fmt"
	"sort"
)

func main() {

	s := sort.IntSlice{1, 1}
	t := sort.StringSlice{"a", "b", "a"}
	u := sort.StringSlice{"a", "b", "b"}

	fmt.Println("Is Palindrome??", s, isPalindrome(s))
	fmt.Println("Is Palindrome??", t, isPalindrome(t))
	fmt.Println("Is Palindrome??", u, isPalindrome(u))


}

func isPalindrome(s sort.Interface) bool {
	var i int
	for j := s.Len() - 1; j > s.Len()/2; j-- {
		if s.Less(i, j) || s.Less(j, i) {
			return false
		}
		i++
	}
	return true
}
