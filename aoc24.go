package main

import (
	"aoc24/day1"
	"flag"
	"fmt"
	"log"
)

type Solver interface {
	Solve() error
}

func main() {
	var day_flag int
	flag.IntVar(&day_flag, "day", 1, "which day to run")
	var input_size bool
	flag.BoolVar(&input_size, "input_size", false, "which test data to use")
	flag.Parse()

	var day_solver Solver
	switch day_flag {
	case 1:
		fmt.Println("Selected day 1")
		day_solver = &day1.Day1Solver{}
	default:
		fmt.Println("Selected day not implemented/released")
	}

	err := day_solver.Solve()
	if err != nil {
		log.Fatalf("Error trying to solve day %d:\n%s", day_flag, err)
	}
}
