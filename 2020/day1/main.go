package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func findSumThree(want int, nums []int) (int, int, int, error) {
	for _, n := range nums {
		for _, m := range nums {
			for _, x := range nums {
				if n+m+x == want {
					return n, m, x, nil
				}
			}
		}
	}
	return 0, 0, 0, fmt.Errorf("Could not find three numbers that added to %d", want)
}

func findSum(want int, nums []int) (int, int, error) {
	for _, n := range nums {
		for _, m := range nums {
			if n+m == want {
				return n, m, nil
			}
		}
	}
	return 0, 0, fmt.Errorf("Could not find two numbers that added to %d", want)
}

func main() {
	f, err := os.OpenFile("input", os.O_RDONLY, os.ModePerm)
	if err != nil {
		panic(err)
	}
	s := bufio.NewScanner(f)
	var ins []int
	for s.Scan() {
		line := s.Text()
		line = strings.TrimSpace(line)
		i, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		ins = append(ins, i)
	}
	// Part 1
	// first, second, err := findSum(2020, ins)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(first * second)

	// Part 2
	first, second, third, err := findSumThree(2020, ins)
	if err != nil {
		panic(err)
	}
	fmt.Println(first * second * third)
}
