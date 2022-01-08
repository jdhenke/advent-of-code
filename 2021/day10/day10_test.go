package day10_test

import (
	"strings"
	"testing"

	"github.com/jdhenke/advent-of-code/2021/day10"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var testData = `[({(<(())[]>[[{[]{<()<>>
[(()[<>])]({[<{<<[]>>(
{([(<{}[<>[]}>{[]{[(<()>
(((({<>}<{<{<>}{[]{[]{}
[[<[([]))<([[{}[[()]]]
[{[{({}]{}}([{[{{{}}([]
{<[[]]>}<{[{[{[]{()[[[]
[<(<(<(<{}))><([]([]()
<{([([[(<>()){}]>(<<{{
<{([{{}}[<[[[<>{}]]]>[]]`

func TestPart1(t *testing.T) {
	got, err := day10.Part1(strings.NewReader(testData))
	require.NoError(t, err)
	assert.Equal(t, 26397, got)
}

func TestPart2(t *testing.T) {
	got, err := day10.Part2(strings.NewReader(testData))
	require.NoError(t, err)
	assert.Equal(t, 288957, got)
}
