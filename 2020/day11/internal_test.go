package day11

import (
	"fmt"
	"strings"
	"testing"

	"github.com/jdhenke/advent-of-code/input"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPart2GetOccupied(t *testing.T) {
	for i, tc := range []struct {
		grid string
		i, j int
		want int
	}{
		{
			grid: `.......#.
...#.....
.#.......
.........
..#L....#
....#....
.........
#........
...#.....`,
			i:    4,
			j:    3,
			want: 8,
		}, {
			grid: `.............
.L.L.#.#.#.#.
.............`,
			i:    1,
			j:    3,
			want: 1,
		}, {
			grid: `.............
.L.L.#.#.#.#.
.............`,
			i:    1,
			j:    2,
			want: 0,
		}, {
			grid: `.##.##.
#.#.#.#
##...##
...L...
##...##
#.#.#.#
.##.##.`,
			i:    3,
			j:    3,
			want: 0,
		},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			g, err := input.NewGrid(strings.NewReader(tc.grid))
			require.NoError(t, err)
			got := part2GetOccupied(g, tc.i, tc.j)
			assert.Equal(t, tc.want, got)
		})
	}
}
