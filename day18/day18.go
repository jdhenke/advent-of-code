package day18

import (
	"advent-of-code/input"
	"bytes"
	"fmt"
	"io"
	"strconv"
)

type SnailFishNumber struct {
	Value *int
	Left  *SnailFishNumber
	Right *SnailFishNumber
}

func (snf *SnailFishNumber) Reduce() {
	reduced := false
	for !reduced {
		// process explosions first
		var (
			last               *int
			explodedRightValue *int
		)
		snf.traverse(0, func(depth int, f *SnailFishNumber) (stop bool) {
			if explodedRightValue != nil && f.Value != nil {
				*f.Value = *f.Value + *explodedRightValue
				return true
			}
			if depth == 4 && f.Value == nil {
				// To explode a pair, the pair's left value is added to the first regular number to the left of the
				// exploding pair (if any), and the pair's right value is added to the first regular number to the right of
				// the exploding pair (if any). Exploding pairs will always consist of two regular numbers. Then, the entire
				// exploding pair is replaced with the regular number 0.
				if last != nil {
					*last = *last + *f.Left.Value
				}
				explodedRightValue = f.Right.Value
				f.Left = nil
				f.Right = nil
				z := 0
				f.Value = &z
				return false
			}
			// ignore the left child of the thing that's about to be exploded
			if depth <= 4 && f.Value != nil {
				last = f.Value
			}
			return false
		})
		if explodedRightValue != nil {
			continue
		}

		// process splits
		if snf.traverse(0, func(depth int, f *SnailFishNumber) (stop bool) {
			if f.Value != nil && *f.Value > 9 {
				val := *f.Value
				leftVal := val / 2
				rightVal := (val + 1) / 2
				f.Left = &SnailFishNumber{Value: &leftVal}
				f.Right = &SnailFishNumber{Value: &rightVal}
				f.Value = nil
				return true
			}
			return false
		}) {
			continue
		}

		reduced = true
	}
}

func (snf *SnailFishNumber) traverse(depth int, visit func(depth int, f *SnailFishNumber) (stop bool)) (stop bool) {
	if snf.Left != nil {
		if snf.Left.traverse(depth+1, visit) {
			return true
		}
	}
	if visit(depth, snf) {
		return true
	}
	if snf.Right != nil {
		if snf.Right.traverse(depth+1, visit) {
			return true
		}
	}
	return false
}

func (snf *SnailFishNumber) String() string {
	var buf bytes.Buffer
	snf.string(&buf)
	return buf.String()
}

func (snf *SnailFishNumber) string(w io.Writer) {
	if snf.Value != nil {
		_, _ = fmt.Fprint(w, *snf.Value)
	} else {
		_, _ = fmt.Fprint(w, "[")
		snf.Left.string(w)
		_, _ = fmt.Fprint(w, ",")
		snf.Right.string(w)
		_, _ = fmt.Fprint(w, "]")
	}
}

func Part1(r io.Reader) (ans int, err error) {
	var sum *SnailFishNumber
	if err := input.ForEachLine(r, func(line string) error {
		snf := parse(line)
		if sum == nil {
			sum = snf
		} else {
			sum = add(sum, snf)
		}
		sum.Reduce()
		return nil
	}); err != nil {
		return 0, err
	}
	return magnitude(sum), nil
}

func parse(line string) *SnailFishNumber {
	if line[0:1] == "[" {
		// stop when [i:i+1] is the middle comma
		i, depth := 0, 0
		for ; !(depth == 1 && line[i:i+1] == ","); i++ {
			if line[i:i+1] == "[" {
				depth++
			} else if line[i:i+1] == "]" {
				depth--
			}
		}
		return &SnailFishNumber{
			Left:  parse(line[1:i]),
			Right: parse(line[i+1 : len(line)-1]),
		}
	} else {
		val, err := strconv.Atoi(line)
		if err != nil {
			panic("bad literal number: " + line)
		}
		return &SnailFishNumber{
			Value: &val,
		}
	}
}

func add(a, b *SnailFishNumber) *SnailFishNumber {
	return &SnailFishNumber{
		Left:  a,
		Right: b,
	}
}

func magnitude(snf *SnailFishNumber) int {
	if snf.Value != nil {
		return *snf.Value
	}
	return 3*magnitude(snf.Left) + 2*magnitude(snf.Right)
}
