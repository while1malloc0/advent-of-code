package main

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/while1malloc0/advent-of-code/2020/challenge"
)

type queue struct {
	items []int64
}

func (q *queue) enqueue(num int64) {
	q.items = append(q.items, num)
}

func (q *queue) dequeue() (int64, error) {
	if len(q.items) == 0 {
		return 0, errors.New("Empty queue")
	}
	ret := q.items[0]
	q.items = q.items[1:]
	return ret, nil
}

func isSumOfTwoNums(needle int64, haystack []int64) bool {
	for _, outer := range haystack {
		for _, inner := range haystack {
			if outer+inner == needle {
				return true
			}
		}
	}
	return false
}

func findFirstNonSum(preambleLen int, haystack []int64) (int64, error) {
	q := &queue{}
	for i := 0; i < preambleLen; i++ {
		q.enqueue(haystack[i])
	}
	for i := preambleLen; i < len(haystack); i++ {
		if !isSumOfTwoNums(haystack[i], q.items) {
			return haystack[i], nil
		}
		q.dequeue()
		q.enqueue(haystack[i])
	}
	return 0, errors.New("No non sum found")
}

func sumsToNumber(want int64, candidates []int64) bool {
	var sum int64
	for i := range candidates {
		sum += candidates[i]
	}
	return sum == want
}

func findContiguousRange(needle int64, haystack []int64) ([]int64, error) {
	rangeSize := 2
	for {
		start := 0
		end := rangeSize
		if rangeSize >= len(haystack) {
			break
		}
		for end < len(haystack) {
			rng := haystack[start:end]
			if sumsToNumber(needle, rng) {
				return haystack[start:end], nil
			}
			start++
			end++
		}
		rangeSize++
	}
	return nil, errors.New("Range not found")
}

// Seriously Go?!
func findMin(haystack []int64) int64 {
	min := haystack[0]
	for _, num := range haystack {
		if num < min {
			min = num
		}
	}
	return min
}

func findMax(haystack []int64) int64 {
	max := haystack[0]
	for _, num := range haystack {
		if num > max {
			max = num
		}
	}
	return max
}

func main() {
	var in []int64
	challenge.InputScanFunc("input", func(s string) error {
		num, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			panic(err)
		}
		in = append(in, num)
		return nil
	})

	partOneFunc := func() error {
		result, err := findFirstNonSum(25, in)
		if err != nil {
			return err
		}
		fmt.Println(result)
		return nil
	}

	partTwoFunc := func() error {
		result, err := findContiguousRange(542529149, in)
		if err != nil {
			return err
		}
		smallest := findMin(result)
		largest := findMax(result)
		fmt.Println(smallest, largest)
		fmt.Println(smallest + largest)
		return nil
	}

	challenge.Run(partOneFunc, partTwoFunc)
}
