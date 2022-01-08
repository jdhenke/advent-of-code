package day25_test

import (
	"strings"
	"testing"

	"github.com/jdhenke/advent-of-code/2021/day25"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var testData = `v...>>.vv>
.vv>>.vv..
>>.>v>...v
>>v>>.>.v.
v>v.vv.v..
>.>>..v...
.vv..>.>v.
v.v..>>v.v
....v..v.>`

func TestPart1(t *testing.T) {
	ans, err := day25.Part1(strings.NewReader(testData))
	require.NoError(t, err)
	assert.Equal(t, 58, ans)
}

func TestPart2(t *testing.T) {
	ans, err := day25.Part2(strings.NewReader(testData))
	require.NoError(t, err)
	assert.Equal(t, 0, ans)
}
