package day6

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Guard struct {
	x  int
	y  int
	dx int
	dy int
}

func (g *Guard) will_be_in_bounds(grid [][]int) bool {
	next_x, next_y := g.x+g.dx, g.y+g.dy
	return next_y >= 0 && next_y < len(grid) && next_x >= 0 && next_x < len(grid[g.y])
}
func (g Guard) is_equal(other Guard) bool {
	return g.x == other.x && g.y == other.y && g.dx == other.dx && g.dy == other.dy
}

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

	grid := make([][]int, 0)
	guard := Guard{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		row := make([]int, 0, len(text))
		for i, char := range text {
			switch char {
			case '#':
				row = append(row, 1)

			default:
				row = append(row, 0)
			}

			switch char {
			case '>':
				guard.x = i
				guard.y = len(grid)
				guard.dx = 1
			case '<':
				guard.x = i
				guard.y = len(grid)
				guard.dx = -1
			case 'v':
				guard.x = i
				guard.y = len(grid)
				guard.dy = 1
			case '^':
				guard.x = i
				guard.y = len(grid)
				guard.dy = -1
			}

		}
		grid = append(grid, row)
	}
	starting_guard := guard
	for guard.y < len(grid) && guard.y >= 0 && guard.x < len(grid[guard.y]) && guard.x >= 0 {

		grid[guard.y][guard.x] = -1
		next_x, next_y := guard.x+guard.dx, guard.y+guard.dy
		target_is_in_bounds := next_x >= 0 && next_x < len(grid[guard.y]) && next_y >= 0 && next_y < len(grid)
		if target_is_in_bounds && grid[next_y][next_x] == 1 {
			guard.dx, guard.dy = -guard.dy, guard.dx
			continue
		}
		guard.x += guard.dx
		guard.y += guard.dy

	}
	count := 0
	for _, row := range grid {
		for _, v := range row {
			if v == -1 {
				count += 1
			}
		}
	}
	cycle_obs_count := 0
	for y_i, row := range grid {
		for x_i, v := range row {
			if v == -1 {
				candidate_grid := add_obstacle(grid, x_i, y_i)
				if has_cycle(starting_guard, candidate_grid) {
					cycle_obs_count += 1
				}
			}
		}
	}
	fmt.Println(count)
	fmt.Println(cycle_obs_count)
	return nil
}

func has_cycle(guard Guard, grid [][]int) bool {
	seen := make([]Guard, 0)
	for guard.y < len(grid) && guard.y >= 0 && guard.x < len(grid[guard.y]) && guard.x >= 0 {
		for _, s := range seen {
			if guard.is_equal(s) {
				return true
			}
		}
		seen_pos := guard
		seen = append(seen, seen_pos)
		grid[guard.y][guard.x] = -1
		next_x, next_y := guard.x+guard.dx, guard.y+guard.dy
		target_is_in_bounds := next_x >= 0 && next_x < len(grid[guard.y]) && next_y >= 0 && next_y < len(grid)
		if target_is_in_bounds && grid[next_y][next_x] == 1 {
			guard.dx, guard.dy = -guard.dy, guard.dx
			continue
		}
		guard.x += guard.dx
		guard.y += guard.dy

	}
	return false

}

func add_obstacle(grid [][]int, x int, y int) [][]int {
	duplicate := make([][]int, len(grid))
	for i := range grid {
		duplicate[i] = make([]int, len(grid[i]))
		copy(duplicate[i], grid[i])
	}

	duplicate[y][x] = 1
	return duplicate
}
