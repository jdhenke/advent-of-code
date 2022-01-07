package day1_test

import (
	"advent-of-code/2021/day1"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testData = `199
200
208
210
200
207
240
269
260
263`

func TestPart1(t *testing.T) {
	ans, err := day1.Part1(strings.NewReader(testData))
	assert.NoError(t, err)
	assert.Equal(t, 7, ans)
}

func TestPart2(t *testing.T) {
	ans, err := day1.Part2(strings.NewReader(testData))
	assert.NoError(t, err)
	assert.Equal(t, 5, ans)
}
