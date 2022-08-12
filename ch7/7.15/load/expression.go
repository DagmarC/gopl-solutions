package load

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/DagmarC/gopl-solutions/ch7/7.13/eval"
)

func LoadExpresion(r io.Reader) (eval.Expr, error) {

	scanner := bufio.NewScanner(r)

	fmt.Println("Enter the expression eg. x-1 or min(1, x, 2, 3), ... :")
	scanner.Scan()
	inputExpr := strings.TrimSpace(scanner.Text())

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	e, err := eval.Parse(inputExpr)
	if err != nil {
		return nil, err
	}

	return e, nil
}

func LoadEnvVars(r io.Reader, exp string) (eval.Env, error) {
	scanner := bufio.NewScanner(r)
	env := make(eval.Env, 0)
	fmt.Println("Enter all given ENV VARS: variable=number.")
	fmt.Println("Multiple CSV format: x=1,y=2,z=3,...")

	scanner.Scan()
	vars := scanner.Text()

	for _, v := range strings.Split(vars, ",") {
		err := parseEnvVars(exp, strings.TrimSpace(v), env)
		if err != nil {
			return env, err
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return env, nil
}

func parseEnvVars(e string, input string, env eval.Env) error {

	vars := strings.Split(input, "=")
	if len(vars) != 2 {
		return fmt.Errorf("wrong format of env variable, got %s want sth like x=1", input)
	}

	// Variable extraction and small check from expression.
	varE := vars[0]
	if !strings.Contains(e, varE) {
		return fmt.Errorf("wrong env variable, got %s want variable from expression %s", varE, e)
	}

	// Value conversion and small check.
	val, err := strconv.ParseFloat(vars[1], 64)
	if err != nil {
		return fmt.Errorf("conversion error, got %s but want number.\n %s", vars[1], err)
	}

	env[eval.Var(varE)] = val

	return nil
}
