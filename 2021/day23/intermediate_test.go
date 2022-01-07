package day23

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
	A3: 12,
	A4: 13,
	B3: 22,
	B4: 23,
	C3: 32,
	C4: 33,
	D3: 42,
	D4: 43,
}

func TestSolveSolved(t *testing.T) {
	ans, ok := solve(Board{
		A1: 10,
		A2: 11,
		B1: 20,
		B2: 21,
		C1: 30,
		C2: 31,
		D1: 40,
		D2: 41,
		A3: 12,
		A4: 13,
		B3: 22,
		B4: 23,
		C3: 32,
		C4: 33,
		D3: 42,
		D4: 43,
	})
	assert.True(t, ok)
	assert.Equal(t, 0, ans)
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
