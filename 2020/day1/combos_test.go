package day1

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCombos(t *testing.T) {
	var got [][]int
	combos([]int{1, 2, 3, 4, 5}, 3, 0, func(c []int) {
		d := make([]int, len(c))
		copy(d, c)
		sort.Ints(d)
		got = append(got, d)
	})
	want := [][]int{
		{1, 2, 3},
		{1, 2, 5},
		{1, 2, 4},
		{1, 3, 4},
		{1, 3, 5},
		{1, 4, 5},
		{2, 5, 4},
		{2, 5, 3},
		{2, 4, 3},
		{5, 3, 4},
	}
	assert.Len(t, got, len(want))
	for _, e := range want {
		sort.Ints(e)
		assert.Contains(t, got, e)
	}
}
