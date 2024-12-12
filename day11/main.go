package day11

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Solver struct{}

type StackItem struct {
	num       int
	remaining int
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
	stones := make([]int, 0)
	for scanner.Scan() {
		text := scanner.Text()

		for _, v := range strings.Fields(text) {
			num, err := strconv.Atoi(v)
			if err != nil {
				return err
			}
			stones = append(stones, num)
		}
		fmt.Println(text)
	}

	// 0 <= i < big number
	// 0 <= j < 75
	// memoized[i][j] == number of stones generated starting from i after j updates
	// memoized := make(map[int][]int)
	// func(num, 0) = 1
	// func(num, remaining) = sum(for all generated, func(generated, remaining - 1))
	memoized := make(map[StackItem]int)
	stack := make([]StackItem, 0)
	iteration_count := 75
	for _, stone := range stones {
		stack = append(stack, StackItem{
			num:       stone,
			remaining: iteration_count,
		})
	}
	total := 0

	stack_count := 0
	for len(stack) > 0 {
		stack_count += 1
		var current StackItem
		current, stack = stack[len(stack)-1], stack[:len(stack)-1]
		if current.remaining == 0 {
			memoized[current] = 1
			total += 1
			continue
		}

		next := update_stone(current.num)
		found := 0
		all_found := true
		for _, v := range next {

			s := StackItem{
				num:       v,
				remaining: current.remaining - 1,
			}
			if count, ok := memoized[s]; ok {
				// fmt.Println("hit")
				found += count
			} else {
				stack = append(stack, s)
				all_found = false
			}
		}
		total += found
		if all_found {
			memoized[current] = found

		}
	}

	fmt.Println("stack_count:", stack_count)
	fmt.Println(total)

	return nil
}

func update_stone(s int) []int {
	next := make([]int, 0)
	if s == 0 {
		next = append(next, 1)
		return next
	}
	is_even_size, tens := has_even_digits(s)
	if is_even_size {
		left := s / tens
		right := s - (left * tens)
		next = append(next, left)
		next = append(next, right)
		return next
	}

	next = append(next, s*2024)
	return next

}

// returns is even number of digits and 10 raised to half number of digits
func has_even_digits(num int) (bool, int) {
	if num == 0 {
		return false, 1
	}
	tens := 1
	count := 0
	for num != 0 {
		num = num / 10
		if count%2 == 0 {

			tens = tens * 10
		}
		count += 1
	}
	return count%2 == 0, tens

}
