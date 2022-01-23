package input_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/jdhenke/advent-of-code/input"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGrid(t *testing.T) {
	g, err := input.NewGrid(strings.NewReader(`abc
def`))
	require.NoError(t, err)
	assert.Equal(t, 2, g.Height())
	assert.Equal(t, 3, g.Width())
	for i, tc := range []struct {
		i, j int
		want string
	}{
		{0, 0, "a"},
		{1, 1, "e"},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			assert.Equal(t, tc.want, g.Get(tc.i, tc.j))
		})
	}
}

func TestGridInfinite(t *testing.T) {
	r := strings.NewReader(`abc
def`)
	g, err := input.NewGrid(r, input.WithInfiniteRows(), input.WithInfiniteColumns())
	require.NoError(t, err)
	for i, tc := range []struct {
		i, j int
		want string
	}{
		{0, 0, "a"},
		{1, 1, "e"},
		{2, 0, "a"},
		{1, 4, "e"},
		{3, 5, "f"},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			assert.Equal(t, tc.want, g.Get(tc.i, tc.j))
		})
	}
}
