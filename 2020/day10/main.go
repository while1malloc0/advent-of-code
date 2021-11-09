package main

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/while1malloc0/advent-of-code/2020/challenge"
)

func countOnesAndThrees(in []int) (int, int) {
	sorted := []int{0}

	for i := range in {
		sorted = append(sorted, in[i])
	}
	sort.Ints(sorted)
	sorted = append(sorted, sorted[len(sorted)-1]+3)
	var ones int
	var threes int
	for i := 0; i <= len(sorted)-2; i++ {
		if sorted[i+1]-sorted[i] == 3 {
			threes++
		}
		if sorted[i+1]-sorted[i] == 1 {
			ones++
		}
	}

	return ones, threes
}

func findNumPaths(in []int) int {
	sorted := []int{0}
	for _, n := range in {
		sorted = append(sorted, n)
	}
	sort.Ints(sorted)
	sorted = append(sorted, sorted[len(sorted)-1]+3)

	paths := map[int]int{}
	paths[0] = 1
	for _, num := range sorted {
		if num == 0 {
			continue
		}
		if val, ok := paths[num-1]; ok {
			paths[num] += val
		}
		if val, ok := paths[num-2]; ok {
			paths[num] += val
		}
		if val, ok := paths[num-3]; ok {
			paths[num] += val
		}
	}
	return paths[sorted[len(sorted)-1]]
}

func main() {
	var in []int
	challenge.InputScanFunc("input", func(s string) error {
		num, err := strconv.Atoi(s)
		if err != nil {
			return err
		}
		in = append(in, num)
		return nil
	})

	partOneFunc := func() error {
		ones, threes := countOnesAndThrees(in)
		fmt.Println(ones * threes)
		return nil
	}

	partTwoFunc := func() error {
		result := findNumPaths(in)
		fmt.Println(result)
		return nil
	}

	challenge.Run(partOneFunc, partTwoFunc)
}
