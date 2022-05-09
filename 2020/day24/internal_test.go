package day24

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChanges(t *testing.T) {
	x, y := 0, 0
	n := 0
	forEachInstruction("nwwswee", func(delta change) {
		x, y = x+delta.dx, y+delta.dy
		n++
	})
	assert.Equal(t, 5, n)
	assert.Equal(t, 0, x)
	assert.Equal(t, 0, y)
}
