package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"

	"github.com/DagmarC/gopl-solutions/ch7/7.15/load"
)

func main() {

	http.HandleFunc("/", display)       //register handler
	http.HandleFunc("/calc", calculate) //register handler

	http.ListenAndServe("localhost:8000", nil)
}

func display(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles("calcexpr.gohtml")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	fmt.Println("expression:", r.Form["exp"])
	fmt.Println("env vars:", r.Form["env"])

}

// calculate endpoint: /calc?exp=x-1
func calculate(w http.ResponseWriter, r *http.Request) {
	var result float64

	err := r.ParseForm()
	if err != nil {
		fmt.Fprintf(w, "error: parsing response")
	}
	//Parse url parameters passed, then parse the response packet for the POST body (request body)
	// attention: If you do not call ParseForm method, the following data can not be obtained form

	resTmpl, err := template.ParseFiles("result.gohtml")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println("method:", r.Method) //get request method
	if r.Method == "POST" {
		//Call to ParseForm makes form fields available.

		expS := r.PostFormValue("exp")
		envS := r.PostFormValue("env")

		exp, err := load.LoadExpresion(strings.NewReader(expS))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// http://localhost:8000/calc?exp=x-1&env=x%3D2
		env, err := load.LoadEnvVars(strings.NewReader(envS), exp.String())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		result = exp.Eval(env)

		if err := resTmpl.Execute(w, result); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	} else {
		if err := resTmpl.Execute(w, result); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}

}
