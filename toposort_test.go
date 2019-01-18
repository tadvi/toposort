package toposort

import (

	"testing"
)

func TestSort(t *testing.T) {
	inp := [][]string{
		{"D", "F"},
		{"A", "C"},
		{"B", "C"},
		{"C", "E"},
		{"E", "F"},
		{"F", "G"},
		{"B", "D"},
	}
	want := []string{"B", "A", "C", "E", "D", "F", "G"}
	ord := Sort(inp)
	if !equal(ord, want) {
		t.Errorf("got: %v, want %v", ord, want)
	}

	inp = [][]string{
		{"A", "B"},
		{"A", "F"},
		{"G", "A"},
		{"G", "B"},
		{"B", "H"},
		{"G", "C"},
		{"D", "C"},
		{"D", "I"},
		{"I", "C"},
		{"D", "E"},
		{"J", "E"},
		{"E", "I"},
	}
	want = []string{"J", "D", "E", "I", "G", "C", "A", "F", "B", "H"}
	ord = Sort(inp)
	if !equal(ord, want) {
		t.Errorf("got: %v, want %v", ord, want)
	}
}

func equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
