package day2

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Solver struct {
}

func abs_difference(a int, b int) int {
	if a-b < 0 {
		return b - a
	}
	return a - b
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

	count := 0
	loose_count := 0
	for scanner.Scan() {
		values := strings.Fields(scanner.Text())

		levels := make([]int, len(values))
		for i, v := range values {

			level, err := strconv.Atoi(v)
			if err != nil {
				return err
			}
			levels[i] = level

		}

		fmt.Println(levels)
		if check_valid(levels, -1) {
			count += 1
			loose_count += 1
		} else {

			for i := 0; i < len(levels); i += 1 {
				if check_valid(levels, i) {
					loose_count += 1
					break
				}
			}
		}
	}

	fmt.Println("Solution a: ", count)
	fmt.Println("Solution b: ", loose_count)

	return nil
}

func check_valid(levels []int, skip_index int) bool {
	start_i := 1
	stop_i := len(levels)
	if skip_index == 0 {
		start_i += 1
	}
	if skip_index == len(levels) {
		stop_i -= 1
	}

	var is_increasing bool
	var first_pair bool
	for i := start_i; i < stop_i; i += 1 {
		level := levels[i]
		last := levels[i-1]
		if skip_index != -1 && skip_index == i-1 {
			if i-2 >= 0 {
				last = levels[i-2]
			} else {
				continue
			}
		}
		if skip_index != -1 && skip_index == i {
			continue
		}
		difference := abs_difference(level, last)
		if (difference < 1) || (difference > 3) {
			return false
		}
		if !first_pair {
			is_increasing = (level - last) > 0
			first_pair = true
			continue
		}
		if is_increasing != ((level - last) > 0) {
			return false
		}
	}
	return true
}
