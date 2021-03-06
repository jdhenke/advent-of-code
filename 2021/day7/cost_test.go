package day7

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCost(t *testing.T) {
	for _, tc := range []struct{ from, to, cost int }{
		{16, 5, 66},
		{1, 5, 10},
		{2, 5, 6},
	} {
		assert.Equal(t, tc.cost, part2Cost(tc.from, tc.to), tc)
	}
}
