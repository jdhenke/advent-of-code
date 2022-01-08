package day9_test

import (
	"strings"
	"testing"

	"github.com/jdhenke/advent-of-code/2021/day9"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var testData = `2199943210
3987894921
9856789892
8767896789
9899965678`

func TestPart1(t *testing.T) {
	got, err := day9.Part1(strings.NewReader(testData))
	require.NoError(t, err)
	assert.Equal(t, 15, got)
}

func TestPart2(t *testing.T) {
	got, err := day9.Part2(strings.NewReader(testData))
	require.NoError(t, err)
	assert.Equal(t, 1134, got)
}
