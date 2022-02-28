package main

import (
	"bytes"
	"fmt"
)

const ASCII_SPACE = '\x20'

func main() {
	t := []byte("Hello World世界 ")
	r := reverse(t)
	fmt.Println(string(r))
}

func reverse(s []byte) []byte {
	rs := bytes.Runes(s)
	for i, j := 0, len(rs)-1; i < j; i, j = i+1, j-1 {
		rs[i], rs[j] = rs[j], rs[i]
	}
	return []byte(string(rs))
}
