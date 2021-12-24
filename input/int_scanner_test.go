package input_test

import (
	"advent-of-code/input"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestScanInt(t *testing.T) {
	data := `1
0
-2
567`
	var got []int
	input.ScanInt(strings.NewReader(data), func(x int) {
		got = append(got, x)
	})
	assert.Equal(t, []int{1, 0, -2, 567}, got)
}
