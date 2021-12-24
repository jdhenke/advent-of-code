package input_test

import (
	"advent-of-code/input"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func TestForEachLine(t *testing.T) {
	data := `1
alpha

foo bar baz`
	var got []string
	err := input.ForEachLine(strings.NewReader(data), func(line string) error {
		got = append(got, line)
		return nil
	})
	require.NoError(t, err)
	assert.Equal(t, []string{"1", "alpha", "", "foo bar baz"}, got)
}

func TestForEachInt(t *testing.T) {
	data := `1
0
-2
567`
	var got []int
	err := input.ForEachInt(strings.NewReader(data), func(x int) {
		got = append(got, x)
	})
	require.NoError(t, err)
	assert.Equal(t, []int{1, 0, -2, 567}, got)
}

func TestForEachCommand(t *testing.T) {
	data := `up 1
left -3
down 566
right 0
`
	type entry struct {
		command string
		value   int
	}
	var got []entry
	err := input.ForEachCommand(strings.NewReader(data), func(cmd string, val int) {
		got = append(got, entry{cmd, val})
	})
	require.NoError(t, err)
	assert.Equal(t, []entry{
		{"up", 1},
		{"left", -3},
		{"down", 566},
		{"right", 0},
	}, got)
}
