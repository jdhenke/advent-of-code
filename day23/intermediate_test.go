package day23

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

var testBoard = Board{
	A1: 11,
	A2: 41,
	B1: 10,
	B2: 30,
	C1: 20,
	C2: 31,
	D1: 21,
	D2: 40,
}

func TestSolve(t *testing.T) {
	ans, ok := solve(testBoard)
	require.True(t, ok)
	assert.Equal(t, 12521, ans)
}

const testData = `#############
#...........#
###B#C#B#D###
  #A#D#C#A#
  #########`

func TestParseBoard(t *testing.T) {
	b, err := parseBoard(strings.NewReader(testData))
	require.NoError(t, err)
	assert.Equal(t, testBoard, b)
}
