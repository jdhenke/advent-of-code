package day11_test

import (
	"advent-of-code/2021/day11"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var testData = `5483143223
2745854711
5264556173
6141336146
6357385478
4167524645
2176841721
6882881134
4846848554
5283751526`

func TestPart1(t *testing.T) {
	got, err := day11.Part1(strings.NewReader(testData))
	require.NoError(t, err)
	assert.Equal(t, 1656, got)
}

func TestPart2(t *testing.T) {
	got, err := day11.Part2(strings.NewReader(testData))
	require.NoError(t, err)
	assert.Equal(t, 195, got)
}
