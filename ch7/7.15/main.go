package main

import (
	"fmt"
	"log"
	"os"

	"github.com/DagmarC/gopl-solutions/ch7/7.15/load"
)

func main() {
	expr, err := load.LoadExpresion(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	env, err := load.LoadEnvVars(os.Stdin, expr.String())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Given expression: %s \t Result: %g\n", expr.String(), expr.Eval(env))
}