package main

import (
	"flag"
	"fmt"
	"strconv"
)

var part2 = flag.Bool("part2", false, "run part two")

func main() {
	flag.Parse()

	assert(!meetsCriteria("111111"), "111111 should be false")
	assert(!meetsCriteria("123444"), "123444 should be false")
	assert(meetsCriteria("111122"), "111122 should be true")

	start := "357253"
	end := "892942"
	candidates := genRange(start, end)

	var answer int
	for _, candidate := range candidates {
		if meetsCriteria(candidate) {
			answer++
		}
	}
	fmt.Println(answer)
}

func genRange(start, end string) []string {
	out := []string{}
	i, err := strconv.Atoi(start)
	if err != nil {
		panic(err)
	}
	max, err := strconv.Atoi(end)
	if err != nil {
		panic(err)
	}
	for i <= max {
		s := strconv.Itoa(i)
		out = append(out, s)
		i++
	}
	return out
}

func assert(predicate bool, msg string) {
	if !predicate {
		panic(fmt.Sprintf("Failed assertion: %s", msg))
	}
}

func meetsCriteria(s string) bool {
	nums := toIntSlice(s)
	return isIncreasing(nums) && hasDoubleVal(nums)
}

func hasDoubleVal(nums []int) bool {
	counts := make(map[int]int)
	for _, n := range nums {
		_, seen := counts[n]
		if !seen {
			counts[n] = 1
		} else {
			counts[n]++
		}
	}
	for _, count := range counts {
		if count > 1 && count < 3 {
			return true
		}
	}
	return false
}

func isIncreasing(nums []int) bool {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			return false
		}
	}
	return true
}

func toIntSlice(s string) []int {
	nums := []int{}
	for _, str := range s {
		i, err := strconv.Atoi(string(str))
		if err != nil {
			panic(err)
		}
		nums = append(nums, i)
	}
	return nums
}
