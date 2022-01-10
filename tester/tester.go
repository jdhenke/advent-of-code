package tester

import (
	"fmt"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/jdhenke/advent-of-code/solution"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type Tester struct {
	t        *testing.T
	solution solution.Func
}

func New(t *testing.T, solution solution.Func) *Tester {
	return &Tester{
		t:        t,
		solution: solution,
	}
}

func (tst *Tester) Run(cases ...Case) {
	for i, c := range cases {
		tst.t.Helper()
		tst.t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Helper()
			opts := &testCaseOptions{}
			c.apply(opts)
			if opts.want == nil {
				t.Fatal("No want specified for test case.")
			}
			want := *opts.want
			var r io.Reader
			switch {
			case opts.inputString != nil:
				r = strings.NewReader(*opts.inputString)
			case opts.inputPath != nil:
				f, err := os.Open(*opts.inputPath)
				require.NoError(t, err)
				defer func() {
					assert.NoError(t, f.Close())
				}()
				r = f
			default:
				t.Fatal("No input specified for test case.")
			}
			got, err := tst.solution(r)
			require.NoError(t, err)
			assert.Equal(t, want, got)
		})
	}
}

type testCaseOptions struct {
	inputString *string
	inputPath   *string
	want        *int
}

type Case interface {
	apply(opts *testCaseOptions)
}

type caseFunc func(opts *testCaseOptions)

func (cf caseFunc) apply(tc *testCaseOptions) {
	cf(tc)
}

func FromFile(path string) *CaseBuilder {
	return &CaseBuilder{
		cf: func(opts *testCaseOptions) {
			opts.inputPath = &path
		},
	}
}

func FromString(s string) *CaseBuilder {
	return &CaseBuilder{
		cf: func(tc *testCaseOptions) {
			tc.inputString = &s
		},
	}
}

type CaseBuilder struct {
	cf caseFunc
}

func (c *CaseBuilder) Want(ans int) Case {
	return caseFunc(func(opts *testCaseOptions) {
		c.cf.apply(opts)
		opts.want = &ans
	})
}
