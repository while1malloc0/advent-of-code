package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/while1malloc0/advent-of-code/2020/challenge"
)

func runGame(startingNums []int, turns int) int {
	seen := map[int][2]int{}
	for i := range startingNums {
		num := startingNums[i]
		seen[num] = [2]int{-1, i + 1}
	}
	current := 0
	for i := len(startingNums) + 1; i < turns; i++ {
		if _, ok := seen[current]; ok {
			prev := seen[current]
			if prev[0] == -1 {
				prev[0] = i
				seen[current] = prev
				current = prev[0] - prev[1]
			} else {
				prev[1] = prev[0]
				prev[0] = i
				seen[current] = prev
				current = prev[0] - prev[1]
			}
		} else {
			seen[current] = [2]int{-1, i}
			current = 0
		}
	}
	return current
}

func main() {
	partOneFunc := func() error {
		in, err := ioutil.ReadFile("input")
		if err != nil {
			return err
		}
		numStrs := strings.Split(string(in), ",")
		var nums []int
		for i := range numStrs {
			num, err := strconv.Atoi(numStrs[i])
			if err != nil {
				return err
			}
			nums = append(nums, num)
		}
		result := runGame(nums, 2020)
		fmt.Println(result)
		return nil
	}

	partTwoFunc := func() error {
		in, err := ioutil.ReadFile("input")
		if err != nil {
			return err
		}
		numStrs := strings.Split(string(in), ",")
		var nums []int
		for i := range numStrs {
			num, err := strconv.Atoi(numStrs[i])
			if err != nil {
				return err
			}
			nums = append(nums, num)
		}
		result := runGame(nums, 30000000)
		fmt.Println(result)
		return nil
	}

	challenge.Run(partOneFunc, partTwoFunc)
}
