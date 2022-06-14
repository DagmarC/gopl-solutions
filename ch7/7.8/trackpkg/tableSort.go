package trackpkg

type ByColumns struct {
	t    []*Track
	less []Less
}

type Less func(x, y *Track) bool

func ByTitleCol(x, y *Track) bool  { return x.Title < y.Title }
func ByArtistCol(x, y *Track) bool { return x.Artist < y.Artist }
func ByAlbumCol(x, y *Track) bool  { return x.Album < y.Album }
func ByYearCol(x, y *Track) bool   { return x.Year < y.Year }
func ByLengthCol(x, y *Track) bool { return int64(x.Length) < int64(y.Length) }

func (x ByColumns) Len() int { return len(x.t) }
func (x ByColumns) Less(i, j int) bool {
	t1, t2 := x.t[i], x.t[j]
	for _, lt := range x.less {
		if !lt(t1, t2) && !lt(t2, t1) { // So they are equal
			continue
		}
		return lt(t1, t2)
	}
	return ByArtistCol(t1, t2) // default sort, if all are equal
}
func (x ByColumns) Swap(i, j int) { x.t[i], x.t[j] = x.t[j], x.t[i] }

func SortByColumns(t []*Track, cols ...Less) ByColumns {
	return ByColumns{
		t:    t,
		less: cols,
	}
}
