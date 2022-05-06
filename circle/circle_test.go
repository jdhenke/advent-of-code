package circle_test

import (
	"testing"

	"github.com/jdhenke/advent-of-code/circle"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	e := circle.New(1, 2, 3)
	assert.Equal(t, 1, e.Value())
	e = e.Next()
	assert.Equal(t, 2, e.Value())
	e = e.Next()
	assert.Equal(t, 3, e.Value())
	e = e.Next()
	assert.Equal(t, 1, e.Value())
}

func TestFind(t *testing.T) {
	e := circle.New(1, 2, 3)
	two := e.Find(2)
	require.NotNil(t, two)
	assert.Equal(t, 2, two.Value())
	assert.Equal(t, 3, two.Next().Value())
	assert.Nil(t, e.Find(4))
}

func TestSnipAndInsert(t *testing.T) {
	e := circle.New(1, 2, 3)
	s := e.Snip(1)
	assert.Equal(t, 1, e.Value(), 1)
	assert.Equal(t, 3, e.Next().Value())
	assert.Equal(t, 1, e.Next().Next().Value())
	e = e.Next()
	e.Insert(s)
	assert.Equal(t, 3, e.Value())
	assert.Equal(t, 2, e.Next().Value())
	assert.Equal(t, 1, e.Next().Next().Value())
	assert.Equal(t, 3, e.Next().Next().Next().Value())
}
