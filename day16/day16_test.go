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

func TestPart2(t *testing.T) {
	for i, tc := range []struct {
		input string
		want  int
	}{
		{"C200B40A82", 3},
		{"04005AC33890", 54},
		{"880086C3E88112", 7},
		{"CE00C43D881120", 9},
		{"D8005AC2A8F0", 1},
		{"F600BC2D8F", 0},
		{"9C005AC2F8F0", 0},
		{"9C0141080250320F1802104A08", 1},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got, err := day16.Part2(strings.NewReader(tc.input))
			require.NoError(t, err)
			assert.Equal(t, tc.want, got)
		})
	}
}
