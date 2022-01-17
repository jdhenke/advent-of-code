package day1

import (
	"fmt"
	"io"

	"github.com/jdhenke/advent-of-code/input"
)

/*
Part1 Prompt

--- Day 1: Report Repair ---
After saving Christmas five years in a row, you've decided to take a vacation
at a nice resort on a tropical island. Surely, Christmas will go on without
you.

The tropical island has its own currency and is entirely cash-only. The gold
coins used there have a little picture of a starfish; the locals just call them
stars. None of the currency exchanges seem to have heard of them, but somehow,
you'll need to find fifty of these coins by the time you arrive so you can pay
the deposit on your room.

To save your vacation, you need to get all fifty stars by December 25th.

Collect stars by solving puzzles. Two puzzles will be made available on each
day in the Advent calendar; the second puzzle is unlocked when you complete the
first. Each puzzle grants one star. Good luck!

Before you leave, the Elves in accounting just need you to fix your expense
report (your puzzle input); apparently, something isn't quite adding up.

Specifically, they need you to find the two entries that sum to 2020 and then
multiply those two numbers together.

For example, suppose your expense report contained the following:

    1721
    979
    366
    299
    675
    1456

In this list, the two entries that sum to 2020 are 1721 and 299. Multiplying
them together produces 1721 * 299 = 514579, so the correct answer is 514579.

Of course, your expense report is much larger. Find the two entries that sum to
2020; what do you get if you multiply them together?
*/
func Part1(r io.Reader) (answer int, err error) {
	return day1(r, 2)
}

/*
Part2 Prompt

--- Part Two ---
The Elves in accounting are thankful for your help; one of them even offers you
a starfish coin they had left over from a past vacation. They offer you a
second one if you can find three numbers in your expense report that meet the
same criteria.

Using the above example again, the three entries that sum to 2020 are 979, 366,
and 675. Multiplying them together produces the answer, 241861950.

In your expense report, what is the product of the three entries that sum to
2020?
*/
func Part2(r io.Reader) (answer int, err error) {
	return day1(r, 3)
}

func day1(r io.Reader, n int) (answer int, err error) {
	var nums []int
	if err := input.ForEachInt(r, func(x int) {
		nums = append(nums, x)
	}); err != nil {
		return 0, err
	}
	combos(nums, n, 0, func(c []int) {
		if sum(c) == 2020 {
			answer = product(c)
		}
	})
	if answer == 0 {
		return 0, fmt.Errorf("failed to find answer")
	}
	return answer, nil
}

func product(c []int) int {
	p := 1
	for _, x := range c {
		p *= x
	}
	return p
}

func sum(c []int) int {
	s := 0
	for _, x := range c {
		s += x
	}
	return s
}

// Put all combos of size n in [b:] at [b:]
func combos(nums []int, n, b int, f func(c []int)) {
	if b+n > len(nums) {
		return
	}
	if n == 0 {
		f(nums[:b])
		return
	}

	// all combos with the first element
	combos(nums, n-1, b+1, f)

	// all combos without first element
	nums[len(nums)-1], nums[b] = nums[b], nums[len(nums)-1]
	combos(nums[:len(nums)-1], n, b, f)
}
