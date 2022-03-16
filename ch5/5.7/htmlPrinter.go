// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 133.

// Outline prints the outline of an HTML document tree.
package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

var out *os.File = os.Stdout

func main() {
	var output string
	flag.StringVar(&output, "out", "stdout", "output - stdout / file")

	flag.Parse()

	if output == "file" {
		fmt.Println("File creation")
		f, err := os.Create("out.html")
		if err != nil {
			fmt.Errorf("Error while openinf a file.")
		}
		defer f.Close()
		out = f
	}

	for _, url := range os.Args[1:] {
		if strings.Contains(url, "https") {
			outline(url)
		}
	}
}

func outline(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	prettify(resp.Body)
	return nil
}

func prettify(input io.Reader) error {
	doc, err := html.Parse(input)
	if err != nil {
		return err
	}
	forEachNode(doc, startElement, endElement)
	return nil
}

// forEachNode calls the functions pre(x) and post(x) for each node
// x in the tree rooted at n. Both functions are optional.
// pre is called before the children are visited (preorder) and
// post is called after (postorder).
func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil && n.FirstChild != nil {
		post(n)
	}
}

var depth int

func startElement(n *html.Node) {

	switch n.Type {
	case html.ElementNode:
		printElement(n)
	case html.TextNode:
		printText(n)
	case html.CommentNode:
		printComment(n)
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		printf("%*s</%s>\n", depth*2, "", n.Data)
	}
}

func printAttr(n *html.Node) {
	for _, a := range n.Attr {
		printf(" %s=\"%s\"", a.Key, a.Val)
	}
}

func printElement(n *html.Node) {
	printf("%*s<%s", depth*2, "", n.Data)
	printAttr(n)

	end := "/>"
	if n.FirstChild != nil {
		end = ">\n"
		depth++
	}
	printf(end)
}

func printText(n *html.Node) {
	if strings.TrimSpace(n.Data) != "" {
		printf("%*s%s\n", depth*2, "", strings.TrimSpace(n.Data))
	}
}

func printComment(n *html.Node) {
	printf("%*s<!--%s-->\n", depth*2, "", n.Data)
}

// printf will used the defined out in the system.
func printf(format string, a ...interface{}) {
	fmt.Fprintf(out, format, a...)
}
