package day3_test

import (
	"advent-of-code/day3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

var testData = ``

func TestPart1(t *testing.T) {
	ans, err := day3.Part1(strings.NewReader(testData))
	require.NoError(t, err)
	assert.Equal(t, "", ans)
}

func TestPart2(t *testing.T) {
	ans, err := day3.Part2(strings.NewReader(testData))
	require.NoError(t, err)
	assert.Equal(t, "", ans)
}
