package day9

import (
	"fmt"
	"io"
	"sort"

	"github.com/jdhenke/advent-of-code/input"
)

/*
Part1 Prompt

--- Day 9: Encoding Error ---
With your neighbor happily enjoying their video game, you turn your attention
to an open data port on the little screen in the seat in front of you.

Though the port is non-standard, you manage to connect it to your computer
through the clever use of several paperclips. Upon connection, the port outputs
a series of numbers (your puzzle input).

The data appears to be encrypted with the eXchange-Masking Addition System
(XMAS) which, conveniently for you, is an old cypher with an important
weakness.

XMAS starts by transmitting a preamble of 25 numbers. After that, each number
you receive should be the sum of any two of the 25 immediately previous
numbers. The two numbers will have different values, and there might be more
than one such pair.

For example, suppose your preamble consists of the numbers 1 through 25 in a
random order. To be valid, the next number must be the sum of two of those
numbers:

- 26 would be a valid next number, as it could be 1 plus 25 (or many other
pairs, like 2 and 24).
- 49 would be a valid next number, as it is the sum of 24 and 25.
- 100 would not be valid; no two of the previous 25 numbers sum to 100.
- 50 would also not be valid; although 25 appears in the previous 25 numbers,
the two numbers in the pair must be different.

Suppose the 26th number is 45, and the first number (no longer an option, as it
is more than 25 numbers ago) was 20. Now, for the next number to be valid,
there needs to be some pair of numbers among 1-19, 21-25, or 45 that add up to
it:

- 26 would still be a valid next number, as 1 and 25 are still within the
previous 25 numbers.
- 65 would not be valid, as no two of the available numbers sum to it.
- 64 and 66 would both be valid, as they are the result of 19+45 and 21+45
respectively.

Here is a larger example which only considers the previous 5 numbers (and has a
preamble of length 5):

	35
	20
	15
	25
	47
	40
	62
	55
	65
	95
	102
	117
	150
	182
	127
	219
	299
	277
	309
	576

In this example, after the 5-number preamble, almost every number is the sum of
two of the previous 5 numbers; the only number that does not follow this rule
is 127.

The first step of attacking the weakness in the XMAS data is to find the first
number in the list (after the preamble) which is not the sum of two of the 25
numbers before it. What is the first number that does not have this property?
*/
func Part1(r io.Reader) (answer int, err error) {
	return Part1WithN(r, 25)
}

func Part1WithN(r io.Reader, n int) (answer int, err error) {
	var nums []int
	if err := input.ForEachInt(r, func(x int) {
		nums = append(nums, x)
	}); err != nil {
		return 0, err
	}
	return part1FromNumsWithN(nums, n)
}

func part1FromNumsWithN(nums []int, n int) (int, error) {
	rs := New(nums, n)
	for i := n; i < len(nums); i++ {
		if !rs.Contains(nums[i]) {
			return nums[i], nil
		}
		rs.Next()
	}
	return 0, fmt.Errorf("no solution found")
}

type RollingSums struct {
	nums   []int
	n      int
	tail   int
	lookup map[int]int
}

func New(nums []int, n int) *RollingSums {
	lookup := make(map[int]int)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			lookup[nums[i]+nums[j]]++
		}
	}
	return &RollingSums{
		nums:   nums,
		n:      n,
		tail:   0,
		lookup: lookup,
	}
}

func (rs *RollingSums) Contains(x int) bool {
	return rs.lookup[x] > 0
}

func (rs *RollingSums) Next() {
	for i := rs.tail + 1; i < rs.tail+rs.n; i++ {
		tail := rs.nums[rs.tail]
		val := rs.nums[i]
		head := rs.nums[rs.tail+rs.n]
		rs.lookup[tail+val]--
		rs.lookup[val+head]++
	}
	rs.tail++
}

/*
Part2 Prompt

--- Part Two ---
The final step in breaking the XMAS encryption relies on the invalid number you
just found: you must find a contiguous set of at least two numbers in your list
which sum to the invalid number from step 1.

Again consider the above example:

	35
	20
	15
	25
	47
	40
	62
	55
	65
	95
	102
	117
	150
	182
	127
	219
	299
	277
	309
	576

In this list, adding up all of the numbers from 15 through 40 produces the
invalid number from step 1, 127. (Of course, the contiguous set of numbers in
your actual list might be much longer.)

To find the encryption weakness, add together the smallest and largest number
in this contiguous range; in this example, these are 15 and 47, producing 62.

What is the encryption weakness in your XMAS-encrypted list of numbers?
*/
func Part2(r io.Reader) (answer int, err error) {
	return Part2WithN(r, 25)
}

func Part2WithN(r io.Reader, n int) (answer int, err error) {
	var nums []int
	if err := input.ForEachInt(r, func(x int) {
		nums = append(nums, x)
	}); err != nil {
		return 0, err
	}
	target, err := part1FromNumsWithN(nums, n)
	if err != nil {
		return 0, err
	}

	// find a contiguous set of at least two numbers that sum to target
	tail, head := 0, 1
	sum := nums[tail] + nums[head]
	for {
		if sum == target {
			region := nums[tail : head+1]
			sort.Ints(region)
			return region[0] + region[len(region)-1], nil
		}
		if sum < target {
			head++
			sum += nums[head]
		} else if sum > target {
			tail++
			head = tail + 1
			sum = nums[tail] + nums[head]
		}
	}
}
