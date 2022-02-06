package day18

import (
	"bufio"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var testData = `1 + 2 * 3 + 4 * 5 + 6`

func TestScanner(t *testing.T) {
	scan := bufio.NewScanner(strings.NewReader(testData))
	scan.Split(splitFunc)
	var got []string
	for scan.Scan() {
		got = append(got, scan.Text())
	}
	require.NoError(t, scan.Err())
	assert.Equal(t, []string{"1", "+", "2", "*", "3", "+", "4", "*", "5", "+", "6"}, got)
}
