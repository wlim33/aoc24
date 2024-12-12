package day9

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Solver struct{}
type FreeRange struct {
	start int
	end   int
}

func (f FreeRange) length() int {
	return f.end - f.start

}

type Block = int

func is_empty(b Block) bool {
	return b < 0

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

	disk_map := make([]Block, 0)
	free_queue := make([]FreeRange, 0)

	file_id := 0
	count := 0
	for scanner.Scan() {
		text := scanner.Text()

		is_file := true
		for _, v := range text {
			num := int(v) - '0'
			count += num
			if is_file {
				for i := 0; i < num; i += 1 {
					disk_map = append(disk_map, Block(file_id))
				}
				file_id += 1
			} else {
				f_range := FreeRange{
					start: len(disk_map),
					end:   0,
				}
				for i := 0; i < num; i += 1 {
					disk_map = append(disk_map, Block(-1))
				}
				f_range.end = len(disk_map)

				free_queue = append(free_queue, f_range)

			}
			is_file = !is_file

		}
	}
	file_id -= 1
	fmt.Println("check:", count, len(disk_map))
	fmt.Println("first file:", file_id)

	// fmt.Println(disk_map)
	// fmt.Println(free_queue)
	file_index := len(disk_map) - 1
	file_length := 0
	for file_index > 0 {
		file_length = 0
		for file_index-file_length >= 0 && disk_map[file_index-file_length] == file_id {

			file_length += 1
		}
		// fmt.Printf("file %d is length %d\n", file_id, file_length)
		// found_flag := false
		for f_i := 0; f_i < len(free_queue); f_i += 1 {
			if free_queue[f_i].start > file_index {
				break
			}
			if free_queue[f_i].length() >= file_length {
				// fmt.Printf("moving file %d to diskmap[%d:%d]\n", file_id, free_queue[f_i].start, free_queue[f_i].start+file_length)
				for i := free_queue[f_i].start; i < free_queue[f_i].start+file_length; i += 1 {
					disk_map[i] = file_id
				}
				for i := file_index; file_index-i < file_length; i -= 1 {
					disk_map[i] = -1
				}

				free_queue[f_i].start += file_length
				file_index -= file_length
				// found_flag = true
				break
			}
		}
		// if !found_flag {
		// 	fmt.Printf("failed to move file %d\n", file_id)
		// }
		file_id -= 1
		for file_index >= 0 && disk_map[file_index] != file_id {
			file_index -= 1
		}

	}

	// fmt.Println(disk_map)
	fmt.Println(checksum(disk_map))
	return nil
}

func checksum(disk []Block) int {
	acc := 0
	for i, v := range disk {
		if v < 0 {
			continue
		}
		acc += i * v
	}
	return acc

}
