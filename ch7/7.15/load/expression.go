package load

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/DagmarC/gopl-solutions/ch7/7.13/eval"
)

func LoadExpresion(r io.Reader) (eval.Expr, eval.Env, error) {

	scanner := bufio.NewScanner(r)
	env := make(eval.Env, 0)

	fmt.Println("Enter the expression eg. x-1 or min(1, x, 2, 3), ... :")
	scanner.Scan()
	inputExpr := scanner.Text()

	fmt.Println("Enter all given ENV VARS eg. x=1 in fact variable=number. To stop type DONE anycase:")
	fmt.Println("Note: If you dont enter the env var, it will have the default value 0. To stop type DONE anycase:")

	for scanner.Scan() {
		vars := scanner.Text()
		if strings.EqualFold(strings.ToUpper(vars), "DONE") {
			break
		}
		err := parseEnvVars(inputExpr, scanner.Text(), env)
		if err != nil {
			return nil, env, err
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, env, err
	}
	e, err := eval.Parse(inputExpr)
	if err != nil {
		return nil, env, err
	}

	return e, env, nil
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
