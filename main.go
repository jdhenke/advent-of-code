package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("input-1a.txt")
	if err != nil {
		log.Fatal(err)
	}
	s := bufio.NewScanner(f)
	var last = math.MaxInt
	var increases int
	for s.Scan() {
		text := s.Text()
		n, err := strconv.Atoi(text)
		if err != nil {
			log.Fatal(err)
		}
		if n > last {
			increases++
		}
		last = n
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(increases, "Increases")
}
