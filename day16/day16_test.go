package day16_test

import (
	"advent-of-code/day16"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

var testData1 = `8A004A801A8002F478`

var testData2 = `620080001611562C8802118E34`

var testData3 = `C0015000016115A2E0802F182340`

var testData4 = `A0016C880162017C3686B18A3D4780`

func TestPart1(t *testing.T) {
	for i, tc := range []struct {
		input string
		want  int
	}{
		{testData1, 16},
		{testData2, 12},
		{testData3, 23},
		{testData4, 31},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got, err := day16.Part1(strings.NewReader(tc.input))
			require.NoError(t, err)
			assert.Equal(t, tc.want, got)
		})
	}
}
