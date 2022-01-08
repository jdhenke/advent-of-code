package day15_test

import (
	"strings"
	"testing"

	"github.com/jdhenke/advent-of-code/2021/day15"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var testData = `1163751742
1381373672
2136511328
3694931569
7463417111
1319128137
1359912421
3125421639
1293138521
2311944581`

func TestPart1(t *testing.T) {
	got, err := day15.Part1(strings.NewReader(testData))
	require.NoError(t, err)
	assert.Equal(t, 40, got)
}

func TestPart2(t *testing.T) {
	got, err := day15.Part2(strings.NewReader(testData))
	require.NoError(t, err)
	assert.Equal(t, 315, got)
}
