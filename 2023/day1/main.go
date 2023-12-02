package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
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

func ExtractNamedDigits(in string) int {
	var firstDigit string
	for i := 0; i < len(in); i++ {
		if unicode.IsDigit(rune(in[i])) {
			firstDigit = string(in[i])
			break
		}
		if strings.HasSuffix(in[0:i+1], "one") {
			firstDigit = "1"
			break
		}
		if strings.HasSuffix(in[0:i+1], "two") {
			firstDigit = "2"
			break
		}
		if strings.HasSuffix(in[0:i+1], "three") {
			firstDigit = "3"
			break
		}
		if strings.HasSuffix(in[0:i+1], "four") {
			firstDigit = "4"
			break
		}
		if strings.HasSuffix(in[0:i+1], "five") {
			firstDigit = "5"
			break
		}
		if strings.HasSuffix(in[0:i+1], "six") {
			firstDigit = "6"
			break
		}
		if strings.HasSuffix(in[0:i+1], "seven") {
			firstDigit = "7"
			break
		}
		if strings.HasSuffix(in[0:i+1], "eight") {
			firstDigit = "8"
			break
		}
		if strings.HasSuffix(in[0:i+1], "nine") {
			firstDigit = "9"
			break
		}
	}

	var lastDigit string
	for i := len(in) - 1; i >= 0; i-- {
		if unicode.IsDigit(rune(in[i])) {
			lastDigit = string(in[i])
			break
		}
		if strings.HasPrefix(in[i:], "one") {
			lastDigit = "1"
			break
		}
		if strings.HasPrefix(in[i:], "two") {
			lastDigit = "2"
			break
		}
		if strings.HasPrefix(in[i:], "three") {
			lastDigit = "3"
			break
		}
		if strings.HasPrefix(in[i:], "four") {
			lastDigit = "4"
			break
		}
		if strings.HasPrefix(in[i:], "five") {
			lastDigit = "5"
			break
		}
		if strings.HasPrefix(in[i:], "six") {
			lastDigit = "6"
			break
		}
		if strings.HasPrefix(in[i:], "seven") {
			lastDigit = "7"
			break
		}
		if strings.HasPrefix(in[i:], "eight") {
			lastDigit = "8"
			break
		}
		if strings.HasPrefix(in[i:], "nine") {
			lastDigit = "9"
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
	numbers := []int{}
	for _, line := range bytes.Split(input, []byte("\n")) {
		number := ExtractNamedDigits(string(line))
		numbers = append(numbers, number)
	}

	var sum int
	for _, num := range numbers {
		sum += num
	}

	return sum
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
