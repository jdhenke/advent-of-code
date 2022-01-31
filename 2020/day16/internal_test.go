package day16

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDay16(t *testing.T) {
	for _, tc := range []struct {
		input string
		want  map[string]int
	}{
		{
			input: `class: 0-1 or 4-19
row: 0-5 or 8-19
seat: 0-13 or 16-19

your ticket:
11,12,13

nearby tickets:
3,9,18
15,1,5
5,14,9`,
			want: map[string]int{"class": 12, "row": 11, "seat": 13},
		},
	} {
		_, ticket, err := day16(strings.NewReader(tc.input), true)
		require.NoError(t, err)
		assert.Equal(t, tc.want, ticket)
	}
}
