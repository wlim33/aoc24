package day5

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
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
	ordering_flag := true
	rules := make(map[int]map[int]struct{})
	pages := make([][]int, 0)
	wrong_pages := make([][]int, 0)
	for scanner.Scan() {
		text := scanner.Text()
		if len(text) == 0 {
			ordering_flag = false
			continue
		}

		if ordering_flag {
			raw := strings.Split(text, "|")
			first, err := strconv.Atoi(raw[0])
			if err != nil {
				return err
			}
			after, err := strconv.Atoi(raw[1])

			if err != nil {
				return err
			}
			if _, ok := rules[first]; ok {
				rules[first][after] = struct{}{}
			} else {
				rules[first] = map[int]struct{}{after: {}}
			}
		} else {
			line, err := convert_all(strings.Split(text, ","))
			if err != nil {
				return err
			}
			is_valid, rights, wrongs := is_valid(line, rules)
			if is_valid {
				pages = append(pages, line)
			} else {
				// TODO: check all possible combinations of wrongs[i] in right

				constructed := generate_correct(rights, wrongs, rules)
				wrong_pages = append(wrong_pages, constructed)
			}

		}
	}

	sum := 0
	for _, v := range pages {
		sum += v[len(v)/2]
	}

	wrong_sum := 0
	for _, v := range wrong_pages {
		wrong_sum += v[len(v)/2]
	}
	fmt.Println(sum)
	fmt.Println(wrong_sum)

	return nil
}
func is_valid(line []int, rules map[int]map[int]struct{}) (bool, []int, []int) {

	is_valid := true
	seen := map[int]struct{}{}
	wrongs := make([]int, 0)
	rights := make([]int, 0)

	for i := len(line) - 1; i >= 0; i -= 1 {
		num := line[i]
		rule, ok := rules[num]
		found_one := false
		for after := range seen {
			if (ok && !contains(after, rule)) || contains(num, rules[after]) {
				is_valid = false
				found_one = true
				wrongs = append(wrongs, num)
				break
			}
		}
		if !found_one {
			rights = append([]int{num}, rights...)
		}
		seen[num] = struct{}{}

	}
	return is_valid, rights, wrongs
}

func contains(a int, in map[int]struct{}) (ok bool) {
	_, ok = in[a]
	return
}
func convert_all(split_values []string) ([]int, error) {
	converted := make([]int, 0, len(split_values))
	for _, v := range split_values {
		num, err := strconv.Atoi(v)
		if err != nil {
			return converted, err
		}
		converted = append(converted, num)
	}

	return converted, nil

}

type StackItem struct {
	vs    []int
	index int
}

func generate_correct(right []int, wrongs []int, rules map[int]map[int]struct{}) []int {
	stack := make([]StackItem, 0)
	stack = append(stack,
		StackItem{
			vs:    right,
			index: 0,
		},
	)
	for len(stack) > 0 {
		var current StackItem
		current, stack = stack[len(stack)-1], stack[:len(stack)-1]

		for i := 0; i <= len(current.vs); i += 1 {
			cloned := append(current.vs[:0:0], current.vs...)
			candidate := slices.Insert(cloned, i, wrongs[current.index])
			is_valid, _, _ := is_valid(candidate, rules)
			if is_valid {
				if current.index >= len(wrongs)-1 {
					return candidate
				}
				stack = append(stack, StackItem{
					vs:    candidate,
					index: current.index + 1,
				})
			}
		}
	}
	return right

}
