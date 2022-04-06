package main

import (
	"log"
	"net/http"
	"os"

	"github.com/DagmarC/gopl-solutions/ch5"
	"golang.org/x/net/html"
)

func main() {
	resp, err := http.Get(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	ElementsByTagName(doc, "h1", "h2", "head")
}

func ElementsByTagName(doc *html.Node, findNodes ...string) []*html.Node {

	var result []*html.Node

	nodes := make(map[string]bool, len(findNodes))
	for _, node := range findNodes {
		nodes[node] = true
	}

	pre := func(n *html.Node) {
		if n.Type == html.ElementNode && nodes[n.Data] {
			result = append(result, n)
		}
	}
	ch5.ForEachNode(doc, pre, nil)

	return result
}
