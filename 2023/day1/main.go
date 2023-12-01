package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

var (
	part = flag.Int("part", 1, "the part to run")
	file = flag.String("file", "example.txt", "the file to read")
)

func PartOne(input []byte) int {
	numbers := []int{}
	for _, line := range bytes.Split(input, []byte("\n")) {
		number := ExtractDigits(string(line))
		numbers = append(numbers, number)
	}

	var sum int
	for _, num := range numbers {
		sum += num
	}

	return sum
}

func ExtractDigits(in string) int {
	var firstDigit string
	for i := 0; i < len(in); i++ {
		if unicode.IsDigit(rune(in[i])) {
			firstDigit = string(in[i])
			break
		}
	}

	var lastDigit string
	for i := len(in) - 1; i >= 0; i-- {
		if unicode.IsDigit(rune(in[i])) {
			lastDigit = string(in[i])
			break
		}
	}

	twoDigit := firstDigit + lastDigit

	out, err := strconv.Atoi(twoDigit)
	if err != nil {
		panic(err)
	}

	return out
}

func PartTwo(input []byte) int {
	return 0
}

func main() {
	flag.Parse()

	in, err := os.ReadFile(*file)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	switch *part {
	case 1:
		fmt.Println(PartOne(in))
	case 2:
		fmt.Println(PartTwo(in))
	}
}
