package intcode_test

import (
	"strings"
	"testing"

	"github.com/jdhenke/advent-of-code/intcode"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParse(t *testing.T) {
	want := intcode.Code{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50}
	got, err := intcode.Parse(strings.NewReader(`1,9,10,3,2,3,11,0,99,30,40,50`))
	require.NoError(t, err)
	assert.Equal(t, want, got)
}
