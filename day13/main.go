package day13

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Solver struct{}

type Vector2 struct {
	x int
	y int
}

type Machine struct {
	a     Vector2
	b     Vector2
	prize Vector2
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
	machines := make([]Machine, 0)

	var sb strings.Builder

	for scanner.Scan() {
		text := scanner.Text()
		if len(text) == 0 {
			if sb.Len() > 0 {
				pattern := regexp.MustCompile(`Button A: X\+(\d+), Y\+(\d+)Button B: X\+(\d+), Y\+(\d+)Prize: X\=(\d+), Y\=(\d+)`)
				matches := pattern.FindAllStringSubmatch(sb.String(), sb.Len())[0]
				nums := make([]int, 0, 6)
				for _, v := range matches[1:] {
					num, err := strconv.Atoi(v)
					if err != nil {
						return err
					}
					nums = append(nums, num)
				}

				machines = append(machines, Machine{
					a:     Vector2{nums[0], nums[1]},
					b:     Vector2{nums[2], nums[3]},
					prize: Vector2{nums[4], nums[5]},
				})
			}
			sb.Reset()
			continue
		}
		sb.WriteString(text)
	}

	if sb.Len() > 0 {
		pattern := regexp.MustCompile(`Button A: X\+(\d+), Y\+(\d+)Button B: X\+(\d+), Y\+(\d+)Prize: X\=(\d+), Y\=(\d+)`)
		matches := pattern.FindAllStringSubmatch(sb.String(), sb.Len())[0]
		nums := make([]int, 0, 6)
		for _, v := range matches[1:] {
			num, err := strconv.Atoi(v)
			if err != nil {
				return err
			}
			nums = append(nums, num)
		}

		machines = append(machines, Machine{
			a:     Vector2{nums[0], nums[1]},
			b:     Vector2{nums[2], nums[3]},
			prize: Vector2{nums[4], nums[5]},
		})
	}

	token_count := 0
	for _, machine := range machines {
		m_count := get_token_count(machine)
		if m_count > 0 {
			token_count += m_count
		}
	}
	fmt.Println(token_count)

	prepended_count := 0
	for _, machine := range machines {
		prepended_m := machine
		prepended_m.prize.x = 10000000000000 + prepended_m.prize.x
		prepended_m.prize.y = 10000000000000 + prepended_m.prize.y
		m_count := get_token_count(prepended_m)
		if m_count > 0 {
			prepended_count += m_count
		}

	}
	fmt.Println(prepended_count)

	return nil
}

// returns -1 if impossible
// gaussian eliminationnnnn :)
func get_token_count(m Machine) int {
	mat := [][]int{{m.a.x, m.b.x, m.prize.x}, {m.a.y, m.b.y, m.prize.y}}

	y_scaled := make([][]int, len(mat))
	y_scaled[0] = scalar_mult(mat[1][0], mat[0])
	y_scaled[1] = scalar_mult(mat[0][0], mat[1])

	if (y_scaled[0][2]-y_scaled[1][2])%(y_scaled[0][1]-y_scaled[1][1]) != 0 {
		return -1
	}
	second_token := (y_scaled[0][2] - y_scaled[1][2]) / (y_scaled[0][1] - y_scaled[1][1])
	if (mat[0][2]-second_token*mat[0][1])%mat[0][0] != 0 {
		return -1

	}
	first_token := (mat[0][2] - second_token*mat[0][1]) / mat[0][0]
	if first_token <= 0 || second_token <= 0 {
		log.Fatal(first_token, second_token)
	}
	return second_token + (3 * first_token)

}

func scalar_mult(scalar int, row []int) []int {
	multiplied := make([]int, len(row))
	for i, v := range row {
		multiplied[i] = v * scalar
	}
	return multiplied
}
