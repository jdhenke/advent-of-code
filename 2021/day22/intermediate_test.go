package day22

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func TestCube_Split(t *testing.T) {
	cubes, err := readCubes(strings.NewReader(`on x=0..3,y=0..3,z=0..0
off x=2..5,y=2..5,z=0..0`))
	require.NoError(t, err)
	c1, c2 := cubes[0], cubes[1]
	assert.True(t, c1.Intersects(c2))
	assert.True(t, c2.Intersects(c1))
	splits := c1.Split(c2)
	for _, s := range splits {
		assert.True(t, c1.Intersects(s))
		assert.False(t, c2.Intersects(s))
	}
}
