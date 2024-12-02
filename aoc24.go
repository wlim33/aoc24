package main

import (
	"aoc24/day1"
	"flag"
	"fmt"
	"log"
)

type Solver interface {
	Solve() error
	LoadData(bool) error
}

func main() {
	var day_flag int
	flag.IntVar(&day_flag, "d", 1, "Which day to run.")
	var use_test_data bool
	flag.BoolVar(&use_test_data, "t", false, "Use actual inputs instead of shared examples.")
	flag.Parse()

	var day_solver Solver
	switch day_flag {
	case 1:
		day_solver = new(day1.Day1Solver)
	default:
		fmt.Println("Selected day not implemented/released")
	}

	if use_test_data {
		fmt.Printf("Selected day %d\n", day_flag)
	} else {
		fmt.Printf("Selected day %d(with example data)\n", day_flag)
	}
	err := day_solver.LoadData(use_test_data)
	if err != nil {
		log.Fatalf("Error trying to load data for day %d:\n%s", day_flag, err)
	}

	err = day_solver.Solve()
	if err != nil {
		log.Fatalf("Error trying to solve day %d:\n%s", day_flag, err)
	}
}
