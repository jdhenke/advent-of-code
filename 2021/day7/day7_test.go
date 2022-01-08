package day7_test

import (
	"strings"
	"testing"

	"github.com/jdhenke/advent-of-code/2021/day7"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var testData = `16,1,2,0,4,2,7,1,2,14`

func TestPart1(t *testing.T) {
	got, err := day7.Part1(strings.NewReader(testData))
	require.NoError(t, err)
	assert.Equal(t, 37, got)
}

func TestPart2(t *testing.T) {
	got, err := day7.Part2(strings.NewReader(testData))
	require.NoError(t, err)
	assert.Equal(t, 168, got)
}
