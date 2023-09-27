package day21

import (
	"fmt"
	"github.com/jdhenke/advent-of-code/input"
	"io"
	"strings"
)

/*
Part1 Prompt

--- Day 21: Monkey Math ---
The monkeys are back! You're worried they're going to try to steal your stuff
again, but it seems like they're just holding their ground and making various
monkey noises at you.

Eventually, one of the elephants realizes you don't speak monkey and comes over
to interpret. As it turns out, they overheard you talking about trying to find
the grove; they can show you a shortcut if you answer their riddle.

Each monkey is given a job: either to yell a specific number or to yell the
result of a math operation. All of the number-yelling monkeys know their number
from the start; however, the math operation monkeys need to wait for two other
monkeys to yell a number, and those two other monkeys might also be waiting on
other monkeys.

Your job is to work out the number the monkey named root will yell before the
monkeys figure it out themselves.

For example:

	root: pppw + sjmn
	dbpl: 5
	cczh: sllz + lgvd
	zczc: 2
	ptdq: humn - dvpt
	dvpt: 3
	lfqf: 4
	humn: 5
	ljgn: 2
	sjmn: drzm * dbpl
	sllz: 4
	pppw: cczh / lfqf
	lgvd: ljgn * ptdq
	drzm: hmdt - zczc
	hmdt: 32

Each line contains the name of a monkey, a colon, and then the job of that
monkey:

- A lone number means the monkey's job is simply to yell that number.
- A job like aaaa + bbbb means the monkey waits for monkeys aaaa and bbbb to
yell each of their numbers; the monkey then yells the sum of those two numbers.
- aaaa - bbbb means the monkey yells aaaa's number minus bbbb's number.
- Job aaaa * bbbb will yell aaaa's number multiplied by bbbb's number.
- Job aaaa / bbbb will yell aaaa's number divided by bbbb's number.

So, in the above example, monkey drzm has to wait for monkeys hmdt and zczc to
yell their numbers. Fortunately, both hmdt and zczc have jobs that involve
simply yelling a single number, so they do this immediately: 32 and 2. Monkey
drzm can then yell its number by finding 32 minus 2: 30.

Then, monkey sjmn has one of its numbers (30, from monkey drzm), and already
has its other number, 5, from dbpl. This allows it to yell its own number by
finding 30 multiplied by 5: 150.

This process continues until root yells a number: 152.

However, your actual situation involves considerably more monkeys. What number
will the monkey named root yell?
*/
func Part1(r io.Reader) (answer int, err error) {
	m, err := parse(r)
	if err != nil {
		return 0, err
	}
	return part1(m, "root"), nil
}

/*
Part2 Prompt

--- Part Two ---
Due to some kind of monkey-elephant-human mistranslation, you seem to have
misunderstood a few key details about the riddle.

First, you got the wrong job for the monkey named root; specifically, you got
the wrong math operation. The correct operation for monkey root should be =,
which means that it still listens for two numbers (from the same two monkeys as
before), but now checks that the two numbers match.

Second, you got the wrong monkey for the job starting with humn:. It isn't a
monkey - it's you. Actually, you got the job wrong, too: you need to figure out
what number you need to yell so that root's equality check passes. (The number
that appears after humn: in your input is now irrelevant.)

In the above example, the number you need to yell to pass root's equality test
is 301. (This causes root to get the same number, 150, from both of its
monkeys.)

What number do you yell to pass root's equality test?
*/
func Part2(r io.Reader) (answer int, err error) {
	m, err := parse(r)
	if err != nil {
		return 0, err
	}
	return part2(m), nil
}

type expr struct {
	k    string
	op   string
	l, r string
	v    int
}

const (
	opVal  = "v"
	opHumn = "h"
)

func parse(r io.Reader) (map[string]*expr, error) {
	m := make(map[string]*expr)
	if err := input.ForEachLine(r, func(line string) error {
		var (
			key string
			v   int
		)
		if _, err := fmt.Sscanf(line, "%s %d", &key, &v); err == nil {
			key = key[:4]
			m[key] = &expr{k: key, op: opVal, v: v}
			return nil
		}
		var l, op, r string
		if _, err := fmt.Sscanf(line, "%s %s %s %s", &key, &l, &op, &r); err != nil {
			return fmt.Errorf("%s: %w", line, err)
		}
		key = key[:4]
		m[key] = &expr{k: key, op: op, l: l, r: r}
		return nil
	}); err != nil {
		return nil, err
	}
	return m, nil
}

var ops = map[string]func(l, r int) int{
	"+": func(l, r int) int {
		return l + r
	},
	"-": func(l, r int) int {
		return l - r
	},
	"*": func(l, r int) int {
		return l * r
	},
	"/": func(l, r int) int {
		return l / r
	},
}

func part1(m map[string]*expr, k string) int {
	e := m[k]
	if e.op == opVal {
		return e.v
	}
	l, r := part1(m, e.l), part1(m, e.r)
	return ops[e.op](l, r)
}

func part2(m map[string]*expr) int {
	root, humn := m["root"], m["humn"]
	root.op, humn.op, humn.v = "=", opHumn, 0

	// collapse all equations involving constants that we can
	for {
		changed := false
		for _, e := range m {
			if strings.Contains("+-*/", e.op) && m[e.l].op == opVal && m[e.r].op == opVal {
				changed = true
				e.v = ops[e.op](m[e.l].v, m[e.r].v)
				e.op = opVal
			}
		}
		if !changed {
			break
		}
	}

	// now force each side of root to equal each other
	l, r := m[root.l], m[root.r]
	force(m, l, r.v)
	force(m, r, l.v)
	return humn.v
}

// Note: being called when e is a val expression is a noop, so is safe to call on both sides to avoid always having to
// check which side is a constant and which still needs solving. This is only OK because we can assume the exactly one
// side will always be solved until we get to the humn expression.
func force(m map[string]*expr, e *expr, val int) {
	switch e.op {
	case opHumn:
		e.v = val
	case "+": // l + r = val
		force(m, m[e.l], val-m[e.r].v)
		force(m, m[e.r], val-m[e.l].v)
	case "-": // l - r = val
		force(m, m[e.l], val+m[e.r].v)
		force(m, m[e.r], m[e.l].v-val)
	case "*": // l * r = val
		if m[e.r].op == opVal {
			force(m, m[e.l], val/m[e.r].v)
		}
		if m[e.l].op == opVal {
			force(m, m[e.r], val/m[e.l].v)
		}
	case "/": // l / r = val
		force(m, m[e.l], val*m[e.r].v)
		force(m, m[e.r], m[e.l].v/val)
	case opVal:
	default:
		panic("? op: " + e.op)
	}
}
