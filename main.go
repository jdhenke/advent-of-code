package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	aoc2020day1 "github.com/jdhenke/advent-of-code/2020/day1"
	aoc2020day2 "github.com/jdhenke/advent-of-code/2020/day2"
	aoc2020day3 "github.com/jdhenke/advent-of-code/2020/day3"
	aoc2020day4 "github.com/jdhenke/advent-of-code/2020/day4"
	aoc2020day5 "github.com/jdhenke/advent-of-code/2020/day5"
	aoc2020day6 "github.com/jdhenke/advent-of-code/2020/day6"
	aoc2020day7 "github.com/jdhenke/advent-of-code/2020/day7"
	aoc2020day8 "github.com/jdhenke/advent-of-code/2020/day8"
	aoc2020day9 "github.com/jdhenke/advent-of-code/2020/day9"
	aoc2021day1 "github.com/jdhenke/advent-of-code/2021/day1"
	aoc2021day10 "github.com/jdhenke/advent-of-code/2021/day10"
	aoc2021day11 "github.com/jdhenke/advent-of-code/2021/day11"
	aoc2021day12 "github.com/jdhenke/advent-of-code/2021/day12"
	aoc2021day13 "github.com/jdhenke/advent-of-code/2021/day13"
	aoc2021day14 "github.com/jdhenke/advent-of-code/2021/day14"
	aoc2021day15 "github.com/jdhenke/advent-of-code/2021/day15"
	aoc2021day16 "github.com/jdhenke/advent-of-code/2021/day16"
	aoc2021day17 "github.com/jdhenke/advent-of-code/2021/day17"
	aoc2021day18 "github.com/jdhenke/advent-of-code/2021/day18"
	aoc2021day19 "github.com/jdhenke/advent-of-code/2021/day19"
	aoc2021day2 "github.com/jdhenke/advent-of-code/2021/day2"
	aoc2021day20 "github.com/jdhenke/advent-of-code/2021/day20"
	aoc2021day21 "github.com/jdhenke/advent-of-code/2021/day21"
	aoc2021day22 "github.com/jdhenke/advent-of-code/2021/day22"
	aoc2021day23 "github.com/jdhenke/advent-of-code/2021/day23"
	aoc2021day24 "github.com/jdhenke/advent-of-code/2021/day24"
	aoc2021day25 "github.com/jdhenke/advent-of-code/2021/day25"
	aoc2021day3 "github.com/jdhenke/advent-of-code/2021/day3"
	aoc2021day4 "github.com/jdhenke/advent-of-code/2021/day4"
	aoc2021day5 "github.com/jdhenke/advent-of-code/2021/day5"
	aoc2021day6 "github.com/jdhenke/advent-of-code/2021/day6"
	aoc2021day7 "github.com/jdhenke/advent-of-code/2021/day7"
	aoc2021day8 "github.com/jdhenke/advent-of-code/2021/day8"
	aoc2021day9 "github.com/jdhenke/advent-of-code/2021/day9"
	"github.com/jdhenke/advent-of-code/solution"
)

func main() {
	var (
		year, day, part int
		file            string
	)
	flag.IntVar(&year, "year", 0, "year")
	flag.IntVar(&day, "day", 0, "day")
	flag.IntVar(&part, "part", 0, "part")
	flag.StringVar(&file, "file", "", "file (defaults to checked input file)")
	flag.Parse()
	if year == 0 {
		log.Fatal("Error: -year required.")
	}
	if day == 0 {
		log.Fatal("Error: -day required.")
	}
	if part == 0 {
		log.Fatal("Error: -part required.")
	}
	if file == "" {
		file = filepath.Join(fmt.Sprint(year), fmt.Sprintf("day%d", day), "input.txt")
	}
	fn, ok := getSolution(year, day, part)
	if !ok {
		log.Fatalf("Error: No solution present for %v %v %v.", year, day, part)
	}
	f, err := os.Open(file)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Printf("Error closing file: %v", err)
		}
	}()
	ans, err := fn(f)
	if err != nil {
		log.Fatalf("Error running solution: %v", err)
	}
	fmt.Println(ans)
}

type key struct {
	year, day, part int
}

