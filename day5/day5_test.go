package day5_test

import (
	"advent-of-code/day5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

var testData = `0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2`

func TestPart1(t *testing.T) {
	got, err := day5.Part1(strings.NewReader(testData))
	require.NoError(t, err)
	assert.Equal(t, 5, got)
}
