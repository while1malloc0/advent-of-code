package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/while1malloc0/advent-of-code/2020/challenge"
)

type instruction struct {
	location int
	value    int
}

type program struct {
	mask         string
	instructions []instruction
}

var memSetRegex = regexp.MustCompile("mem\\[(\\d+)\\] = (\\d+)")

func runProgram(in string) (int, error) {
	r := strings.NewReader(in)
	s := bufio.NewScanner(r)
	// Largest 36 bit number according to Wikipedia
	var memory [4294967295]int

	var mask string
	for s.Scan() {
		line := s.Text()

		if strings.HasPrefix(line, "mask") {
			newMask := strings.TrimSpace(strings.Split(line, "=")[1])
			mask = newMask
			continue
		}

		if strings.HasPrefix(line, "mem") {
			matched := memSetRegex.FindStringSubmatch(line)

			value, err := strconv.Atoi(matched[2])
			if err != nil {
				return 0, err
			}
			result := applyMask(mask, value)

			location, err := strconv.Atoi(matched[1])
			if err != nil {
				return 0, err
			}
			memory[location] = result
		}
	}

	var sum int
	for i := range memory {
		sum += memory[i]
	}

	return sum, nil
}

// Go HaS sUcH a GrEaT sTaNdArD lIbRaRy
func reverseString(s string) string {
	var out string
	for i := len(s) - 1; i >= 0; i-- {
		out += string(s[i])
	}
	return out
}

func applyMask(mask string, input int) int {
	result := input
	m := reverseString(mask)
	for i := range m {
		if m[i] == 'X' {
			continue
		}
		var val int
		if m[i] == '1' {
			val = 1
		}
		result = result&^(1<<i) | val<<i
	}
	return int(result)
}

func applyFloatingBitmask(mask string, in int) string {
	bin := fmt.Sprintf("%.36b", in)
	if len(bin) != len(mask) {
		panic("Not going to figure out padding right now")
	}

	var result string
	for i := range bin {
		switch mask[i] {
		case '0':
			result += string(bin[i])
		case '1':
			result += "1"
		case 'X':
			result += "X"
		}
	}
	return result
}

func generateAllMasks(mask string) []int64 {
	var positions []int
	for i := range mask {
		if mask[i] == 'X' {
			positions = append(positions, len(mask)-1-i)
		}
	}

	seen := map[int64]struct{}{}
	startingMask := strings.ReplaceAll(mask, "X", "0")
	startingNum, err := strconv.ParseInt(startingMask, 2, 64)
	if err != nil {
		panic(err)
	}
	queue := []int64{startingNum}
	for len(queue) > 0 {
		current := queue[0]
		for _, pos := range positions {
			result := current | (1 << pos)
			if _, ok := seen[result]; !ok {
				seen[result] = struct{}{}
				queue = append(queue, result)
			}
		}
		queue = queue[1:]
	}
	results := []int64{startingNum}
	for num := range seen {
		results = append(results, num)
	}
	sort.Slice(results, func(i, j int) bool {
		return results[i] < results[j]
	})
	return results
}

func runFloatingProgram(in string) (int, error) {
	r := strings.NewReader(in)
	s := bufio.NewScanner(r)

	var mask string
	var memory [67682591465 + 1]int

	for s.Scan() {
		line := s.Text()
		if strings.HasPrefix(line, "mask") {
			newMask := strings.TrimSpace(strings.Split(line, "=")[1])
			mask = newMask
			continue
		}

		if strings.HasPrefix(line, "mem") {
			matched := memSetRegex.FindStringSubmatch(line)

			value, err := strconv.Atoi(matched[2])
			if err != nil {
				return 0, err
			}

			location, err := strconv.Atoi(matched[1])
			if err != nil {
				return 0, err
			}

			floatingMask := applyFloatingBitmask(mask, location)
			positions := generateAllMasks(floatingMask)
			for _, position := range positions {
				memory[position] = value
			}
		}
	}

	var sum int
	for i := range memory {
		sum += memory[i]
	}
	return sum, nil
}

func main() {
	partOneFunc := func() error {
		in, err := ioutil.ReadFile("input")
		if err != nil {
			return err
		}
		result, err := runProgram(string(in))
		if err != nil {
			return err
		}
		fmt.Println(result)
		return nil
	}

	partTwoFunc := func() error {
		in, err := ioutil.ReadFile("input")
		if err != nil {
			return err
		}

		result, err := runFloatingProgram(string(in))
		if err != nil {
			return err
		}

		fmt.Println(result)
		return nil
	}

	challenge.Run(partOneFunc, partTwoFunc)
}
