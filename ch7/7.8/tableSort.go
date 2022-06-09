package main

type byColumns struct {
	t    []*Track
	less []less
}

type less func(x, y *Track) bool

func byTitleCol(x, y *Track) bool  { return x.Title < y.Title }
func byArtistCol(x, y *Track) bool { return x.Artist < y.Artist }
func byAlbumCol(x, y *Track) bool  { return x.Album < y.Album }
func byYearCol(x, y *Track) bool   { return x.Year < y.Year }
func byLengthCol(x, y *Track) bool { return int64(x.Length) < int64(y.Length) }

func (x byColumns) Len() int { return len(x.t) }
func (x byColumns) Less(i, j int) bool {
	t1, t2 := x.t[i], x.t[j]
	for _, lt := range x.less {
		if !lt(t1, t2) && !lt(t2, t1) { // So they are equal
			continue
		}
		return lt(t1, t2)
	}
	return byArtistCol(t1, t2) // default sort, if all are equal
}
func (x byColumns) Swap(i, j int) { x.t[i], x.t[j] = x.t[j], x.t[i] }

func sortByColumns(t []*Track, cols ...less) byColumns {
	return byColumns{
		t:    t,
		less: cols,
	}
}
