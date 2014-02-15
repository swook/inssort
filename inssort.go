// Package inssort provides an insertion sort function with entry order changes
// indication and the specification of indices for changed entries for more
// efficient sorting.
package inssort

import (
	"errors"
	"sort"
)

var ErrInvalidIndices = errors.New("inssort: Invalid indices provided.")

// Sort sorts data using insertion sort. It considers data[idxs[0]:idxs[1]] as
// new data, where the indices considered are:
//  - No range specified:       1:
//  - Only one index specified: idxs[0]:idxs[0]+1
//  - Two indices specified:    idxs[0]:idxs[1]
// If idxs[2] specified, it only flags changes on idxs[2] first entries
func Sort(data sort.Interface, idxs ...int) (changed bool) {
	n := data.Len()

	// Nothing to sort if none or single entry
	if n < 2 {
		return false
	}

	// Determine index range (b:e) to check
	var b, e int
	l := n // Define limit of indices checked for 'changedness'
	switch len(idxs) {
	case 0:
		b, e = 1, n
	case 1:
		b, e = idxs[0], idxs[0]+1
	case 3:
		l = idxs[2]
		fallthrough
	default:
		b, e = idxs[0], idxs[1]
	}

	// Abort if invalid indices supplied
	if b > e || b < 0 || e-1 > n {
		panic(ErrInvalidIndices)
	}

	// For each entry that needs to be considered
	for i := b; i < e; i++ {
		for j, j_ := i, i-1; j > 0; j-- {
			j_ = j - 1

			// If current entry less than previous, swap
			if data.Less(j, j_) {
				data.Swap(j, j_)

				// Mark as having changed order of entries
				// if lower index swapped is under limit
				if j_ < l {
					changed = true
				}
			} else {
				break
			}
		}
	}
	return
}
