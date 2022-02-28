package main

import (
	"testing"

	"github.com/DagmarC/gopl-solutions/utils"
)

type data struct {
	input    []string
	expected []string
}

var testData = []data{
	{input: []string{"a", "a", "b"}, expected: []string{"a", "b"}},
	{input: []string{"a", "b", "b"}, expected: []string{"a", "b"}},
	{input: []string{"a", "c", "b"}, expected: []string{"a", "c", "b"}},
	{input: []string{"a", "c", "b", "b", "b", "b"}, expected: []string{"a", "c", "b"}},
	{input: []string{"a", "b", "b", "b", "a", "a"}, expected: []string{"a", "b", "a"}},
	{input: []string{"a", "a", "a", "a"}, expected: []string{"a"}},
}

func TestRemoveAdjDuplicates(t *testing.T) {
	for _, td := range testData {
		actual := removeAdjDuplicates(td.input)
		if !utils.EqualsS(actual, td.expected) {
			t.Errorf("error expected: %v, actual: %v.", td.expected, td.input)
		}
	}
}
