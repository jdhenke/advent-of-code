package day14

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart2Addresses(t *testing.T) {
	for i, tc := range []struct {
		mask string
		addr int
		want []int
	}{
		{
			mask: "0000",
			addr: 0,
			want: []int{0},
		},
		{
			mask: "0101",
			addr: 0,
			want: []int{5},
		},
		{
			mask: "0101",
			addr: 10,
			want: []int{15},
		},
		{
			mask: "010X",
			addr: 10,
			want: []int{14, 15},
		},
		{
			mask: "01XX",
			addr: 10,
			want: []int{12, 13, 14, 15},
		},
		{
			mask: "000000000000000000000000000000X1001X",
			addr: 42,
			want: []int{26, 27, 58, 59},
		},
		{
			mask: "00000000000000000000000000000000X0XX",
			addr: 26,
			want: []int{16, 17, 18, 19, 24, 25, 26, 27},
		},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := part2Addresses(tc.mask, tc.addr)
			assert.Equal(t, tc.want, got)
		})
	}
}
