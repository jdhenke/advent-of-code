package day14_test

import (
	"strings"
	"testing"

	"github.com/jdhenke/advent-of-code/2021/day14"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var testData = `NNCB

CH -> B
HH -> N
CB -> H
NH -> C
HB -> C
HC -> B
HN -> C
NN -> C
BH -> H
NC -> B
NB -> B
BN -> B
BB -> N
BC -> B
CC -> N
CN -> C`

func TestPart1(t *testing.T) {
	got, err := day14.Part1(strings.NewReader(testData))
	require.NoError(t, err)
	assert.Equal(t, 1588, got)
}

func TestPart2(t *testing.T) {
	got, err := day14.Part2(strings.NewReader(testData))
	require.NoError(t, err)
	assert.Equal(t, 2188189693529, got)
}
