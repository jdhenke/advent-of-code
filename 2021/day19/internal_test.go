package day19

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

const allSame6 = `--- scanner 0 ---
-1,-1,1
-2,-2,2
-3,-3,3
-2,-3,1
5,6,-4
8,0,7

--- scanner 0 ---
1,-1,1
2,-2,2
3,-3,3
2,-1,3
-5,4,-6
-8,-7,0

--- scanner 0 ---
-1,-1,-1
-2,-2,-2
-3,-3,-3
-1,-3,-2
4,6,5
-7,0,8

--- scanner 0 ---
1,1,-1
2,2,-2
3,3,-3
1,3,-2
-4,-6,5
7,0,8

--- scanner 0 ---
1,1,1
2,2,2
3,3,3
3,1,2
-6,-4,-5
0,7,-8`

const allSame3 = `--- scanner 0 ---
0,2
4,1
3,3

--- scanner 1 ---
-1,-1
-5,0
-2,1`

func TestDay19(t *testing.T) {
	ans, err := day19(strings.NewReader(allSame6), 6)
	require.NoError(t, err)
	assert.Equal(t, 6, ans)
}

func TestDay19_AllSame3(t *testing.T) {
	ans, err := day19(strings.NewReader(allSame3), 3)
	require.NoError(t, err)
	assert.Equal(t, 3, ans)
}
