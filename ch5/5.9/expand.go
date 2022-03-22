package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {

	fmt.Println(expand("$foo", strings.ToUpper))
}

// expand replace "$foo" in s with foo("foo") and returns the result
func expand(s string, f func(string) string) string {
	re := regexp.MustCompile(`\$\w+`)

	return re.ReplaceAllStringFunc(s, f)
}