func getSolution(year, day, part int) (solution.Func, bool) {
	solutions := map[key]solution.Func{
		{2020, 1, 1}:  aoc2020day1.Part1,
		{2020, 1, 2}:  aoc2020day1.Part2,
		{2020, 2, 1}:  aoc2020day2.Part1,
		{2020, 2, 2}:  aoc2020day2.Part2,
		{2020, 3, 1}:  aoc2020day3.Part1,
		{2020, 3, 2}:  aoc2020day3.Part2,
		{2020, 4, 1}:  aoc2020day4.Part1,
		{2020, 4, 2}:  aoc2020day4.Part2,
		{2020, 5, 1}:  aoc2020day5.Part1,
		{2020, 5, 2}:  aoc2020day5.Part2,
		{2020, 6, 1}:  aoc2020day6.Part1,
		{2020, 6, 2}:  aoc2020day6.Part2,
		{2020, 7, 1}:  aoc2020day7.Part1,
		{2020, 7, 2}:  aoc2020day7.Part2,
		{2020, 8, 1}:  aoc2020day8.Part1,
		{2020, 8, 2}:  aoc2020day8.Part2,
		{2020, 9, 1}:  aoc2020day9.Part1,
		{2020, 9, 2}:  aoc2020day9.Part2,
		{2021, 1, 1}:  aoc2021day1.Part1,
		{2021, 1, 2}:  aoc2021day1.Part2,
		{2021, 2, 1}:  aoc2021day2.Part1,
		{2021, 2, 2}:  aoc2021day2.Part2,
		{2021, 3, 1}:  aoc2021day3.Part1,
		{2021, 3, 2}:  aoc2021day3.Part2,
		{2021, 4, 1}:  aoc2021day4.Part1,
		{2021, 4, 2}:  aoc2021day4.Part2,
		{2021, 5, 1}:  aoc2021day5.Part1,
		{2021, 5, 2}:  aoc2021day5.Part2,
		{2021, 6, 1}:  aoc2021day6.Part1,
		{2021, 6, 2}:  aoc2021day6.Part2,
		{2021, 7, 1}:  aoc2021day7.Part1,
		{2021, 7, 2}:  aoc2021day7.Part2,
		{2021, 8, 1}:  aoc2021day8.Part1,
		{2021, 8, 2}:  aoc2021day8.Part2,
		{2021, 9, 1}:  aoc2021day9.Part1,
		{2021, 9, 2}:  aoc2021day9.Part2,
		{2021, 10, 1}: aoc2021day10.Part1,
		{2021, 10, 2}: aoc2021day10.Part2,
		{2021, 11, 1}: aoc2021day11.Part1,
		{2021, 11, 2}: aoc2021day11.Part2,
		{2021, 12, 1}: aoc2021day12.Part1,
		{2021, 12, 2}: aoc2021day12.Part2,
		{2021, 13, 1}: aoc2021day13.Part1,
		{2021, 13, 2}: aoc2021day13.Part2,
		{2021, 14, 1}: aoc2021day14.Part1,
		{2021, 14, 2}: aoc2021day14.Part2,
		{2021, 15, 1}: aoc2021day15.Part1,
		{2021, 15, 2}: aoc2021day15.Part2,
		{2021, 16, 1}: aoc2021day16.Part1,
		{2021, 16, 2}: aoc2021day16.Part2,
		{2021, 17, 1}: aoc2021day17.Part1,
		{2021, 17, 2}: aoc2021day17.Part2,
		{2021, 18, 1}: aoc2021day18.Part1,
		{2021, 18, 2}: aoc2021day18.Part2,
		{2021, 19, 1}: aoc2021day19.Part1,
		{2021, 19, 2}: aoc2021day19.Part2,
		{2021, 20, 1}: aoc2021day20.Part1,
		{2021, 20, 2}: aoc2021day20.Part2,
		{2021, 21, 1}: aoc2021day21.Part1,
		{2021, 21, 2}: aoc2021day21.Part2,
		{2021, 22, 1}: aoc2021day22.Part1,
		{2021, 22, 2}: aoc2021day22.Part2,
		{2021, 23, 1}: aoc2021day23.Part1,
		{2021, 23, 2}: aoc2021day23.Part2,
		{2021, 24, 1}: aoc2021day24.Part1,
		{2021, 24, 2}: aoc2021day24.Part2,
		{2021, 25, 1}: aoc2021day25.Part1,
		{2021, 25, 2}: aoc2021day25.Part2,
	}
	sol, ok := solutions[key{year, day, part}]
	return sol, ok
}
