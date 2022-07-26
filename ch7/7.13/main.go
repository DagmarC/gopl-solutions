package main

import (
	"fmt"
	"math"

	"github.com/DagmarC/gopl-solutions/ch7/7.13/eval"
)

func main() {
	problems := []struct {
		expr string
		env  eval.Env
		want string
	}{
		{"sqrt(A / pi)", eval.Env{"A": 87616, "pi": math.Pi}, "167"},
		{"pow(x, 3) + pow(y, 3)",eval. Env{"x": 12, "y": 1}, "1729"},
		{"5 / 9 * (F - 32)", eval.Env{"F": -40}, "-40"},
		// {"5 / 9 * (F - 32)", eval.Env{"F": 212}, "100"},
		// //!-Eval
		// // additional tests that don't appear in the book
		{"-1 + -x", eval.Env{"x": 1}, "-2"},
		{"-1 - x", eval.Env{"x": 1}, "-2"},
		// //!+Eval
	}
	for _, e := range problems {
		expr, err := eval.Parse(e.expr)
		if err != nil {
			continue
		}
		fmt.Println(expr)
	}

}

