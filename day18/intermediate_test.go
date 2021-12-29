package day18

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParse(t *testing.T) {
	input := `[[[[7,7],2],[[9,2],4]],[[[9,1],5],[[9,6],[6,4]]]]`
	snf := parse(input)
	assert.Equal(t, input, snf.String())
	var nums []int
	snf.traverse(0, func(_ int, f *SnailFishNumber) bool {
		if f.Value != nil {
			nums = append(nums, *f.Value)
		}
		return false
	})
	assert.Equal(t, []int{7, 7, 2, 9, 2, 4, 9, 1, 5, 9, 6, 6, 4}, nums)
}

func TestExplode(t *testing.T) {
	for i, tc := range []struct {
		input string
		want  string
	}{
		{`[[[[[9,8],1],2],3],4]`, `[[[[0,9],2],3],4]`},
		{`[7,[6,[5,[4,[3,2]]]]]`, `[7,[6,[5,[7,0]]]]`},
		{`[[6,[5,[4,[3,2]]]],1]`, `[[6,[5,[7,0]]],3]`},
		{`[[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]]`, `[[3,[2,[8,0]]],[9,[5,[7,0]]]]`},
		{`[[[[[4,3],4],4],[7,[[8,4],9]]],[1,1]]`, `[[[[0,7],4],[[7,8],[6,0]]],[8,1]]`},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			snf := parse(tc.input)
			snf.Reduce()
			assert.Equal(t, tc.want, snf.String())
		})
	}
}
