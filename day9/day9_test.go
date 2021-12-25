package day9_test

import (
	"advent-of-code/day9"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
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
