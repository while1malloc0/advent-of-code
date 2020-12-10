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

	challenge.Run(partOneFunc, nil)
}
