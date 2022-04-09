package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

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
	_ = ElementsByTagName(doc,  "title", "h1", "a")

}

func ElementsByTagName(doc *html.Node, findNodes ...string) []*html.Node {

	var result []*html.Node
	fmt.Println(findNodes)
	nodes := make(map[string]bool, len(findNodes))
	for _, node := range findNodes {
		nodes[node] = true
	}

	pre := func(n *html.Node) {
		if n.Type == html.ElementNode && nodes[n.Data] && n.FirstChild != nil && strings.TrimSpace(n.FirstChild.Data) != "" {
			result = append(result, n.FirstChild)
			fmt.Println(n.Data, n.FirstChild.Data)
		}
	}
	ch5.ForEachNode(doc, pre, nil)

	return result
}
