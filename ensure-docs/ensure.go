package main

import (
	"flag"
	"log"
	"os"

	"github.com/jdhenke/advent-of-code/ensure-docs/ensuredocs"
)

func main() {
	var (
		year, day, part int
	)
	flag.IntVar(&year, "year", 0, "year")
	flag.IntVar(&day, "day", 0, "day")
	flag.IntVar(&part, "part", 0, "part")
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
	session := os.Getenv("SESSION")
	if session == "" {
		log.Fatal("SESSION env var must be set.")
	}
	if err := ensuredocs.Ensure(year, day, part, session); err != nil {
		log.Fatal(err)
	}
	log.Println("Success")
}
