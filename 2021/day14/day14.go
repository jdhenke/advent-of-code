package day14

import (
	"bufio"
	"io"
	"strings"
)

/*
Part1 Prompt

--- Day 14: Extended Polymerization ---
The incredible pressures at this depth are starting to put a strain on your
submarine. The submarine has polymerization equipment that would produce
suitable materials to reinforce the submarine, and the nearby
volcanically-active caves should even have the necessary input elements in
sufficient quantities.

The submarine manual contains instructions for finding the optimal polymer
formula; specifically, it offers a polymer template and a list of pair
insertion rules (your puzzle input). You just need to work out what polymer
would result after repeating the pair insertion process a few times.

For example:

	NNCB

	CH -> B
	HH -> N
	CB -> H
	NH -> C
	HB -> C
	HC -> B
	HN -> C
	NN -> C
	BH -> H
	NC -> B
	NB -> B
	BN -> B
	BB -> N
	BC -> B
	CC -> N
	CN -> C

The first line is the polymer template - this is the starting point of the
process.

The following section defines the pair insertion rules. A rule like AB -> C
means that when elements A and B are immediately adjacent, element C should be
inserted between them. These insertions all happen simultaneously.

So, starting with the polymer template NNCB, the first step simultaneously
considers all three pairs:

- The first pair (NN) matches the rule NN -> C, so element C is inserted
between the first N and the second N.
- The second pair (NC) matches the rule NC -> B, so element B is inserted
between the N and the C.
- The third pair (CB) matches the rule CB -> H, so element H is inserted
between the C and the B.

Note that these pairs overlap: the second element of one pair is the first
element of the next pair. Also, because all pairs are considered
simultaneously, inserted elements are not considered to be part of a pair until
the next step.

After the first step of this process, the polymer becomes NCNBCHB.

Here are the results of a few steps using the above rules:

	Template:     NNCB
	After step 1: NCNBCHB
	After step 2: NBCCNBBBCBHCB
	After step 3: NBBBCNCCNBBNBNBBCHBHHBCHB
	After step 4: NBBNBNBBCCNBCNCCNBBNBBNBBBNBBNBBCBHCBHHNHCBBCBHCB

This polymer grows quickly. After step 5, it has length 97; After step 10, it
has length 3073. After step 10, B occurs 1749 times, C occurs 298 times, H
occurs 161 times, and N occurs 865 times; taking the quantity of the most
common element (B, 1749) and subtracting the quantity of the least common
element (H, 161) produces 1749 - 161 = 1588.

Apply 10 steps of pair insertion to the polymer template and find the most and
least common elements in the result. What do you get if you take the quantity
of the most common element and subtract the quantity of the least common
element?
*/
func Part1(r io.Reader) (ans int, err error) {
	return day14(r, 10)
}

/*
Part2 Prompt

--- Part Two ---
The resulting polymer isn't nearly strong enough to reinforce the submarine.
You'll need to run more steps of the pair insertion process; a total of 40
steps should do it.

In the above example, the most common element is B (occurring 2192039569602
times) and the least common element is H (occurring 3849876073 times);
subtracting these produces 2188189693529.

Apply 40 steps of pair insertion to the polymer template and find the most and
least common elements in the result. What do you get if you take the quantity
of the most common element and subtract the quantity of the least common
element?
*/
func Part2(r io.Reader) (ans int, err error) {
	return day14(r, 40)
}

func day14(r io.Reader, steps int) (ans int, err error) {
	scan := bufio.NewScanner(r)
	scan.Scan()
	text := scan.Text()
	scan.Scan()
	rules := make(map[string]string)
	for scan.Scan() {
		parts := strings.Split(scan.Text(), " -> ")
		rules[parts[0]] = parts[1]
	}
	if err := scan.Err(); err != nil {
		return 0, err
	}

	type entry struct {
		text   string
		rounds int
	}
	memo := make(map[entry]map[string]int)
	var getCounts func(text string, rounds int) (counts map[string]int)

	// only gets the counts of all new characters introduced
	getCounts = func(text string, rounds int) (counts map[string]int) {
		e := entry{
			text:   text,
			rounds: rounds,
		}
		if counts, ok := memo[e]; ok {
			return counts
		}
		defer func() {
			memo[e] = counts
		}()
		if rounds == 0 {
			return nil
		}
		counts = map[string]int{}
		for i := 0; i+1 < len(text); i++ {
			if m, ok := rules[text[i:i+2]]; ok {
				counts[m]++
				for c, v := range getCounts(text[i:i+1]+m, rounds-1) {
					counts[c] += v
				}
				for c, v := range getCounts(m+text[i+1:i+2], rounds-1) {
					counts[c] += v
				}
			}
		}
		return counts
	}
	freqs := getCounts(text, steps)

	// update to include the characters in the original text as well
	for i := 0; i < len(text); i++ {
		freqs[text[i:i+1]]++
	}

	var min, max string
	for c := range freqs {
		if min == "" || freqs[c] < freqs[min] {
			min = c
		}
		if max == "" || freqs[c] > freqs[max] {
			max = c
		}
	}
	return freqs[max] - freqs[min], nil
}
