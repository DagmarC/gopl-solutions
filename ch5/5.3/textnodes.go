package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
		os.Exit(1)
	}
	visit(doc)
}

func visit(n *html.Node) {
	if n.Type == html.TextNode {
		if n.Parent.Data != "script" && n.Parent.Data != "style" {
			fmt.Println(n.Data)
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		visit(c)
	}
}
