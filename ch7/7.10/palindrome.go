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
	for j, i := s.Len() - 1, 0; j > s.Len()/2; j,i = j-1, i+1 {
		if s.Less(i, j) || s.Less(j, i) {
			return false
		}
	}
	return true
}
