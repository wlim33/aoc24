package day4

import (
	"bufio"
	"fmt"
	"log"
	"os"
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
	xmas_count := 0
	grid := make([]string, 0)
	for scanner.Scan() {
		text := scanner.Text()
		grid = append(grid, text)

	}
	d_count := 0
	for r_i := 0; r_i < len(grid)-3; r_i += 1 {
		for c_i := 0; c_i < len(grid[r_i])-3; c_i += 1 {
			// diagonals
			d_count += check(string(grid[r_i][c_i]), string(grid[r_i+1][c_i+1]), string(grid[r_i+2][c_i+2]), string(grid[r_i+3][c_i+3]))
			d_count += check(string(grid[r_i+3][c_i]), string(grid[r_i+2][c_i+1]), string(grid[r_i+1][c_i+2]), string(grid[r_i][c_i+3]))
		}
	}
	v_count := 0
	for r_i := 0; r_i < len(grid)-3; r_i += 1 {
		for c_i := 0; c_i < len(grid[r_i]); c_i += 1 {
			// fmt.Println(r_i, c_i, ",", r_i+1, c_i, ",", r_i+2, c_i, ",", r_i+3, c_i)
			v_count += check(string(grid[r_i][c_i]), string(grid[r_i+1][c_i]), string(grid[r_i+2][c_i]), string(grid[r_i+3][c_i]))
		}
	}
	h_count := 0
	for _, row := range grid {
		h_count += strings.Count(row, "XMAS") + strings.Count(row, "SAMX")
	}

	mas_count := 0
	for r_i := 0; r_i < len(grid)-2; r_i += 1 {
		for c_i := 0; c_i < len(grid[r_i])-2; c_i += 1 {
			// diagonals
			l_r := check_mas(string(grid[r_i][c_i]), string(grid[r_i+1][c_i+1]), string(grid[r_i+2][c_i+2]))
			if l_r > 0 {
				r_l := check_mas(string(grid[r_i+2][c_i]), string(grid[r_i+1][c_i+1]), string(grid[r_i][c_i+2]))
				if r_l > 0 {
					mas_count += 1
				}
			}
		}
	}
	xmas_count += h_count
	xmas_count += d_count
	xmas_count += v_count
	fmt.Println(xmas_count)
	fmt.Println(mas_count)
	// fmt.Println(grid)
	return nil
}

func check_mas(s_1 string, s_2 string, s_3 string) int {
	count := 0
	if s_1 == "M" && s_2 == "A" && s_3 == "S" {
		count += 1
	}
	if s_1 == "S" && s_2 == "A" && s_3 == "M" {
		count += 1
	}
	return count

}

func check(s_1 string, s_2 string, s_3 string, s_4 string) int {
	count := 0
	if s_1 == "X" && s_2 == "M" && s_3 == "A" && s_4 == "S" {
		count += 1
	}
	if s_1 == "S" && s_2 == "A" && s_3 == "M" && s_4 == "X" {
		count += 1
	}
	return count

}
