package day1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Day1Solver struct{}

func (d *Day1Solver) Solve() error {
	file, err := os.Open("day1/small.txt")
	if err != nil {
		return err
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	scanner := bufio.NewScanner(file)
	l_acc, r_acc := 0, 0
	for scanner.Scan() {
		values := strings.Fields(scanner.Text())
		var left int
		var right int
		var conv_err error
		left, conv_err = strconv.Atoi(values[0])
		if conv_err != nil {
			return conv_err
		}
		right, conv_err = strconv.Atoi(values[1])
		if conv_err != nil {
			return conv_err
		}
		l_acc += left
		r_acc += right
	}
	fmt.Println("Solution: ", r_acc-l_acc)
	return nil
}
