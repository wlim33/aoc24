package day1

import (
	"bufio"
	"container/heap"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type Solver struct {
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
	l_heap, r_heap := IntHeap{}, IntHeap{}
	heap.Init(&l_heap)
	heap.Init(&r_heap)

	l_freq, r_freq := make(map[int]int), make(map[int]int)
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

		heap.Push(&l_heap, left)
		heap.Push(&r_heap, right)

		if val, ok := l_freq[left]; ok {
			l_freq[left] = val + 1
		} else {
			l_freq[left] = 1
		}

		if val, ok := r_freq[right]; ok {
			r_freq[right] = val + 1
		} else {
			r_freq[right] = 1
		}

	}
	total := 0
	for l_heap.Len() > 0 && r_heap.Len() > 0 {
		l := heap.Pop(&l_heap).(int)
		r := heap.Pop(&r_heap).(int)
		fmt.Println(l, r)
		if l < r {
			total += (r - l)
		} else {
			total += (l - r)
		}
	}
	fmt.Println("Solution a: ", total)

	similarity_score := 0
	for k, v := range l_freq {

		if val, ok := r_freq[k]; ok {
			similarity_score += (v * val * k)
		}
	}

	fmt.Println("Solution b: ", similarity_score)
	return nil
}
