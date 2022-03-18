// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 133.

// Outline prints the outline of an HTML document tree.
package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

var out *os.File = os.Stdout

var doc *html.Node

var id = flag.String("id", "head", "ID of the element you want to search for.")

const ID string = "id"

func main() {
	flag.Parse()
	for _, url := range os.Args[1:] {
		if !strings.Contains(url, "https://") {
			continue
		}
		doc = outline(url)

		match := ElementByID(doc, *id)
		if match == nil {
			printf("Main: ID %v not found \n", *id)
			continue
		}
		printf("Main: ID %v found, element=%v\n", *id, match.Data)
	}
}

func ElementByID(doc *html.Node, id string) *html.Node {
	return forEachNode(doc, id, startElement, endElement)
}

func outline(url string) *html.Node {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return doc
}

func GetElementByIDWrapper(input io.Reader, id string) *html.Node {
	doc, err := html.Parse(input)
	if err != nil {
		log.Fatal(err)
	}

	match := ElementByID(doc, id)
	if match == nil {
		printf("W: ID %v not found \n", id)
		return nil
	}

	printf("W: ID %v found, element=%v\n", id, match.Data)
	return match
}

//!+forEachNode
// forEachNode calls the functions pre(x) and post(x) for each node
// x in the tree rooted at n. Both functions are optional.
// pre is called before the children are visited (preorder) and
// post is called after (postorder).
func forEachNode(n *html.Node, id string, pre, post func(n *html.Node, id string) bool) *html.Node {
	if n == nil {
		return nil
	}
	if pre != nil && pre(n, id) {
		return n
	}
	m := forEachNode(n.NextSibling, id, pre, post)
	if m != nil {
		return m
	}
	m = forEachNode(n.FirstChild, id, pre, post)
	if m != nil {
		return m
	}
	if post != nil && post(n, id) {
		return n
	}
	return m
}

//!-forEachNode

//!+startend
var depth int

func startElement(n *html.Node, id string) bool {
	if n.Type == html.ElementNode {
		if byID(n, id) {
			return true
		}
	}
	return false
}

func byID(n *html.Node, id string) bool {
	for _, a := range n.Attr {
		if a.Key == ID && a.Val == id {
			return true
		}
	}
	return false
}

func endElement(n *html.Node, id string) bool {
	if n.Type == html.ElementNode {
		if byID(n, id) {
			return true
		}
	}
	return false
}

// printf will used the defined out in the system.
func printf(format string, a ...interface{}) {
	fmt.Fprintf(out, format, a...)
}

//!-startend
