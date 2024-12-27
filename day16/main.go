package day14

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

	for scanner.Scan() {
		text := scanner.Text()
	}
	fmt.Println(text)
	return nil
}
