package day17_test

import (
	"advent-of-code/2021/day17"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var testData = `target area: x=20..30, y=-10..-5`

func TestPart1(t *testing.T) {
	got, err := day17.Part1(strings.NewReader(testData))
	require.NoError(t, err)
	assert.Equal(t, 45, got)
}

func TestPart2(t *testing.T) {
	got, err := day17.Part2(strings.NewReader(testData))
	require.NoError(t, err)
	assert.Equal(t, 112, got)
}
