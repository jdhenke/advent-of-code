package day25

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoop(t *testing.T) {
	assert.Equal(t, 5764801, loop(7, 8))
	assert.Equal(t, 17807724, loop(7, 11))
	assert.Equal(t, 14897079, loop(5764801, 11))
	assert.Equal(t, 14897079, loop(17807724, 8))
	assert.Equal(t, loop(loop(7, 8), 11), loop(loop(7, 11), 8))
}

func TestFindLoop(t *testing.T) {
	l1, ok1 := findLoopNum(5764801)
	assert.True(t, ok1)
	assert.Equal(t, 8, l1)
	l2, ok2 := findLoopNum(17807724)
	assert.True(t, ok2)
	assert.Equal(t, 11, l2)
}
