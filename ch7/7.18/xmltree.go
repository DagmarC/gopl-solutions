package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

// Construct a tree of generic nodes that represents it.
// Nodes are of two kinds:CharDatanodes represent text strings,
// andElementnodes represent named elements and their attributes.

type Node interface{} // CharData or *Element

type CharData string

type Element struct {
	Type     xml.Name
	Attr     []xml.Attr
	Children []Node
}
type Tree struct {
	root Node
}

var fPtr = flag.String("xml", "cd_catalog.xml", "xml file, if none then cd_catalog.xml is used")

// go run . -xml=books.xml
func main() {
	flag.Parse()

	f, err := os.Open(*fPtr)
	if err != nil {
		log.Fatal(err)
	}
	t := XmlTree(f)
	t.Print()
}

func XmlTree(r io.Reader) Tree {
	dec := xml.NewDecoder(r)
	var stack []Node
	var tree Tree
	var tailn Node

	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
			os.Exit(1)
		}

		switch tok := tok.(type) {
		case xml.StartElement:
			var el Element
			el.Attr = tok.Attr
			el.Type = tok.Name

			stack = append(stack, &el) // push
			if tree.root == nil {
				tree.root = &el
			} else {
				switch tailn := tailn.(type) {
				case *Element:
					tailn.Children = append(tailn.Children, &el)
				}
			}
			tailn = stack[len(stack)-1]

		case xml.EndElement:
			stack = stack[:len(stack)-1] // pop
			if len(stack) > 0 {
				tailn = stack[len(stack)-1]
			}
		case xml.CharData:
			switch tailn := tailn.(type) {
			case *Element:
				tailn.Children = append(tailn.Children, CharData(tok))
			}
		}
	}
	return tree
}

func (t *Tree) Print() {
	traverseRec(t.root, 0)
	fmt.Println()
}

func traverseRec(n Node, depth int) {
	if n == nil {
		return
	}

	switch n := n.(type) {
	case *Element:
		tab(depth)
		fmt.Printf("%s\t", n.Type.Local)

		for _, nn := range n.Children {
			traverseRec(nn, depth+1)
		}
	case CharData:
		if strings.TrimSpace(string(n)) != "" {
			fmt.Printf(": %s", n)
		}
		fmt.Println()
	}
}
func tab(d int) {
	for d > 0 {
		fmt.Printf("\t")
		d--
	}
}
