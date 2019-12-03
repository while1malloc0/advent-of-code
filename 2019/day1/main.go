package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("input")
	defer f.Close()
	if err != nil {
		panic(err)
	}
	firstPart(f)
	f, err = os.Open("input")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	secondPart(f)
}

func firstPart(f *os.File) {
	r := bufio.NewScanner(f)
	var sum int64
	for r.Scan() {
		text := r.Text()
		text = strings.TrimSpace(text)
		num, err := strconv.ParseInt(text, 0, 0)
		if err != nil {
			panic(err)
		}
		sum += int64(math.Floor(float64(num/3)) - 2)
	}
	fmt.Println(sum)
}

func secondPart(f *os.File) {
	r := bufio.NewScanner(f)
	var sum int64
	for r.Scan() {
		text := r.Text()
		text = strings.TrimSpace(text)
		num, err := strconv.ParseInt(text, 0, 0)
		if err != nil {
			panic(err)
		}
		sum += getFuelWithSelf(num)
	}
	fmt.Println(sum)
}

func getFuelWithSelf(n int64) int64 {
	rem := int64(math.Floor(float64(n/3)) - 2)
	sum := rem
	for {
		rem = int64(math.Floor(float64(rem/3)) - 2)
		if rem < 1 {
			break
		}
		sum += rem
	}
	return int64(sum)
}
