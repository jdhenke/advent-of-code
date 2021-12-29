package day18_test

import (
	"advent-of-code/day18"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

var testData = `[[[0,[5,8]],[[1,7],[9,6]]],[[4,[1,2]],[[1,4],2]]]
[[[5,[2,8]],4],[5,[[9,9],0]]]
[6,[[[6,2],[5,6]],[[7,6],[4,7]]]]
[[[6,[0,7]],[0,9]],[4,[9,[9,0]]]]
[[[7,[6,4]],[3,[1,3]]],[[[5,5],1],9]]
[[6,[[7,3],[3,2]]],[[[3,8],[5,7]],4]]
[[[[5,4],[7,7]],8],[[8,3],8]]
[[9,3],[[9,9],[6,[4,9]]]]
[[2,[[7,7],7]],[[5,8],[[9,3],[0,2]]]]
[[[[5,2],5],[8,[3,7]]],[[5,[7,5]],[4,4]]]`

func TestPart1(t *testing.T) {
	for i, tc := range []struct {
		input string
		want  int
	}{
		{testData, 4140},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got, err := day18.Part1(strings.NewReader(tc.input))
			require.NoError(t, err)
			assert.Equal(t, tc.want, got)
		})
	}
}
