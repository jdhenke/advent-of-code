package day4

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMatchesPart1(t *testing.T) {
	for i, tc := range []struct {
		matchesFunc func(digits []int) bool
		num         int
		want        bool
	}{
		{matchesPart1, 111111, true},
		{matchesPart1, 223450, false},
		{matchesPart1, 123789, false},
		{matchesPart2, 112233, true},
		{matchesPart2, 123444, false},
		{matchesPart2, 111122, true},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := tc.matchesFunc(toDigits(tc.num))
			assert.Equal(t, tc.want, got)
		})
	}
}
