package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/DagmarC/gopl-solutions/ch4/4.10/github"
)

// 1. declare the template
const githubTempl = `<h1> Github Issues </h1>
<table>
	<tr style='text-align: left'> 
		<th>Number</th>
		<th>Title</th>
		<th>State</th>
		<th>User</th>
	</tr>
	{{range .Items}}
	<tr>
		<td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
		<td>{{.Title}}</td>
		<td>{{.State}}</td>
		<td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
	</tr>
	{{end}}
</table>`

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		githubRender(w)
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func githubRender(out io.Writer) {
	issues, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	// 2. Create template and parse (use the declared template here)
	var template = template.Must(template.New("githubIssues").Parse(githubTempl))
	if err != nil {
		log.Fatal(err)
	}
	// 3. Execute on newly created template - supply the struct here.
	if err = template.Execute(out, issues); err != nil {
		log.Fatal(err)
	}
}
