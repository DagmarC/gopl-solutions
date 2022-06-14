package main

// import (
// 	"reflect"
// 	"strings"
// 	"time"
// )

// //!+7.7 recentlyClickedSort
// type recentlyClickedSort struct {
// 	t      []*Track
// 	recent []string
// }

// func (ts recentlyClickedSort) Len() int      { return len(ts.t) }
// func (ts recentlyClickedSort) Swap(i, j int) { ts.t[i], ts.t[j] = ts.t[j], ts.t[i] }

// func (ts recentlyClickedSort) Less(i, j int) bool {

// 	if ts.recent == nil || len(ts.recent) == 0 {
// 		return ts.t[i].Artist < ts.t[j].Artist // Default, none was recently clicked
// 	}
// 	for _, r := range ts.recent {
// 		if strings.ToLower(r) == "year" {
// 			t1 := getFieldInteger(ts.t[i], r)
// 			t2 := getFieldInteger(ts.t[j], r)
// 			if t1 != t2 {
// 				return t1 < t2
// 			}
// 		} else if strings.ToLower(r) == "length" {
// 			t1 := getFieldDuration(ts.t[i], r)
// 			t2 := getFieldDuration(ts.t[j], r)
// 			if t1 != t2 {
// 				return t1 < t2
// 			}
// 		} else {
// 			t1 := getFieldString(ts.t[i], r)
// 			t2 := getFieldString(ts.t[j], r)
// 			if t1 != t2 {
// 				return t1 < t2
// 			}
// 		}
// 	}
// 	return false
// }

// func getFieldString(t *Track, field string) string {
// 	r := reflect.ValueOf(t)
// 	f := reflect.Indirect(r).FieldByName(field)
// 	return f.String()
// }

// func getFieldInteger(t *Track, field string) int {
// 	r := reflect.ValueOf(t)
// 	f := reflect.Indirect(r).FieldByName(field)
// 	return int(f.Int())
// }

// func getFieldDuration(t *Track, field string) time.Duration {
// 	r := reflect.ValueOf(t)
// 	f := reflect.Indirect(r).FieldByName(field)
// 	return time.Duration(f.Int())
// }

// //!-7.7 recentlyClickedSort
