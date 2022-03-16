package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks5.5: %v\n", err)
		os.Exit(1)
	}
	counts := make(map[string]int, 2)
	visit(counts, doc)
	fmt.Println(counts)

}

//!-main

//!+visit
// visit appends to links each link found in n and returns the result.
func visit(counts map[string]int, n *html.Node) {
	if n == nil {
		return
	}
	switch n.Type {
	case html.TextNode:
		if n.Data != "script" && n.Data != "style" {
			splitWords(n.Data, counts)
		}
	case html.ElementNode:
		if n.Data == "img" {
			counts["IMG"]++
		}
	}
	visit(counts, n.NextSibling)
	visit(counts, n.FirstChild)
}

func splitWords(s string, counts map[string]int) {
	input := bufio.NewScanner(strings.NewReader(s))
	input.Split(bufio.ScanWords)

	for input.Scan() {
		input.Text()
		counts["WORD"]++
	}
}

//!-visit

/*
//!+html
package html

type Node struct {
	Type                    NodeType
	Data                    string
	Attr                    []Attribute
	FirstChild, NextSibling *Node
}

type NodeType int32

const (
	ErrorNode NodeType = iota
	TextNode
	DocumentNode
	ElementNode
	CommentNode
	DoctypeNode
)

type Attribute struct {
	Key, Val string
}

func Parse(r io.Reader) (*Node, error)
//!-html
*/
