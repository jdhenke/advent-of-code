package day19

import (
	"advent-of-code/input"
	"fmt"
	"io"
	"sort"
	"strconv"
	"strings"
)

type Vector []int

func (v Vector) String() string {
	return fmt.Sprint([]int(v))
}

type VectorSet []Vector

func (vs VectorSet) Column(idx int) []int {
	var out []int
	for _, v := range vs {
		out = append(out, v[idx])
	}
	return out
}

func (vs VectorSet) translate(rs []result) VectorSet {
	var out VectorSet
	for _, v := range vs {
		v2 := make(Vector, len(v))
		for _, r := range rs {
			val := v[r.oldDimension]
			if r.negated {
				val = -val
			}
			val -= r.offset
			v2[r.newDimension] = val
		}
		out = append(out, v2)
	}
	return out
}

func Part1(r io.Reader) (answer int, err error) {
	return day19(r, 12)
}

type result struct {
	negated bool
	offset       int
	newDimension int
	oldDimension int
}

func day19(r io.Reader, n int) (answer int, err error) {
	vectorSets, err := getVectorSets(r)
	if err != nil {
		return 0, err
	}
	solved := []VectorSet{vectorSets[0]}
	toMatch := []VectorSet{vectorSets[0]}
	unsolved := vectorSets[1:]

	for len(unsolved) > 0 {
		if len(toMatch) == 0 {
			panic("ran out of solutions to match against")
		}
		m := toMatch[0]
		toMatch = toMatch[1:]
		// unsolved index
		matches := map[int][]result{}
		for mi := 0; mi < len(m[0]); mi++ { // 3
			mc := m.Column(mi)
			sort.Ints(mc)
			for ui, vs := range unsolved { // 32
				for uvsi := 0; uvsi < len(vs[0]); uvsi++ { // 3
					uvsc := vs.Column(uvsi)
					sort.Ints(uvsc)
					if offset, ok := same(mc, uvsc, n); ok {
						matches[ui] = append(matches[ui], result{
							newDimension: mi,
							oldDimension: uvsi,
							negated:      false,
							offset:       offset,
						})
						continue
					}
					nvscn := negate(uvsc)
					sort.Ints(nvscn)
					if offset, ok := same(mc, nvscn, n); ok {
						matches[ui] = append(matches[ui], result{
							newDimension: mi,
							oldDimension: uvsi,
							negated:      true,
							offset:       offset,
						})
					}
				}
			}
		}
		// by count
		for _, rs := range matches {
			if len(rs) != len(solved[0][0]) {
				panic("our hopes that things only match perfectly is flawed")
			}
		}
		var newUnsolved []VectorSet
		for i, vs := range unsolved {
			if rs, ok := matches[i]; ok {
				vs = vs.translate(rs)
				toMatch = append(toMatch, vs)
				solved = append(solved, vs)
			} else {
				newUnsolved = append(newUnsolved, vs)
			}
		}
		unsolved = newUnsolved
		if len(toMatch) == 0 {
			panic("nothing matched")
		}
	}

	// find unique number of points
	uniq := make(map[string]bool)
	for _, vs := range solved {
		for _, v := range vs {
			uniq[v.String()] = true
		}
	}
	return len(uniq), nil
}

func Part2(r io.Reader) (answer int, err error) {
	return 0, nil
}

func getVectorSets(r io.Reader) ([]VectorSet, error) {
	var out []VectorSet
	var current VectorSet
	if err := input.ForEachLine(r, func(line string) error {
		if len(line) == 0 {
			out = append(out, current)
		} else if strings.HasPrefix(line, "---") {
			current = nil
		} else {
			current = append(current, parseVector(line))
		}
		return nil
	}); err != nil {
		return nil, err
	}
	out = append(out, current)
	return out, nil
}

func parseVector(line string) Vector {
	parts := strings.Split(line, ",")
	var v Vector
	for _, p := range parts {
		d, err := strconv.Atoi(p)
		if err != nil {
			panic("bad number: " + p)
		}
		v = append(v, d)
	}
	return v
}

// v1 and v2 must already be sorted, looks for an overlap of n
func same(v1, v2 []int, n int) (v2Offset int, ok bool) {
	for i := 0; i <= len(v1)-n; i++ {
		freqs1 := freqs(v1, i)
		for j := 0; j <= len(v2)-n; j++ {
			freqs2 := freqs(v2, j)
			o := overlap(freqs1, freqs2)
			if o >= n {
				return v2[j]-v1[i], true
			}
		}
	}
	return 0, false
}

func freqs(v []int, i int) map[int]int {
	out := make(map[int]int)
	offset := v[i]
	for j := i; j < len(v); j++ {
		val := v[j] - offset
		out[val]++
	}
	return out
}

func overlap(f1, f2 map[int]int) int {
	var count int
	for k, v := range f1 {
		count += min(v, f2[k])
	}
	return count
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func negate(v []int) []int {
	var out []int
	for _, x := range v {
		out = append(out, -x)
	}
	return out
}