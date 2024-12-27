package day14

import (
	"bufio"
	"cmp"
	"fmt"
	"log"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

type Solver struct{}
type Vec2 struct {
	x int
	y int
}

type Robot struct {
	p Vec2
	v Vec2
}
type World struct {
	bounds Vec2
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

	robots := make([]Robot, 0)
	for scanner.Scan() {
		text := scanner.Text()
		robot, parse_err := parse_robot(text)
		if parse_err != nil {
			return parse_err
		}
		robots = append(robots, robot)
	}
	world := World{
		bounds: Vec2{
			x: 101,
			y: 103,
		},
	}

	const tick_count = 8000

	all_scores := make(map[int]int)
	sorted_scores := make([]int, 0, tick_count)
	for i := 0; i < tick_count; i += 1 {

		current_score := neighbor_score(world.update_tick(robots, i))
		all_scores[current_score] = i
		sorted_scores = Insert(sorted_scores, current_score)
	}

	world.draw(world.update_tick(robots, all_scores[sorted_scores[len(sorted_scores)-1]]))

	fmt.Println(all_scores[sorted_scores[len(sorted_scores)-1]])
	// quadrants_count := []int{0, 0, 0, 0}
	// fmt.Println(robots)
	// for _, robot := range robots {
	// 	q_i := world.get_quadrant_index(robot)
	// 	if q_i < 0 {
	// 		continue
	// 	}
	// 	quadrants_count[q_i] += 1
	// }

	// safety_factor := quadrants_count[0] * quadrants_count[1] * quadrants_count[2] * quadrants_count[3]
	// fmt.Println(quadrants_count)
	// fmt.Println(safety_factor)
	return nil
}

func Insert[T cmp.Ordered](ts []T, t T) []T {
	i, _ := slices.BinarySearch(ts, t) // find slot
	return slices.Insert(ts, i, t)
}

func (w World) get_quadrant_index(r Robot) int {
	midx := w.bounds.x / 2
	midy := w.bounds.y / 2
	if r.p.x == midx || r.p.y == midy {
		return -1
	}

	if r.p.x >= 0 && r.p.x < midx && r.p.y >= 0 && r.p.y < midy {
		return 0
	}
	if r.p.x >= 0 && r.p.x < midx && r.p.y > midy && r.p.y < w.bounds.y {
		return 1
	}
	if r.p.x > midx && r.p.x < w.bounds.x && r.p.y >= 0 && r.p.y < midy {
		return 2
	}
	if r.p.x > midx && r.p.x < w.bounds.x && r.p.y > midy && r.p.y < w.bounds.y {
		return 3
	}
	return -1
}

func neighbor_score(robots []Robot) int {
	acc := 0
	for i, r := range robots {
		score := 0
		for i_n, r_n := range robots {
			if i_n == i {
				continue
			}
			if abs(r_n.p.x, r.p.x) == 1 && abs(r_n.p.y, r.p.y) == 1 {
				score += 1
			}

		}
		acc += score
	}
	return acc

}
func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a

}

func (w *World) update_tick(robots []Robot, tick_amt int) []Robot {
	next_robots := make([]Robot, 0, len(robots))
	for _, v := range robots {
		n_r := v
		n_r.p.x += n_r.v.x * tick_amt
		n_r.p.y += n_r.v.y * tick_amt

		n_r.p.x = n_r.p.x % w.bounds.x
		if n_r.p.x < 0 {
			n_r.p.x += w.bounds.x
		}
		n_r.p.y = n_r.p.y % w.bounds.y
		if n_r.p.y < 0 {
			n_r.p.y += w.bounds.y
		}

		next_robots = append(next_robots, n_r)
	}
	return next_robots

}

func (w *World) draw(robots []Robot) {
	fmt.Println("\033[2J")

	var sb strings.Builder
	for y := 0; y < w.bounds.y; y += 1 {
		for x := 0; x < w.bounds.x; x += 1 {
			found := false
			for _, robot := range robots {
				if robot.p.x == x && robot.p.y == y {
					sb.WriteString("X")
					found = true
					break
				}

			}

			if !found {
				sb.WriteString(" ")
			}

		}
		sb.WriteString("\n")
	}
	fmt.Println(sb.String())

}

func (w *World) update(robots []Robot) []Robot {
	next_robots := make([]Robot, 0, len(robots))
	for _, v := range robots {
		n_r := v
		n_r.p.x += n_r.v.x
		n_r.p.y += n_r.v.y

		n_r.p.x = n_r.p.x % w.bounds.x
		if n_r.p.x < 0 {
			n_r.p.x = w.bounds.x + n_r.p.x
		}
		n_r.p.y = n_r.p.y % w.bounds.y
		if n_r.p.y < 0 {
			n_r.p.y = w.bounds.y + n_r.p.y
		}

		next_robots = append(next_robots, n_r)
	}
	return next_robots

}

func parse_robot(line string) (Robot, error) {
	r := Robot{}

	pattern := regexp.MustCompile(`p=(-?\d+),(-?\d+) v=(-?\d+),(-?\d+)`)
	matches := pattern.FindStringSubmatch(line)

	p_x, err := strconv.ParseInt(matches[1], 10, 64)
	if err != nil {
		return r, err
	}
	p_y, err := strconv.ParseInt(matches[2], 10, 64)
	if err != nil {
		return r, err
	}
	v_x, err := strconv.ParseInt(matches[3], 10, 64)
	if err != nil {
		return r, err
	}
	v_y, err := strconv.ParseInt(matches[4], 10, 64)
	if err != nil {
		return r, err
	}

	r.p.x = int(p_x)
	r.p.y = int(p_y)
	r.v.x = int(v_x)
	r.v.y = int(v_y)
	return r, nil

}
