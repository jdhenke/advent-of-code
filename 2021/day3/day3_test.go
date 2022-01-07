package day3_test

import (
	"advent-of-code/2021/day3"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var testData = `00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010`

func TestPart1(t *testing.T) {
	ans, err := day3.Part1(strings.NewReader(testData))
	require.NoError(t, err)
	assert.Equal(t, 198, ans)
}

func TestPart2(t *testing.T) {
	ans, err := day3.Part2(strings.NewReader(testData))
	require.NoError(t, err)
	assert.Equal(t, 230, ans)
}
