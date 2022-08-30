package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"

	myxml "github.com/DagmarC/gopl-solutions/ch7/7.17/xml"
)

// func main() {
// 	attMatchers := []string{"id", "class"}
// 	dec := xml.NewDecoder(os.Stdin)
// 	var stack []string // stack of element names

// 	for {
// 		tok, err := dec.Token()
// 		if err == io.EOF {
// 			break
// 		} else if err != nil {
// 			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
// 			os.Exit(1)
// 		}
// 		switch tok := tok.(type) {
// 		case xml.StartElement:
// 			stack = append(stack, tok.Name.Local) // push
// 		case xml.EndElement:
// 			stack = stack[:len(stack)-1] // pop

// 		case xml.CharData:
// 			if myxml.ContainsAll(stack, os.Args[1:]) {
// 				if myxml.ContainsAll(stack, os.Args[1:]) {
// 					// for _, s := range stack {
// 					// 	if !strings.Contains(s, "=") {
// 					// 		fmt.Printf("%s ", s)
// 					// 	}
// 					// }
// 					fmt.Printf(" %s\n", tok)
// 				}
// 			}
// 		}
// 	}
// }

func main() {
	dec := xml.NewDecoder(os.Stdin)
	var stack []xml.StartElement // stack of element names

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
			stack = append(stack, tok) // push
			
		case xml.EndElement:
			stack = stack[:len(stack)-1] // pop

		case xml.CharData:
			// fmt.Printf("\t===== %s\n", tok)
			if myxml.ContainsAll(stack, os.Args[1:]) {
				for _, s := range stack {
					fmt.Printf("%s ", s.Name.Local)
				}
				fmt.Printf(": \t\t\t\t%s\n", tok)
			}
		}
	}
}
