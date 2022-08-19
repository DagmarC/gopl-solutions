package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"

	myxml "github.com/DagmarC/gopl-solutions/ch7/7.17/xml"
)

func main() {
	dec := xml.NewDecoder(os.Stdin)
	var stack []string // stack of element names
	var elCount int
	attMatchers := []string{"id", "class", "name"}

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
			stack = append(stack, tok.Name.Local) // push
			elCount++                             // count elements and attributes pushed to stack

			// append to stack attributes: class=X or id=Y
			for _, att := range tok.Attr {
				for _, m := range attMatchers {
					if att.Name.Local == m {
						target := strings.Join([]string{att.Name.Local, att.Value}, "=")
						stack = append(stack, target)
						elCount++
					}
				}

			}
		case xml.EndElement:
			for elCount > 0 { // Pop elements + attributes
				stack = stack[:len(stack)-1] // pop
				elCount--
			}
		case xml.CharData:
			if myxml.ContainsAll(stack, os.Args[1:]) {
				fmt.Printf("%s: %s\n", strings.Join(stack, " "), tok)
			}
		}
	}
}
