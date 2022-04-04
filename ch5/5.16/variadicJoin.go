package main

import (
	"bytes"
	"fmt"
)

func main() {
	ss := []string{"Cata", "sthro", "pic"}
	fmt.Printf("Alternate variadic strings.Join called on %v, result: %s", ss, Join("X", ss...))
}

func Join(sep string, ss ...string) string {
	var buffer bytes.Buffer

	for _, s := range ss {
		buffer.WriteString(s)
		buffer.WriteString(sep)
	}
	return buffer.String()
}
