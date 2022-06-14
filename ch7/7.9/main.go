package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
	"sort"

	ti "github.com/DagmarC/gopl-solutions/ch7/7.8/trackpkg"
)

var tracksTempl = template.Must(template.New("tracksTempl").Parse(`
<h1>Tracks</h1>
<table style="width:100%">
<tr style='text-align: left'>
  <th><a href='./table?sort=titleCol'>Title</th>
  <th><a href='./table?sort=artistCol'>Artist</th>
  <th><a href='./table?sort=albumCol'>Album</th>
  <th><a href='./table?sort=yearCol'>Year</th>
  <th><a href='./table?sort=lengthCol'>Length</th>

</tr>
{{range .}}
<tr>
  <td>{{.Title}}</a></td>
  <td>{{.Artist}}</td>
  <td>{{.Album}}</a></td>
  <td>{{.Year}}</a></td>
  <td>{{.Length}}</a></td>
</tr>
{{end}}
</table>
`))

var tracks = []*ti.Track{
	{"Go", "Delilah", "From the Roots Up", 2012, ti.Length("3m38s")},
	{"Go", "Moby", "Moby", 1992, ti.Length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, ti.Length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, ti.Length("4m36s")},
	{"My Angel", "Martin Solveig", "Sth", 2011, ti.Length("4m36s")},
}

var recentCols ti.RecentCols

const columnsN int = 5

func main() {
	http.Handle("/", http.HandlerFunc(tracksSite))
	http.Handle("/table", http.HandlerFunc(tracksTable))

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

// --------------HTTP LISTENERS--------------

func tracksSite(w http.ResponseWriter, r *http.Request) {
	tracksRender(w)
}

func tracksTable(w http.ResponseWriter, r *http.Request) {
	queries := map[string][]string(r.URL.Query())

	if sortCols, ok := queries["sort"]; ok {
		addRecentCols(sortCols)
		sort.Sort(ti.SortByColumns(tracks, recentCols.Reverse()...))
	}

	tracksRender(w) // Re-render it after query processing, eg. sorting.
}

// --------------HTTP TEMPLATE RENDER--------------

func tracksRender(out io.Writer) {
	if err := tracksTempl.Execute(out, tracks); err != nil {
		log.Fatal(err)
	}
}

// --------------SORTING BY COLUMS--------------

func addRecentCols(cols []string) {
	if recentCols.Length() == columnsN {
		for range cols {
			recentCols.RemoveFirst() // VIA API CALL YOU COUL ADD MORE COLS AT ONCE
		}
	}
	for _, col := range cols {
		switch col {
		case "titleCol":
			recentCols.Add(ti.ByTitleCol)
		case "albumCol":
			recentCols.Add(ti.ByAlbumCol)
		case "artistCol":
			recentCols.Add(ti.ByArtistCol)
		case "yearCol":
			recentCols.Add(ti.ByYearCol)
		case "lengthCol":

			recentCols.Add(ti.ByLengthCol)
		}
	}
}
