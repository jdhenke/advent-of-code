package day2_test

import (
	"strings"
	"testing"

	"github.com/jdhenke/advent-of-code/2021/day2"
	"github.com/stretchr/testify/assert"
)

var testData = `forward 5
down 5
forward 8
up 3
down 8
forward 2`

func TestPart1(t *testing.T) {
	ans, err := day2.Part1(strings.NewReader(testData))
	assert.NoError(t, err)
	assert.Equal(t, 150, ans)
}

func TestPart2(t *testing.T) {
	ans, err := day2.Part2(strings.NewReader(testData))
	assert.NoError(t, err)
	assert.Equal(t, 900, ans)
}
