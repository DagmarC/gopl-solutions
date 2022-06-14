// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 187.

// Sorting sorts a music playlist into a variety of orders.
package main

import (
	"sort"

	ti "github.com/DagmarC/gopl-solutions/ch7/7.8/trackpkg"
)

var tracks = []*ti.Track{
	{"Go", "Delilah", "From the Roots Up", 2012, ti.Length("3m38s")},
	{"Go", "Moby", "Moby", 1992, ti.Length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, ti.Length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, ti.Length("4m36s")},
}

func main() {

	// fmt.Println("\nCustom:")
	// //!+customcall
	// sort.Sort(customSort{tracks, func(x, y *Track) bool {
	// 	if x.Title != y.Title {
	// 		return x.Title < y.Title
	// 	}
	// 	if x.Year != y.Year {
	// 		return x.Year < y.Year
	// 	}
	// 	if x.Length != y.Length {
	// 		return x.Length < y.Length
	// 	}
	// 	return false
	// }})

	// sort.Sort(byRecentTableSort{tracks, []string{"Title", "Year"}})
	// sort.Sort(recentlyClickedSort{tracks, []string{"Length", "Artist"}})
	// SECOND SOLUTION
	sort.Sort(ti.SortByColumns(tracks, ti.ByLengthCol, ti.ByYearCol))
	ti.PrintTracks(tracks)
}
