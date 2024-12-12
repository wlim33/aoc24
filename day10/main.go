package day10

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Solver struct{}

func (d *Solver) Solve(file_name string) error {
	file, err := os.Open(file_name)
	if err != nil {
		return err
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	scanner := bufio.NewScanner(file)

	grid := make([][]int, 0)
	starts := make([]Coordinate, 0)
	for scanner.Scan() {
		text := scanner.Text()

		row := make([]int, 0, len(text))
		for _, v := range text {
			num := int(v) - '0'
			if num == 0 {
				starts = append(starts, Coordinate{
					x: len(row),
					y: len(grid),
				})
			}
			row = append(row, num)
		}

		grid = append(grid, row)
	}
	count := 0
	for _, s := range starts {
		count += get_score(grid, s)
	}
	fmt.Println(count)

	total_rating := 0
	for _, s := range starts {
		total_rating += get_rating(grid, s)
	}
	fmt.Println(total_rating)
	return nil
}

type Coordinate struct {
	x int
	y int
}

func get_rating(grid [][]int, start Coordinate) int {
	stack := make([]Coordinate, 0)
	stack = append(stack, start)
	count := 0
	for len(stack) > 0 {
		var current Coordinate
		current, stack = stack[len(stack)-1], stack[:len(stack)-1]
		if grid[current.y][current.x] == 9 {
			count += 1
			continue
		}

		stack = append(stack, current.get_neighbors(grid)...)
	}
	return count

}
func get_score(grid [][]int, start Coordinate) int {
	stack := make([]Coordinate, 0)
	stack = append(stack, start)
	seen := make(map[Coordinate]struct{})
	for len(stack) > 0 {
		var current Coordinate
		current, stack = stack[len(stack)-1], stack[:len(stack)-1]
		if grid[current.y][current.x] == 9 {
			seen[current] = struct{}{}
			continue
		}

		stack = append(stack, current.get_neighbors(grid)...)
	}
	return len(seen)

}

func (c Coordinate) get_neighbors(grid [][]int) []Coordinate {
	neighbors := make([]Coordinate, 0)
	current_level := grid[c.y][c.x]
	if c.y > 0 && current_level == grid[c.y-1][c.x]-1 {
		neighbors = append(neighbors, Coordinate{
			x: c.x,
			y: c.y - 1,
		})
	}
	if c.y < len(grid)-1 && current_level == grid[c.y+1][c.x]-1 {
		neighbors = append(neighbors, Coordinate{
			x: c.x,
			y: c.y + 1,
		})
	}
	if c.x > 0 && current_level == grid[c.y][c.x-1]-1 {
		neighbors = append(neighbors, Coordinate{
			x: c.x - 1,
			y: c.y,
		})
	}
	if c.y < len(grid) && c.x < len(grid[c.y])-1 && current_level == grid[c.y][c.x+1]-1 {
		neighbors = append(neighbors, Coordinate{
			x: c.x + 1,
			y: c.y,
		})
	}
	return neighbors
}
