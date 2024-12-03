package day3

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
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
	file_total := 0
	enabled := true
	for scanner.Scan() {
		var line_val int
		enabled, line_val = parse_input(enabled, scanner.Text())

		file_total += line_val
	}
	fmt.Println(file_total)
	return nil
}

func parse_input(enabled bool, raw string) (bool, int) {
	sum := 0
	pattern := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)|do\(\)|don't\(\)`)
	matches := pattern.FindAllStringSubmatch(raw, len(raw))
	enabled_status := enabled
	for _, v := range matches {
		switch v[0] {
		case "do()":
			enabled_status = true
		case "don't()":
			enabled_status = false
		default:
			if !enabled_status {
				continue
			}
			param_1, conv_err := strconv.Atoi(v[1])
			param_2, conv_err_2 := strconv.Atoi(v[2])

			if conv_err != nil {
				log.Fatal("Bad regexp:", conv_err)
			}
			if conv_err_2 != nil {
				log.Fatal("Bad regexp:", conv_err_2)
			}
			sum += (param_1 * param_2)
		}
	}

	return enabled_status, sum
}
