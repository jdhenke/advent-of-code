package day6_test

import (
	"advent-of-code/day6"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

var testData = `3,4,3,1,2`

func TestPart1(t *testing.T) {
	ans, err := day6.Part1(strings.NewReader(testData))
	require.NoError(t, err)
	assert.Equal(t, 5934, ans)
}

func TestPart2(t *testing.T) {
	ans, err := day6.Part2(strings.NewReader(testData))
	require.NoError(t, err)
	assert.Equal(t, 26984457539, ans)
}
