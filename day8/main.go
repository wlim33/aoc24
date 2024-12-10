package day8

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Solver struct{}
type GridItem struct {
	freq rune
	anti bool
}
type Coordinate struct {
	x int
	y int
}

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

	grid := make([][]GridItem, 0)
	antenna := make(map[rune][]Coordinate, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()

		row := make([]GridItem, 0, len(text))
		for i, char := range text {
			if char != rune('.') {
				if existing, ok := antenna[char]; ok {
					antenna[char] = append(existing, Coordinate{
						x: i,
						y: len(grid),
					})
				} else {
					antenna[char] = []Coordinate{{
						x: i,
						y: len(grid),
					}}
				}
			}
			row = append(row, GridItem{
				freq: char,
				anti: false,
			})

		}
		grid = append(grid, row)
	}

	// fmt.Println(antenna)

	for _, row := range grid {
		fmt.Println(to_string(row))
	}
	for _, as := range antenna {
		for i := 0; i < len(as); i += 1 {
			for j := i + 1; j < len(as); j += 1 {
				add_antis(as[i], as[j], grid)
			}
		}
	}

	anti_count := 0

	fmt.Println("")
	for _, row := range grid {
		for _, item := range row {
			if item.anti {
				anti_count += 1
			}
		}
		fmt.Println(to_string(row))
	}
	fmt.Println(anti_count)
	return nil

}

func (g GridItem) to_string() string {
	if g.freq != '.' {
		return string(g.freq)
	}
	if g.anti {
		return string('X')
	} else {

		return string('.')
	}
}

func to_string(row []GridItem) string {
	var sb strings.Builder
	for _, v := range row {
		sb.WriteString(v.to_string())

	}
	return sb.String()

}
func (c Coordinate) is_in_bounds(grid [][]GridItem) bool {
	return c.y >= 0 && c.y < len(grid) && c.x >= 0 && c.x < len(grid[c.y])

}

func (c Coordinate) subtract(s Coordinate) Coordinate {
	return Coordinate{
		x: c.x - s.x,
		y: c.y - s.y,
	}
}
func (c Coordinate) add(s Coordinate) Coordinate {
	return Coordinate{
		x: c.x + s.x,
		y: c.y + s.y,
	}
}
func (c Coordinate) equals(s Coordinate) bool {
	return c.x == s.x && c.y == s.y
}

func add_antis(a Coordinate, b Coordinate, grid [][]GridItem) {

	grid[a.y][a.x].anti = true
	grid[b.y][b.x].anti = true
	basis := Coordinate{
		x: a.x - b.x,
		y: a.y - b.y,
	}

	current := a.add(basis)
	for current.is_in_bounds(grid) {
		if current.equals(a) || current.equals(b) {
			current = current.add(basis)
			continue
		}

		grid[current.y][current.x].anti = true
		current = current.add(basis)
	}

	reverse_current := a.subtract(basis)
	for reverse_current.is_in_bounds(grid) {
		if reverse_current.equals(a) || reverse_current.equals(b) {
			reverse_current = reverse_current.subtract(basis)
			continue
		}

		grid[reverse_current.y][reverse_current.x].anti = true
		reverse_current = reverse_current.subtract(basis)
	}
}
