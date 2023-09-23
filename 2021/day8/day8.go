package day8

import (
	"io"
	"sort"
	"strings"

	"github.com/jdhenke/advent-of-code/input"
)

/*
Part1 Prompt

--- Day 8: Seven Segment Search ---
You barely reach the safety of the cave when the whale smashes into the cave
mouth, collapsing it. Sensors indicate another exit to this cave at a much
greater depth, so you have no choice but to press on.

As your submarine slowly makes its way through the cave system, you notice that
the four-digit seven-segment displays in your submarine are malfunctioning;
they must have been damaged during the escape. You'll be in a lot of trouble
without them, so you'd better figure out what's wrong.

Each digit of a seven-segment display is rendered by turning on or off any of
seven segments named a through g:

	  0:      1:      2:      3:      4:
	 aaaa    ....    aaaa    aaaa    ....
	b    c  .    c  .    c  .    c  b    c
	b    c  .    c  .    c  .    c  b    c
	 ....    ....    dddd    dddd    dddd
	e    f  .    f  e    .  .    f  .    f
	e    f  .    f  e    .  .    f  .    f
	 gggg    ....    gggg    gggg    ....

	  5:      6:      7:      8:      9:
	 aaaa    aaaa    aaaa    aaaa    aaaa
	b    .  b    .  .    c  b    c  b    c
	b    .  b    .  .    c  b    c  b    c
	 dddd    dddd    ....    dddd    dddd
	.    f  e    f  .    f  e    f  .    f
	.    f  e    f  .    f  e    f  .    f
	 gggg    gggg    ....    gggg    gggg

So, to render a 1, only segments c and f would be turned on; the rest would be
off. To render a 7, only segments a, c, and f would be turned on.

The problem is that the signals which control the segments have been mixed up
on each display. The submarine is still trying to display numbers by producing
output on signal wires a through g, but those wires are connected to segments
randomly. Worse, the wire/segment connections are mixed up separately for each
four-digit display! (All of the digits within a display use the same
connections, though.)

So, you might know that only signal wires b and g are turned on, but that
doesn't mean segments b and g are turned on: the only digit that uses two
segments is 1, so it must mean segments c and f are meant to be on. With just
that information, you still can't tell which wire (b/g) goes to which segment
(c/f). For that, you'll need to collect more information.

For each display, you watch the changing signals for a while, make a note of
all ten unique signal patterns you see, and then write down a single four digit
output value (your puzzle input). Using the signal patterns, you should be able
to work out which pattern corresponds to which digit.

For example, here is what you might see in a single entry in your notes:

	acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab |
	cdfeb fcadb cdfeb cdbaf

(The entry is wrapped here to two lines so it fits; in your notes, it will all
be on a single line.)

Each entry consists of ten unique signal patterns, a | delimiter, and finally
the four digit output value. Within an entry, the same wire/segment connections
are used (but you don't know what the connections actually are). The unique
signal patterns correspond to the ten different ways the submarine tries to
render a digit using the current wire/segment connections. Because 7 is the
only digit that uses three segments, dab in the above example means that to
render a 7, signal lines d, a, and b are on. Because 4 is the only digit that
uses four segments, eafb means that to render a 4, signal lines e, a, f, and b
are on.

Using this information, you should be able to work out which combination of
signal wires corresponds to each of the ten digits. Then, you can decode the
four digit output value. Unfortunately, in the above example, all of the digits
in the output value (cdfeb fcadb cdfeb cdbaf) use five segments and are more
difficult to deduce.

For now, focus on the easy digits. Consider this larger example:

	be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb |
	fdgacbe cefdb cefbgd gcbe
	edbfga begcd cbg gc gcadebf fbgde acbgfd abcde gfcbed gfec |
	fcgedb cgb dgebacf gc
	fgaebd cg bdaec gdafb agbcfd gdcbef bgcad gfac gcb cdgabef |
	cg cg fdcagb cbg
	fbegcd cbd adcefb dageb afcb bc aefdc ecdab fgdeca fcdbega |
	efabcd cedba gadfec cb
	aecbfdg fbg gf bafeg dbefa fcge gcbea fcaegb dgceab fcbdga |
	gecf egdcabf bgf bfgea
	fgeab ca afcebg bdacfeg cfaedg gcfdb baec bfadeg bafgc acf |
	gebdcfa ecba ca fadegcb
	dbcfg fgd bdegcaf fgec aegbdf ecdfab fbedc dacgb gdcebf gf |
	cefg dcbef fcge gbcadfe
	bdfegc cbegaf gecbf dfcage bdacg ed bedf ced adcbefg gebcd |
	ed bcgafe cdgba cbgef
	egadfb cdbfeg cegd fecab cgb gbdefca cg fgcdab egfdb bfceg |
	gbdfcae bgc cg cgb
	gcafb gcf dcaebfg ecagb gf abcdeg gaef cafbge fdbac fegbdc |
	fgae cfgab fg bagce

Because the digits 1, 4, 7, and 8 each use a unique number of segments, you
should be able to tell which combinations of signals correspond to those
digits. Counting only digits in the output values (the part after | on each
line), in the above example, there are 26 instances of digits that use a unique
number of segments (highlighted above).

In the output values, how many times do digits 1, 4, 7, or 8 appear?
*/
func Part1(r io.Reader) (ans int, err error) {
	digits, err := day8(r)
	if err != nil {
		return 0, err
	}
	num1478 := 0
	for _, line := range digits {
		for _, d := range line {
			switch d {
			case 1, 4, 7, 8:
				num1478++
			}
		}
	}
	return num1478, nil
}

