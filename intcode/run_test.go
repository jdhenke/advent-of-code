package intcode_test

import (
	"testing"

	"github.com/jdhenke/advent-of-code/intcode"
	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	got := intcode.Run(intcode.Code{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50})
	assert.Equal(t, 3500, got)
}
