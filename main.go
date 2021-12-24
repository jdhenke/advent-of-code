package main

import (
	"advent-of-code/day1"
	"advent-of-code/day2"
	"advent-of-code/day3"
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
	solvers := map[entry]func(r io.Reader) (answer string, err error){
		{1, 1}: day1.Part1,
		{1, 2}: day1.Part2,
		{2, 1}: day2.Part1,
		{2, 2}: day2.Part2,
		{3, 1}: day3.Part1,
		{3, 2}: day3.Part2,
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
