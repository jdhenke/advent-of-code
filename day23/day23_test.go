package day23_test

import (
	"advent-of-code/day23"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

var testData = `#############
#...........#
###B#C#B#D###
  #A#D#C#A#
  #########`

func TestPart1(t *testing.T) {
	ans, err := day23.Part1(strings.NewReader(testData))
	require.NoError(t, err)
	assert.Equal(t, 12521, ans)
}

func TestPart2(t *testing.T) {
	ans, err := day23.Part2(strings.NewReader(testData))
	require.NoError(t, err)
	assert.Equal(t, 0, ans)
}
