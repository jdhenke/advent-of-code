package day19

import (
	"fmt"
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
		stop bool
		want []string
	}{
		{
			s:    "",
			stop: false,
			want: []string(nil),
		},
		{
			s:    "a",
			stop: false,
			want: []string(nil),
		},
		{
			s:    "ab",
			stop: false,
			want: []string{""},
		},
		{
			s:    "abc",
			stop: false,
			want: []string{"c"},
		},
		{
			s:    "abac",
			stop: false,
			want: []string{"ac", ""},
		},
		{
			s:    "abac",
			stop: true,
			want: []string{"ac"},
		},
		{
			s:    "abacd",
			stop: false,
			want: []string{"acd", "d"},
		},
	} {
		t.Run(fmt.Sprintf("%s:%v", tc.s, tc.stop), func(t *testing.T) {
			var got []string
			r.Match(tc.s, func(remaining string) (ok bool) {
				got = append(got, remaining)
				return tc.stop
			})
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestEqualPartsRule(t *testing.T) {
	r := equalPartsRule(literalRule("a"), literalRule("b"))
	for _, tc := range []struct {
		s    string
		stop bool
		want []string
	}{
		{
			s:    "",
			stop: false,
			want: []string(nil),
		},
		{
			s:    "a",
			stop: false,
			want: []string(nil),
		},
		{
			s:    "ab",
			stop: false,
			want: []string{""},
		},
		{
			s:    "abb",
			stop: false,
			want: []string{"b"},
		},
		{
			s:    "abba",
			stop: false,
			want: []string{"ba"},
		},
	} {
		t.Run(fmt.Sprintf("%s:%v", tc.s, tc.stop), func(t *testing.T) {
			var got []string
			r.Match(tc.s, func(remaining string) (ok bool) {
				got = append(got, remaining)
				return tc.stop
			})
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestEqualPartsMultipleMatches(t *testing.T) {
	r := equalPartsRule(literalRule("a"), literalRule("a"))
	for _, tc := range []struct {
		s    string
		stop bool
		want []string
	}{
		{
			s:    "aa",
			stop: false,
			want: []string{""},
		},
		{
			s:    "aaaa",
			stop: false,
			want: []string{"aa", ""},
		},
		{
			s:    "aaaaaa",
			stop: false,
			want: []string{"aaaa", "aa", ""},
		},
		{
			s:    "aaaaaa",
			stop: true,
			want: []string{"aaaa"},
		},
	} {
		t.Run(fmt.Sprintf("%s:%v", tc.s, tc.stop), func(t *testing.T) {
			var got []string
			r.Match(tc.s, func(remaining string) (ok bool) {
				got = append(got, remaining)
				return tc.stop
			})
			assert.Equal(t, tc.want, got)
		})
	}
}
