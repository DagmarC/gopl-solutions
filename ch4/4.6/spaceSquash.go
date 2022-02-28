package main

import (
	"bytes"
	"fmt"
	"unicode"
)

const ASCII_SPACE = '\x20'

func main() {
	t := []byte("  Hello   \t    \r   World  ")
	t = utf8tasciiSpace(t)
	fmt.Println(string(t))
}

// utf8tasciiSpace squashes each run of adjacent Unicode spaces in a UTF-8-encoded []byte slice into a single ASCII space.
func utf8tasciiSpace(b []byte) []byte {

	rs := bytes.Runes(b)
	for i := 0; i < len(rs); i++ {
		if !unicode.IsSpace(rs[i]) {
			continue
		}
		var j = i + 1
		for ; j < len(rs) && unicode.IsSpace(rs[j]); j++ {
		}
		rs = ASCIISpace(i, j, rs)
	}
	return []byte(string(rs))
}

func ASCIISpace(start, end int, rs []rune) []rune {
	fmt.Println(start, end, rs, len(rs))
	rs[start] = ASCII_SPACE

	if end == start+1 {
		return rs
	}

	if end == len(rs) {
		rs = rs[:start+1]
		return rs
	}

	copy(rs[start+1:], rs[end:])
	rs = rs[:len(rs)-(end-start)+1]
	return rs
}
