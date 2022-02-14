package day19

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOneOrMore(t *testing.T) {
	r := oneOrMoreRule(seqRule([]Rule{
		literalRule("a"),
		orRule([]Rule{
			literalRule("b"),
			literalRule("c"),
		}),
	}))
	for _, tc := range []struct {
		s    string
		ok   bool
		want []string
	}{
		{
			s:    "",
			ok:   true,
			want: []string(nil),
		},
		{
			s:    "a",
			ok:   true,
			want: []string(nil),
		},
		{
			s:    "ab",
			ok:   true,
			want: []string{""},
		},
		{
			s:    "abc",
			ok:   true,
			want: []string{"c"},
		},
		{
			s:    "abac",
			ok:   true,
			want: []string{"ac", ""},
		},
		{
			s:    "abac",
			ok:   false,
			want: []string{"ac"},
		},
		{
			s:    "abacd",
			ok:   true,
			want: []string{"acd", "d"},
		},
	} {
		t.Run(tc.s, func(t *testing.T) {
			var got []string
			r.Match(tc.s, func(remaining string) (ok bool) {
				got = append(got, remaining)
				return tc.ok
			})
			assert.Equal(t, tc.want, got)
		})
	}
}
