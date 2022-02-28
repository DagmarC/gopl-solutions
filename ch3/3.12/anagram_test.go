package main

import (
	"reflect"
	"runtime"
	"testing"
)

func testIsAnagram(t *testing.T, f func(string, string) bool) {
	tcs := []struct {
		s1      string
		s2      string
		expects bool
	}{
		{"abc", "cba", true},
		{"abc", "abc", false},
		{"abc", "abcd", false},
		{"abc", "ab", false},
		{"abc", "", false},
	}

	for _, tc := range tcs {
		ret := f(tc.s1, tc.s2)
		if ret != tc.expects {
			t.Errorf("Failed %v, s1: %s, s2: %s, result: %v", getFunctionName(f), tc.s1, tc.s2, ret)
		}
	}
}

func getFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

func TestIsAnagramReflect(t *testing.T) {
	testIsAnagram(t, anagram)
}

func benchmarkIsAnagram(b *testing.B, f func(string, string) bool) {
	for i := 0; i < b.N; i++ {
		f("abcdefghijklmnopqrstuvwxyz", "zyxwvutsrqponmlkjihgfedcba")
	}
}

func BenchmarkAnagram(b *testing.B) {
	benchmarkIsAnagram(b, anagram)
}
func BenchmarkIsAnagramMap(b *testing.B) {
	benchmarkIsAnagram(b, isAnagramMap)
}

func BenchmarkIsAnagramReflect(b *testing.B) {
	benchmarkIsAnagram(b, isAnagramReflect)
}
