package day21_test

import (
	"advent-of-code/day21"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

var testData = `Player 1 starting position: 4
Player 2 starting position: 8`

func TestPart1(t *testing.T) {
	ans, err := day21.Part1(strings.NewReader(testData))
	require.NoError(t, err)
	assert.Equal(t, 739785, ans)
}

func TestPart2(t *testing.T) {
	ans, err := day21.Part2(strings.NewReader(testData))
	require.NoError(t, err)
	assert.Equal(t, 444356092776315, ans)
}