/*
Part2 Prompt

--- Part Two ---
Through a little deduction, you should now be able to determine the remaining
digits. Consider again the first example above:

	acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab |
	cdfeb fcadb cdfeb cdbaf

After some careful analysis, the mapping between signal wires and segments only
make sense in the following configuration:

	 dddd
	e    a
	e    a
	 ffff
	g    b
	g    b
	 cccc

So, the unique signal patterns would correspond to the following digits:

- acedgfb: 8
- cdfbe: 5
- gcdfa: 2
- fbcad: 3
- dab: 7
- cefabd: 9
- cdfgeb: 6
- eafb: 4
- cagedb: 0
- ab: 1

Then, the four digits of the output value can be decoded:

- cdfeb: 5
- fcadb: 3
- cdfeb: 5
- cdbaf: 3

Therefore, the output value for this entry is 5353.

Following this same process for each entry in the second, larger example above,
the output value of each entry can be determined:

- fdgacbe cefdb cefbgd gcbe: 8394
- fcgedb cgb dgebacf gc: 9781
- cg cg fdcagb cbg: 1197
- efabcd cedba gadfec cb: 9361
- gecf egdcabf bgf bfgea: 4873
- gebdcfa ecba ca fadegcb: 8418
- cefg dcbef fcge gbcadfe: 4548
- ed bcgafe cdgba cbgef: 1625
- gbdfcae bgc cg cgb: 8717
- fgae cfgab fg bagce: 4315

Adding all of the output values in this larger example produces 61229.

For each entry, determine all of the wire/segment connections and decode the
four-digit output values. What do you get if you add up all of the output
values?
*/
func Part2(r io.Reader) (ans int, err error) {
	digits, err := day8(r)
	if err != nil {
		return 0, err
	}
	sum := 0
	for _, line := range digits {
		n := 0
		for _, d := range line {
			n *= 10
			n += d
		}
		sum += n
	}
	return sum, nil
}

func day8(r io.Reader) ([][]int, error) {
	var out [][]int
	if err := input.ForEachLine(r, func(line string) error {
		parts := strings.SplitN(line, " | ", 2)
		inputs, outputs := parts[0], parts[1]
		out = append(out, solveLine(sortItems(strings.Split(inputs, " ")), sortItems(strings.Split(outputs, " "))))
		return nil
	}); err != nil {
		return nil, err
	}
	return out, nil
}

func solveLine(inputs, outputs []string) []int {
	options := make(map[string]map[int]bool)
	answers := make(map[int]string)
	// solve 1, 4, 7, 8 first
	for _, c := range inputs {
		switch len(c) {
		case 2:
			answers[1] = c
			options[c] = map[int]bool{1: true}
		case 3:
			answers[7] = c
			options[c] = map[int]bool{7: true}
		case 4:
			answers[4] = c
			options[c] = map[int]bool{4: true}
		case 5:
			options[c] = map[int]bool{
				2: true,
				3: true,
				5: true,
			}
		case 6:
			options[c] = map[int]bool{
				0: true,
				6: true,
				9: true,
			}
		case 7:
			answers[8] = c
			options[c] = map[int]bool{8: true}
		default:
			panic("bad input")
		}
	}

	// solve 9
	for _, c := range inputs {
		if len(options[c]) == 1 {
			continue
		}
		if contains(c, answers[4]) && contains(c, answers[7]) {
			answers[9] = c
			options[c] = map[int]bool{9: true}
		} else {
			delete(options[c], 9)
		}
	}

	// solve 5
	for _, c := range inputs {
		if len(options[c]) == 1 {
			continue
		}
		if !contains(c, answers[1]) {
			delete(options[c], 3)
			delete(options[c], 4)
			delete(options[c], 9)
		}
		if !contains(c, answers[4]) {
			delete(options[c], 9)
		}
		if !contains(c, answers[7]) {
			delete(options[c], 9)
			delete(options[c], 0)
			delete(options[c], 3)
		}
		if contains(answers[9], c) {
			delete(options[c], 6)
			delete(options[c], 2)
			delete(options[c], 0)
		}
		if len(options[c]) == 1 && options[c][5] {
			answers[5] = c
		}
	}
	if _, ok := answers[5]; !ok {
		panic("flawed")
	}

	// solve 6
	for c, o := range options {
		if c != answers[5] {
			delete(o, 5)
		}
		if len(options[c]) == 1 && options[c][6] {
			answers[6] = c
		}
	}

	if _, ok := answers[6]; !ok {
		panic("flawed")
	}

	// solve 2, 3, 0
	for c, o := range options {
		if c != answers[6] {
			delete(o, 6)
		}
		if len(options[c]) == 1 && options[c][2] {
			answers[2] = c
		}
		if len(options[c]) == 1 && options[c][3] {
			answers[3] = c
		}
		if len(options[c]) == 1 && options[c][0] {
			answers[0] = c
		}
	}

	var out []int
	for _, o := range outputs {
		out = append(out, only(options[o]))
	}
	return out
}

func contains(s, c string) bool {
	for i := range c {
		if !strings.Contains(s, c[i:i+1]) {
			return false
		}
	}
	return true
}

func sortItems(in []string) (out []string) {
	for _, s := range in {
		out = append(out, sortString(s))
	}
	return out
}

func sortString(s string) string {
	b := []byte(s)
	sort.Slice(b, func(i, j int) bool {
		return b[i] < b[j]
	})
	return string(b)
}

func only(m map[int]bool) int {
	for x := range m {
		return x
	}
	panic("bad")
}
