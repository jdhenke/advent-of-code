package day7

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCost(t *testing.T) {
	for _, tc := range []struct{ from, to, cost int }{
		{16, 5, 66},
		{1, 5, 10},
		{2, 5, 6},
	} {
		assert.Equal(t, tc.cost, moveCost(tc.from, tc.to), tc)
	}
}
