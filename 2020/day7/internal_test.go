package day7

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseBags(t *testing.T) {
	// 3 vibrant green bags, 4 plaid blue bags, 2 drab brown bags.
	// 1 pale magenta bag.
	// no other bags.
	for _, tc := range []struct {
		line string
		want map[string]int
	}{
		{
			"3 vibrant green bags, 4 plaid blue bags, 2 drab brown bags.", map[string]int{
				"vibrant green": 3,
				"plaid blue":    4,
				"drab brown":    2,
			},
		}, {
			"1 pale magenta bag.", map[string]int{
				"pale magenta": 1,
			},
		}, {
			"no other bags.", map[string]int{},
		},
	} {
		assert.Equal(t, tc.want, parseBags(tc.line))
	}
}
