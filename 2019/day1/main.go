package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const minNeededFuelToMatter = 1

func main() {
	nums := processInput("input")
	calcFuelNaive(nums)
	calcFuelAndDealWithIt(nums)
}

func processInput(path string) []int64 {
	var out []int64
	f, err := os.Open(path)
	defer f.Close()
	if err != nil {
		panic(err)
	}
	r := bufio.NewScanner(f)
	for r.Scan() {
		text := strings.TrimSpace(r.Text())
		num, err := strconv.ParseInt(text, 0, 0)
		if err != nil {
			panic(err)
		}
		out = append(out, num)
	}
	return out
}

func calcFuelNaive(nums []int64) {
	var sum int64
	for _, num := range nums {
		sum += calcFuel(num)
	}
	fmt.Println(sum)
}

func calcFuelAndDealWithIt(nums []int64) {
	var sum int64
	for _, num := range nums {
		sum += calcFuelWhilePositive(num)
	}
	fmt.Println(sum)
}

func calcFuel(n int64) int64 {
	return int64(math.Floor(float64(n/3)) - 2)
}

func calcFuelWhilePositive(n int64) int64 {
	neededFuel := calcFuel(n)
	totalNeededFuel := neededFuel
	for {
		neededFuel = calcFuel(neededFuel)
		if neededFuel < minNeededFuelToMatter {
			break
		}
		totalNeededFuel += neededFuel
	}
	return totalNeededFuel
}
