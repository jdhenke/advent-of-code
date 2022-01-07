package main

import (
	aoc2021day1 "advent-of-code/2021/day1"
	"advent-of-code/solution"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	var (
		year, day, part int
		file            string
	)
	flag.IntVar(&year, "year", 0, "year")
	flag.IntVar(&day, "day", 0, "day")
	flag.IntVar(&part, "part", 0, "part")
	flag.StringVar(&file, "file", "", "file")
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
		log.Fatal("Error: -file required.")
	}
	solution, ok := getSolution(year, day, part)
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
	ans, err := solution(f)
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
		{2021, 1, 1}: aoc2021day1.Part1,
		{2021, 1, 2}: aoc2021day1.Part2,
	}
	sol, ok := solutions[key{year, day, part}]
	return sol, ok
}
