package day12

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
)

type Solver struct{}

type Plant = rune
type Coordinate struct {
	x int
	y int
}

type Island = map[Coordinate]struct{}

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
	plot := make([][]Plant, 0)
	for scanner.Scan() {
		text := scanner.Text()

		plot = append(plot, []rune(text))

	}

	regions := make(map[Plant][]Island)

	for y := 0; y < len(plot); y += 1 {
		for x := 0; x < len(plot[y]); x += 1 {
			c := Coordinate{
				x, y,
			}
			if islands, plant_found := regions[plot[y][x]]; plant_found {
				coordinate_visited := false
				for _, island := range islands {
					if _, ok := island[c]; ok {
						coordinate_visited = true
						break
					}
				}
				if !coordinate_visited {
					island := get_contiguous(c, plot)
					regions[plot[y][x]] = append(regions[plot[y][x]], island)
				}
			} else {
				island := get_contiguous(c, plot)
				regions[plot[y][x]] = []map[Coordinate]struct{}{island}
			}

		}
	}

	price := 0
	for _, islands := range regions {
		for _, island := range islands {
			area := len(island)
			perimeter := get_perimeter(island)
			// fmt.Println("perimeter, edge_count", perimeter, get_edge_count(island))
			price += (area * perimeter)
		}

	}

	fmt.Println("total price pt 1: ", price)

	price_b := 0
	for _, islands := range regions {
		for _, island := range islands {
			area := len(island)
			edges := get_edge_count(island)
			price_b += (area * edges)
		}

	}
	fmt.Println("total price pt 2: ", price_b)
	return nil
}

const (
	NORTH = iota
	SOUTH
	EAST
	WEST
)

type Edge struct {
	island_boundary Coordinate
	direction       int
}

func get_edges(island Island) map[Edge]struct{} {
	edge_points := make(map[Edge]struct{})

	for point := range island {
		c_east := Coordinate{point.x + 1, point.y}
		if _, ok := island[c_east]; !ok {
			edge_points[Edge{
				c_east,
				EAST,
			}] = struct{}{}
		}

		c_west := Coordinate{point.x - 1, point.y}
		if _, ok := island[c_west]; !ok {
			edge_points[Edge{
				c_west,
				WEST,
			}] = struct{}{}
		}
		c_south := Coordinate{point.x, point.y + 1}
		if _, ok := island[c_south]; !ok {
			edge_points[Edge{
				c_south,
				SOUTH,
			}] = struct{}{}
		}
		c_north := Coordinate{point.x, point.y - 1}
		if _, ok := island[c_north]; !ok {
			edge_points[Edge{
				c_north,
				NORTH,
			}] = struct{}{}
		}

	}

	return edge_points

}

func all_boundaries_on_edge(edges map[Edge]struct{}, start Edge) []Edge {
	neighbors := []Edge{start}

	if start.direction == NORTH || start.direction == SOUTH {
		left, right := start, start
		left_exists := true
		for left_exists {

			if !slices.Contains(neighbors, left) {
				neighbors = append(neighbors, left)
			}
			left = Edge{
				Coordinate{
					left.island_boundary.x - 1, left.island_boundary.y,
				},
				left.direction,
			}

			_, left_exists = edges[left]
		}
		right_exists := true
		for right_exists {
			if !slices.Contains(neighbors, right) {
				neighbors = append(neighbors, right)
			}
			right = Edge{
				Coordinate{
					right.island_boundary.x + 1, right.island_boundary.y,
				},
				right.direction,
			}

			_, right_exists = edges[right]

		}

	}
	if start.direction == EAST || start.direction == WEST {
		left, right := start, start
		left_exists := true
		for left_exists {
			if !slices.Contains(neighbors, left) {
				neighbors = append(neighbors, left)
			}
			left = Edge{
				Coordinate{
					left.island_boundary.x, left.island_boundary.y - 1,
				},
				left.direction,
			}

			_, left_exists = edges[left]

		}
		right_exists := true
		for right_exists {
			if !slices.Contains(neighbors, right) {
				neighbors = append(neighbors, right)
			}
			right = Edge{
				Coordinate{
					right.island_boundary.x, right.island_boundary.y + 1,
				},
				right.direction,
			}

			_, right_exists = edges[right]

		}

	}

	return neighbors
}

func get_edge_count(island Island) int {
	edges_count := 0
	edges := get_edges(island)
	seen := make(map[Edge]struct{})

	for e := range edges {
		if _, ok := seen[e]; ok {
			continue
		}
		edges_count += 1
		for _, same_edge := range all_boundaries_on_edge(edges, e) {
			seen[same_edge] = struct{}{}

		}

	}
	return edges_count
}

func get_perimeter(island map[Coordinate]struct{}) int {
	perimeter := 0

	for point := range island {
		if _, ok := island[Coordinate{point.x + 1, point.y}]; !ok {
			perimeter += 1
		}

		if _, ok := island[Coordinate{point.x - 1, point.y}]; !ok {
			perimeter += 1
		}
		if _, ok := island[Coordinate{point.x, point.y + 1}]; !ok {
			perimeter += 1
		}
		if _, ok := island[Coordinate{point.x, point.y - 1}]; !ok {
			perimeter += 1
		}
	}
	return perimeter

}

func get_contiguous(start Coordinate, plot [][]Plant) map[Coordinate]struct{} {
	visited := make(map[Coordinate]struct{})

	stack := make([]Coordinate, 0)
	stack = append(stack, start)
	for len(stack) > 0 {
		var current Coordinate
		current, stack = stack[len(stack)-1], stack[:len(stack)-1]
		visited[current] = struct{}{}

		for _, neighbor := range current.get_neighbors(plot, visited) {
			stack = append(stack, neighbor)
		}

	}
	return visited
}

func (c Coordinate) get_neighbors(grid [][]Plant, visited map[Coordinate]struct{}) []Coordinate {
	neighbors := make([]Coordinate, 0)
	current_plant := grid[c.y][c.x]

	if c.y > 0 && current_plant == grid[c.y-1][c.x] {
		candidate := Coordinate{
			x: c.x,
			y: c.y - 1,
		}
		if _, ok := visited[candidate]; !ok {
			neighbors = append(neighbors, candidate)
		}
	}
	if c.y < len(grid)-1 && current_plant == grid[c.y+1][c.x] {
		candidate := Coordinate{
			x: c.x,
			y: c.y + 1,
		}
		if _, ok := visited[candidate]; !ok {
			neighbors = append(neighbors, candidate)
		}
	}
	if c.x > 0 && current_plant == grid[c.y][c.x-1] {
		candidate := Coordinate{
			x: c.x - 1,
			y: c.y,
		}
		if _, ok := visited[candidate]; !ok {
			neighbors = append(neighbors, candidate)
		}
	}
	if c.y < len(grid) && c.x < len(grid[c.y])-1 && current_plant == grid[c.y][c.x+1] {
		candidate := Coordinate{
			x: c.x + 1,
			y: c.y,
		}
		if _, ok := visited[candidate]; !ok {
			neighbors = append(neighbors, candidate)
		}
	}
	return neighbors
}
