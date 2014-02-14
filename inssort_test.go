package inssort

import (
	"reflect"
	"sort"
	"testing"
)

var sortTests = []struct {
	in      sort.IntSlice
	idxs    []int
	out     sort.IntSlice
	changed bool
}{
	// Simple no change situation
	{sort.IntSlice{1, 2, 3}, []int{}, sort.IntSlice{1, 2, 3}, false},

	// Complete sort situation
	{sort.IntSlice{4, 3, 2, 1}, []int{}, sort.IntSlice{1, 2, 3, 4}, true},

	// Partial sort with changed entries
	{sort.IntSlice{4, 3, 2, 1}, []int{3}, sort.IntSlice{1, 4, 3, 2}, true},

	// Partial sort without changed entries
	{sort.IntSlice{3, 2, 4}, []int{2}, sort.IntSlice{3, 2, 4}, false},

	// Partial sort with changed entries and defined indices range
	{sort.IntSlice{3, 4, 2, 5, 3}, []int{2, 4}, sort.IntSlice{2, 3, 4, 5, 3}, true},

	// Sort where changed elements outside diff limit
	{sort.IntSlice{2, 3, 4, 6, 5}, []int{2, 5, 3}, sort.IntSlice{2, 3, 4, 5, 6}, false},
}

func TestSort(t *testing.T) {
	var in []int
	var changed bool
	for i, test := range sortTests {
		in = make(sort.IntSlice, test.in.Len())
		copy(in, test.in)
		changed = Sort(test.in, test.idxs...)
		if changed != test.changed || !reflect.DeepEqual(test.in, test.out) {
			t.Fatalf("[%d] Sort(%v, %v) = (%v, %v), want (%v, %v)",
				i, in, test.idxs, test.in, changed, test.out, test.changed)
		}
	}
}
