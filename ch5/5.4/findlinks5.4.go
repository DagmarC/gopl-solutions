package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

var mapping = map[string]string{"img": "src", "script": "src", "link": "href"}

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

//!-main

//!+visit
// visit appends to links each link found in n and returns the result.
func visit(links []string, n *html.Node) []string {
	if n == nil {
		return links
	}
	if n.Type == html.ElementNode {
		if _, ok := mapping[n.Data]; ok {
			for _, a := range n.Attr {
				if a.Key == mapping[n.Data] {
					links = append(links, a.Val)
				}
			}
		}
	}
	links = visit(links, n.NextSibling)
	links = visit(links, n.FirstChild)

	return links
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
