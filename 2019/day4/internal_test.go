package day4

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMatches(t *testing.T) {
	for _, tc := range []struct {
		num  int
		want bool
	}{
		{111111, true},
		{223450, false},
		{123789, false},
	} {
		t.Run(fmt.Sprint(tc.num), func(t *testing.T) {
			got := matches(tc.num)
			assert.Equal(t, tc.want, got)
		})
	}
}
