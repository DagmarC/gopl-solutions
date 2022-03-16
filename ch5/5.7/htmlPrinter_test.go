package main

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func TestPretify(t *testing.T) {
	input :=
		`<html>
  <head id="head">
    <title>
      Reading html
    </title>
  </head>
  <body>
    <p class="test">
      Hello World!
    </p>
    <a href="www.xxx.au">
    </a>
  </body>
</html>
`
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	out = w
	defer w.Close()

	err = prettify(strings.NewReader(input))
	if err != nil {
		t.Log()
		t.Fail()
	}

	buf := make([]byte, 2048)
	n, err := r.Read(buf)
	if err != nil {
		t.Log(err)
		t.Fatal()
	}
	// Back to normal os.Stdout
	out = os.Stdout

	res := string(buf[:n])
	fmt.Println(res)
	fmt.Println("----")

	_, err = html.Parse(strings.NewReader(res)) // Parsing should be successful on res as well
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	if ok := strings.EqualFold(res, input); !ok {
		fmt.Println(input)
		t.Fail()
	}
}
