package main

import (
	"fmt"
	"strconv"

	"github.com/while1malloc0/advent-of-code/2020/challenge"
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
	var ins []int
	err := challenge.InputScanFunc("input", func(s string) error {
		i, err := strconv.Atoi(s)
		if err != nil {
			return err
		}
		ins = append(ins, i)
		return nil
	})
	if err != nil {
		panic(err)
	}
	partOneFn := func() error {
		first, second, err := findSum(2020, ins)
		if err != nil {
			return err
		}
		fmt.Println(first * second)
		return nil
	}
	partTwoFn := func() error {
		first, second, third, err := findSumThree(2020, ins)
		if err != nil {
			return err
		}
		fmt.Println(first * second * third)
		return nil
	}
	err = challenge.Run(partOneFn, partTwoFn)
	if err != nil {
		panic(err)
	}
}
