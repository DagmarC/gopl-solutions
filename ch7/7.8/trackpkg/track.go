// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 187.

// Sorting sorts a music playlist into a variety of orders.
package trackpkg

import (
	"time"
)

//!+main
type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, Length("3m38s")},
	{"Go", "Moby", "Moby", 1992, Length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, Length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, Length("4m36s")},
}

func Length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

//!-main

//!-printTracks

//!+artistcode
type byArtist []*Track

func (x byArtist) Len() int           { return len(x) }
func (x byArtist) Less(i, j int) bool { return x[i].Artist < x[j].Artist }
func (x byArtist) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

//!-artistcode

//!+yearcode
type byYear []*Track

func (x byYear) Len() int           { return len(x) }
func (x byYear) Less(i, j int) bool { return x[i].Year < x[j].Year }
func (x byYear) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

//!-yearcode

//!+customcode
type customSort struct {
	t    []*Track
	less func(x, y *Track) bool
}

func (x customSort) Len() int           { return len(x.t) }
func (x customSort) Less(i, j int) bool { return x.less(x.t[i], x.t[j]) }
func (x customSort) Swap(i, j int)      { x.t[i], x.t[j] = x.t[j], x.t[i] }

//!-customcode

//!+7.9 Sort By RecentColumns

type RecentCols struct {
	cols []Less
}

func (rc *RecentCols) Cols() []Less {
	return rc.cols
}

func (rc *RecentCols) Reverse() []Less {
	for i := 0; i < rc.Length()/2; i++ {
		j := rc.Length() - i - 1
		rc.cols[i], rc.cols[j] = rc.cols[j], rc.cols[i]
	}
	return rc.cols
}

func (rc *RecentCols) Length() int {
	return len(rc.cols)
}

func (rc *RecentCols) RemoveFirst() {
	len := rc.Length()
	if len == 1 {
		rc.cols = []Less{} // make it empty
		return
	}
	rc.cols = rc.cols[1:]
}

func (rc *RecentCols) Add(col Less) {
	rc.cols = append(rc.cols, col)
}

//!-7.9 Sort By RecentColumns
