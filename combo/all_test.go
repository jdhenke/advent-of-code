package combo_test

import (
	"fmt"
	"testing"

	"github.com/jdhenke/advent-of-code/combo"
	"github.com/stretchr/testify/assert"
)

func TestPermute(t *testing.T) {
	for i, tc := range []struct {
		length int
		vals   []int
		want   [][]int
	}{
		{
			length: 0,
			vals:   []int{1, 2, 3},
			want:   [][]int{{}},
		},
		{
			length: 3,
			vals:   []int{1},
			want:   [][]int{{1, 1, 1}},
		},
		{
			length: 1,
			vals:   []int{1, 2, 3},
			want:   [][]int{{1}, {2}, {3}},
		},
		{
			length: 2,
			vals:   []int{0, 1},
			want:   [][]int{{0, 0}, {0, 1}, {1, 0}, {1, 1}},
		},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			buf := make([]int, tc.length)
			var got [][]int
			combo.All(buf, tc.vals, func(x []int) {
				cp := make([]int, len(x))
				copy(cp, x)
				got = append(got, cp)
			})
			assert.Equal(t, tc.want, got)
		})
	}
}
