package main

import (
	"advent-of-code/day1"
	"advent-of-code/day10"
	"advent-of-code/day11"
	"advent-of-code/day12"
	"advent-of-code/day2"
	"advent-of-code/day3"
	"advent-of-code/day4"
	"advent-of-code/day5"
	"advent-of-code/day6"
	"advent-of-code/day7"
	"advent-of-code/day8"
	"advent-of-code/day9"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	var (
		day  = flag.Int("day", 0, "day to solve")
		part = flag.Int("part", 0, "part to solve")
	)
	flag.Parse()
	if *day == 0 || *part == 0 {
		log.Fatalf("Must provide -day and -part.")
	}
	type entry struct {
		day, part int
	}
	solvers := map[entry]func(r io.Reader) (answer int, err error){
		{1, 1}:  day1.Part1,
		{1, 2}:  day1.Part2,
		{2, 1}:  day2.Part1,
		{2, 2}:  day2.Part2,
		{3, 1}:  day3.Part1,
		{3, 2}:  day3.Part2,
		{4, 1}:  day4.Part1,
		{4, 2}:  day4.Part2,
		{5, 1}:  day5.Part1,
		{5, 2}:  day5.Part2,
		{6, 1}:  day6.Part1,
		{6, 2}:  day6.Part2,
		{7, 1}:  day7.Part1,
		{7, 2}:  day7.Part2,
		{8, 1}:  day8.Part1,
		{8, 2}:  day8.Part2,
		{9, 1}:  day9.Part1,
		{9, 2}:  day9.Part2,
		{10, 1}: day10.Part1,
		{10, 2}: day10.Part2,
		{11, 1}: day11.Part1,
		{11, 2}: day11.Part2,
		{12, 1}: day12.Part1,
		{12, 2}: day12.Part2,
	}
	solverFunc, ok := solvers[entry{*day, *part}]
	if !ok {
		log.Fatalf("No solution for day %d part %d.", *day, *part)
	}
	f, err := os.Open(fmt.Sprintf("data/day%d.txt", *day))
	if err != nil {
		log.Fatalf("Failed to open input for %d: %v", *day, err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Printf("Failed to close input file: %v", err)
		}
	}()
	ans, err := solverFunc(f)
	if err != nil {
		log.Fatalf("Solver returned an error: %v", err)
	}
	fmt.Println(ans)
}
