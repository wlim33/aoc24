package day7

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

	inputs := make(map[uint64][]uint64)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		x := strings.Split(text, ":")
		right_exp := strings.Fields(x[1])
		r := make([]uint64, 0)

		for _, v := range right_exp {
			num, err := strconv.ParseUint(v, 10, 64)
			if err != nil {
				return err
			}
			r = append(r, num)
		}
		left, err := strconv.ParseUint(x[0], 10, 64)
		if err != nil {
			return err
		}
		inputs[left] = r
	}
	var acc uint64
	acc = 0

	for left, right := range inputs {
		if can_compute(left, right) {
			fmt.Println(left, right)
			acc += left
		}
	}
	fmt.Println(acc)
	return nil
}

const (
	add = iota
	multiply
)

type StackItem struct {
	operators []uint64
}

func can_compute(total uint64, operands []uint64) bool {
	stack := make([]StackItem, 0)
	stack = append(stack, StackItem{make([]uint64, 0)})
	for len(stack) > 0 {
		var current StackItem
		current, stack = stack[len(stack)-1], stack[:len(stack)-1]
		if len(current.operators) >= len(operands)-1 {
			acc := operands[0]
			for i, operand := range operands {
				if i == 0 {
					continue
				}
				switch current.operators[i-1] {
				case multiply:
					acc = acc * operand
				case add:
					acc = acc + operand
				}
			}
			if acc == total {
				return true
			}
		} else {
			stack = append(stack, StackItem{append(current.operators, add)})
			stack = append(stack, StackItem{append(current.operators, multiply)})
		}
	}
	return false
}
