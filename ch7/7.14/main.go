package main

import (
	"fmt"
	"strconv"

	"github.com/DagmarC/gopl-solutions/ch7/7.13/eval"
)

func main() {
	problems := []struct {
		expr string
		env  eval.Env
		want string
	}{
		{"min(1, 2, 3, 4, 5)", eval.Env{}, "1"},
		{"min(10, 2, 30, -4, 5)", eval.Env{}, "-4"},
		{"min(10, 2, 30, -4, x, 5)", eval.Env{"x": -100}, "-100"},
	}
	for _, p := range problems {
		expr, err := eval.Parse(p.expr)
		if err != nil {
			continue
		}

		fWant, err := strconv.ParseFloat(p.want, 64)
		if err != nil {
			fmt.Println(err)
		}
		res := expr.Eval(p.env)
		if fWant != res {
			fmt.Printf("Error got %g want %g\n", res, fWant)
		}
		fmt.Printf("Got %g <--- %s, where env=%v\n", res, p.expr, p.env)
	}
}
