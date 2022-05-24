package mappings

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
	elements := make(map[string]int, 50)
	Outline(elements, doc)
	fmt.Println(elements)
}

func Outline(elements map[string]int, n *html.Node) {
	if n.Type == html.ElementNode {
		elements[n.Data]++
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		Outline(elements, c)
	}
}
