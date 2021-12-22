package main

import (
	"flag"
	"fmt"

	"github.com/damon-myers/aoc-2021/day1"
	"github.com/damon-myers/aoc-2021/day2"
	"github.com/damon-myers/aoc-2021/day3"
	"github.com/damon-myers/aoc-2021/day4"
)

func main() {
	var days = map[int]func(string){
		1: day1.Run,
		2: day2.Run,
		3: day3.Run,
		4: day4.Run,
	}

	infilePtr := flag.String("infile", "", "path to the input file")
	dayPtr := flag.Int("d", 1, "which day to run")
	flag.Parse()

	if *infilePtr == "" {
		*infilePtr = fmt.Sprintf("inputs/day%d.txt", *dayPtr)
	}

	days[*dayPtr](*infilePtr)
}
