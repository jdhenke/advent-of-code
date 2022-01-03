package day24_test

import (
	"advent-of-code/day24"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

var testData = ""

func TestPart1(t *testing.T) {
	ans, err := day24.Part1(strings.NewReader(testData))
	require.NoError(t, err)
	assert.Equal(t, 0, ans)
}

func TestPart2(t *testing.T) {
	ans, err := day24.Part2(strings.NewReader(testData))
	require.NoError(t, err)
	assert.Equal(t, 0, ans)
}
