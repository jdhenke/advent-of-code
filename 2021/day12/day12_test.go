package day12_test

import (
	"advent-of-code/2021/day12"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var test1 = `start-A
start-b
A-c
A-b
b-d
A-end
b-end`

var test2 = `dc-end
HN-start
start-kj
dc-start
dc-HN
LN-dc
HN-end
kj-sa
kj-HN
kj-dc`

var test3 = `fs-end
he-DX
fs-he
start-DX
pj-DX
end-zg
zg-sl
zg-pj
pj-he
RW-he
fs-DX
pj-RW
zg-RW
start-pj
he-WI
zg-he
pj-fs
start-RW`

func TestPart1(t *testing.T) {
	for i, tc := range []struct {
		data string
		want int
	}{
		{
			data: test1,
			want: 10,
		},
		{
			data: test2,
			want: 19,
		},
		{
			data: test3,
			want: 226,
		},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got, err := day12.Part1(strings.NewReader(tc.data))
			require.NoError(t, err)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestPart2(t *testing.T) {
	for i, tc := range []struct {
		data string
		want int
	}{
		{
			data: test1,
			want: 36,
		},
		{
			data: test2,
			want: 103,
		},
		{
			data: test3,
			want: 3509,
		},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got, err := day12.Part2(strings.NewReader(tc.data))
			require.NoError(t, err)
			assert.Equal(t, tc.want, got)
		})
	}
}