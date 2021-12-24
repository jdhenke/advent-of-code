package main

import (
	"advent-of-code/day1"
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
	type entry struct {
		day, part int
	}
	solvers := map[entry]func(r io.Reader) (answer string, err error){
		{1, 2}: day1.Part2,
	}
	solverFunc, ok := solvers[entry{*day, *part}]
	if !ok {
		log.Fatalf("No solution for day %d part %d.", *day, part)
	}
	f, err := os.Open(fmt.Sprintf("inputs/day%d.txt", *day))
	if err != nil {
		log.Fatalf("No input for day %d.", *day)
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
