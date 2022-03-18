package main

import (
	"strings"
	"testing"
)

func TestGetElementByID(t *testing.T) {
	input :=
		`<html>
  <head id="headid">
    <title id="head">
      Reading html
    </title>
  </head>
  <body>
    <p class="test" id="head2">
      Hello World!
    </p>
    <a href="www.xxx.au">
    </a>
  </body>
</html>
`
	res := GetElementByIDWrapper(strings.NewReader(input), "headid")
	if res == nil || res.Data != "head" {
		t.Fail()
	}

	res = GetElementByIDWrapper(strings.NewReader(input), "head2")
	if res == nil || res.Data != "p" {
		t.Fail()
	}
}
