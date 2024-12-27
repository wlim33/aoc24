package day15

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

type Solver struct{}

const (
	EMPTY = iota
	BOX_LEFT
	BOX_RIGHT
	WAREHOUSE
)

type Vec2 struct {
	x int
	y int
}

type Body struct {
	p Vec2
	v Vec2
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

	scanner := bufio.NewScanner(file)

	warehouse := make([][]int, 0)
	map_completed := false
	moves := make([]rune, 0)
	robot := Vec2{}
	for scanner.Scan() {
		text := scanner.Text()
		if len(text) == 0 {
			map_completed = true
			continue
		}

		if !map_completed {
			row := make([]int, 0)
			for i, r := range text {
				switch r {
				case '#':
					row = append(row, WAREHOUSE)
					row = append(row, WAREHOUSE)
				case 'O':
					row = append(row, BOX_LEFT)
					row = append(row, BOX_RIGHT)
				case '@':
					robot = Vec2{
						x: i * 2,
						y: len(warehouse),
					}
					row = append(row, EMPTY)
					row = append(row, EMPTY)
				default:
					row = append(row, EMPTY)
					row = append(row, EMPTY)
				}

			}

			warehouse = append(warehouse, row)
		} else {
			for _, r := range text {
				moves = append(moves, r)
			}
		}

	}

	render_with_robot(warehouse, robot)
	for _, move := range moves {
		robot, warehouse = update_world(robot, convert(move), warehouse)

	}

	render_with_robot(warehouse, robot)
	fmt.Println(gps_sum(warehouse))
	return nil
}

func render_with_robot(world [][]int, robot Vec2) {
	var sb strings.Builder
	for y, row := range world {
		for x, val := range row {
			if x == robot.x && y == robot.y {
				sb.WriteString("@")
				// sb.WriteString("@")
				continue
			}
			switch val {
			case BOX_LEFT:
				sb.WriteString("[")
			case BOX_RIGHT:
				sb.WriteString("]")
			case WAREHOUSE:
				sb.WriteString("#")
			default:
				sb.WriteString(".")

			}

		}
		sb.WriteString("\n")
	}

	fmt.Println(sb.String())
}

func render(world [][]int) {
	var sb strings.Builder
	for _, row := range world {
		for _, val := range row {
			switch val {
			case BOX_LEFT:
				sb.WriteString("[")
			case BOX_RIGHT:
				sb.WriteString("]")
			case WAREHOUSE:
				sb.WriteString("#")
			default:
				sb.WriteString(".")

			}

		}
		sb.WriteString("\n")
	}

	fmt.Println(sb.String())
}

func gps_sum(world [][]int) int {
	sum := 0
	for y, row := range world {
		for x, val := range row {
			if val == BOX_LEFT {
				sum += (100 * y) + x

			}
		}

	}
	return sum

}

func convert(r rune) Vec2 {

	switch r {
	case 'v':
		return Vec2{0, 1}

	case '^':

		return Vec2{0, -1}
	case '>':

		return Vec2{1, 0}
	case '<':
		return Vec2{-1, 0}
	default:
		log.Fatal("ioanerstoinarsien")
		return Vec2{}

	}

}

func (v Vec2) add(v2 Vec2) Vec2 {
	return Vec2{
		x: v.x + v2.x,
		y: v.y + v2.y,
	}

}
func (v Vec2) subtract(v2 Vec2) Vec2 {
	return Vec2{
		x: v.x - v2.x,
		y: v.y - v2.y,
	}

}

func get_point(pos Vec2, world [][]int) int {
	return world[pos.y][pos.x]
}

func update_world(pos, direction Vec2, old [][]int) (Vec2, [][]int) {
	world := clone_warehouse(old)

	immediate_next := pos.add(direction)
	if get_point(immediate_next, world) == EMPTY {
		// render(world)
		return immediate_next, world

	}
	if get_point(immediate_next, world) == WAREHOUSE {
		return pos, world

	}

	stack := make([]Vec2, 0)

	shift_points := make([]Vec2, 0)
	if direction.y == 0 {
		for distance := 1; world[pos.y][pos.x+(direction.x*distance)] == BOX_RIGHT || world[pos.y][pos.x+(direction.x*distance)] == BOX_LEFT; distance += 1 {
			shift_points = append(shift_points, Vec2{
				x: pos.x + (direction.x * distance),
				y: pos.y,
			})
		}
		if world[pos.y][pos.x+(direction.x*(1+len(shift_points)))] == WAREHOUSE {
			return pos, world
		}
	} else {

		stack = append(stack, immediate_next)
		if world[immediate_next.y][immediate_next.x] == BOX_LEFT {

			stack = append(stack, Vec2{
				x: immediate_next.x + 1,
				y: immediate_next.y,
			})
		}
		if world[immediate_next.y][immediate_next.x] == BOX_RIGHT {

			stack = append(stack, Vec2{
				x: immediate_next.x - 1,
				y: immediate_next.y,
			})
		}
	}

	for len(stack) > 0 {
		var current Vec2
		current, stack = stack[0], stack[1:]
		if world[current.y][current.x] == BOX_LEFT || world[current.y][current.x] == BOX_RIGHT {
			if slices.Contains(shift_points, current) {
				continue
			}
			shift_points = append(shift_points, current)
		}

		if world[current.y][current.x] == WAREHOUSE {
			return pos, world
		}
		if world[current.y][current.x] == EMPTY {
			continue
		}

		stack = append(stack, Vec2{
			x: current.x,
			y: current.y + direction.y,
		})
		if world[current.y+direction.y][current.x] == BOX_LEFT {
			stack = append(stack, Vec2{
				x: current.x + 1,
				y: current.y + direction.y,
			})
		}
		if world[current.y+direction.y][current.x] == BOX_RIGHT {
			stack = append(stack, Vec2{
				x: current.x - 1,
				y: current.y + direction.y,
			})
		}

	}

	for i := len(shift_points) - 1; i >= 0; i -= 1 {
		last := world[shift_points[i].y][shift_points[i].x]
		world[shift_points[i].y][shift_points[i].x] = EMPTY
		world[shift_points[i].y+direction.y][shift_points[i].x+direction.x] = last
	}

	return immediate_next, world

}

func clone_warehouse(world [][]int) [][]int {
	cloned := make([][]int, 0)
	for _, row := range world {
		cloned = append(cloned, append(row[:0:0], row...))

	}
	return cloned

}
