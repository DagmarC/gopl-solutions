// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 214.
//!+
package xml

import (
	"bytes"
	"encoding/xml"
	"strings"
)

// ContainsAll reports whether x contains the elements of y, in order.
func ContainsAll(x []xml.StartElement, y []string) bool {

	for len(y) <= len(x) {
		if len(y) == 0 {
			return true
		}
		xAttrs := GetAttVals(x[0])
		if x[0].Name.Local == y[0] {
			y = y[1:]
		}

		for i := 0; len(y) > i && strings.Contains(y[i], "=") && len(xAttrs) > 0 && strings.Contains(xAttrs, y[i]); i++ {
			y = y[1:]
		}

		x = x[1:]
	}
	return false
}

func GetAttVals(tok xml.StartElement) string {
	attMatchers := []string{"id", "class", "name"}
	var buf bytes.Buffer
	for _, att := range tok.Attr {
		for _, m := range attMatchers {
			if att.Name.Local == m {
				buf.WriteString(strings.Join([]string{att.Name.Local, att.Value}, "="))
				buf.WriteString(" ")
			}
		}
	}
	return strings.TrimSpace(buf.String())
}

//!-
