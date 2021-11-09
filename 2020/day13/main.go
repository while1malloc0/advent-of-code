package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"

	"github.com/while1malloc0/advent-of-code/2020/challenge"
)

const SKIP = -1

func getNearestMultipleOf(from, to int) int {
	dived := to / from
	floored := math.Floor(float64(dived))
	multiplier := int(floored) + 1
	return from * multiplier
}

func getNearestBusID(to int, candidates []int) (int, int) {
	min := math.MaxInt64
	var nearestCandidate int
	var nearestDeparture int
	for i := range candidates {
		departure := getNearestMultipleOf(candidates[i], to)
		if departure < min {
			min = departure
			nearestCandidate = candidates[i]
			nearestDeparture = departure
		}
	}
	return nearestCandidate, nearestDeparture
}

func checkMagicCandidate(candidate int, nums []int) bool {
	for i := range nums {
		if nums[i] == SKIP {
			continue
		}
		if !((candidate+i)%nums[i] == 0) {
			return false
		}
	}
	return true
}

// this likely has a math term, I don't know what it is
// func getMagicTimestamp(nums []int) int {
// 	var i int
// 	for {
// 		candidate := nums[0] * i
// 		if checkMagicCandidate(candidate, nums) {
// 			return candidate
// 		}
// 		i++
// 	}
// }

func getMagicTimestamp(nums []int) int {
	candidate := nums[0]
	increment := nums[0]
	for i := 0; i < len(nums)-1; i++ {
		if nums[i+1] == SKIP {
			continue
		}
		for {
			if (candidate+i+1)%nums[i+1] == 0 {
				increment *= nums[i+1]
				break
			}
			candidate += increment
		}
	}
	return candidate
}

// func getMagicTimestamp(nums []int) int {
// 	mult := 1
// 	for i := range nums {
// 		if nums[i] == SKIP {
// 			continue
// 		}
// 		mult *= nums[i]
// 	}

// 	var j int
// 	for {
// 		if j%mult == j {
// 			return j
// 		}
// 	}
// }

func main() {
	partOneFunc := func() error {
		in, err := ioutil.ReadFile("input")
		if err != nil {
			return err
		}
		lines := strings.Split(string(in), "\n")
		target, err := strconv.Atoi(lines[0])
		if err != nil {
			return err
		}

		candidateStrings := strings.Split(lines[1], ",")
		var candidates []int
		for i := range candidateStrings {
			if candidateStrings[i] == "x" {
				continue
			}
			candidate, err := strconv.Atoi(candidateStrings[i])
			if err != nil {
				return err
			}
			candidates = append(candidates, candidate)
		}

		nearestCandidate, nearestDeparture := getNearestBusID(target, candidates)
		timeWaiting := nearestDeparture - target
		fmt.Println(nearestCandidate * timeWaiting)
		return nil
	}

	partTwoFunc := func() error {
		in, err := ioutil.ReadFile("input")
		if err != nil {
			return err
		}
		lines := strings.Split(string(in), "\n")
		candidateStrings := strings.Split(lines[1], ",")
		var candidates []int
		for i := range candidateStrings {
			if candidateStrings[i] == "x" {
				candidates = append(candidates, SKIP)
				continue
			}
			candidate, err := strconv.Atoi(candidateStrings[i])
			if err != nil {
				return err
			}
			candidates = append(candidates, candidate)
		}
		fmt.Println(getMagicTimestamp(candidates))
		return nil
	}

	challenge.Run(partOneFunc, partTwoFunc)
}
