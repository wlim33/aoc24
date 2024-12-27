package main

import (
	"aoc24/day1"
	"aoc24/day10"
	"aoc24/day11"
	"aoc24/day12"
	"aoc24/day13"
	"aoc24/day14"
	"aoc24/day15"
	"aoc24/day2"
	"aoc24/day3"
	"aoc24/day4"
	"aoc24/day5"
	"aoc24/day6"
	"aoc24/day7"
	"aoc24/day8"
	"aoc24/day9"
	"flag"
	"fmt"
	"log"
)

type Solver interface {
	Solve(string) error
}

const SmallTestDataPath = "small.txt"
const LargeTestDataPath = "full.txt"

func input_path(day int, use_test_data bool) string {
	if !use_test_data {
		return fmt.Sprintf("day%d/%s", day, SmallTestDataPath)
	} else {
		return fmt.Sprintf("day%d/%s", day, LargeTestDataPath)

	}
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
		day_solver = new(day1.Solver)
	case 2:
		day_solver = new(day2.Solver)
	case 3:
		day_solver = new(day3.Solver)
	case 4:
		day_solver = new(day4.Solver)
	case 5:
		day_solver = new(day5.Solver)
	case 6:
		day_solver = new(day6.Solver)
	case 7:
		day_solver = new(day7.Solver)
	case 8:
		day_solver = new(day8.Solver)
	case 9:
		day_solver = new(day9.Solver)
	case 10:
		day_solver = new(day10.Solver)
	case 11:
		day_solver = new(day11.Solver)
	case 12:
		day_solver = new(day12.Solver)
	case 13:
		day_solver = new(day13.Solver)
	case 14:
		day_solver = new(day14.Solver)
	case 15:
		day_solver = new(day15.Solver)
	default:
		log.Fatalf("Selected day not implemented/released")
	}

	if use_test_data {
		fmt.Printf("Selected day %d\n", day_flag)
	} else {
		fmt.Printf("Selected day %d(with example data)\n", day_flag)
	}
	path := input_path(day_flag, use_test_data)

	err := day_solver.Solve(path)
	if err != nil {
		log.Fatalf("Error trying to solve day %d:\n%s", day_flag, err)
	}
}
