package day13_test

import (
	"advent-of-code/day13"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

var testData = `6,10
0,14
9,10
0,3
10,4
4,11
6,0
6,12
4,1
0,13
10,12
3,4
3,0
8,4
1,10
2,14
8,10
9,0

fold along y=7
fold along x=5`

func TestPart1(t *testing.T) {
	got, err := day13.Part1(strings.NewReader(testData))
	require.NoError(t, err)
	assert.Equal(t, 17, got)
}

func TestPart2(t *testing.T) {
	_, err := day13.Part2(strings.NewReader(testData))
	require.NoError(t, err)
}
