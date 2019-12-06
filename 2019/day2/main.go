package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const (
	OPAdd  = 1
	OPMul  = 2
	OPHalt = 99
)

func main() {
	f, err := ioutil.ReadFile("input")
	if err != nil {
		panic(err)
	}
	input := strings.TrimSpace(string(f))
	fmt.Println(execute(input))
}

func execute(program string) string {
	rawTokens := strings.Split(program, ",")
	tokens := parseInts(rawTokens)
	outputTokens := tokens
MAIN_LOOP:
	for i := 0; i < len(tokens); i += 4 {
		opCode := tokens[i]
		if opCode == OPHalt {
			break MAIN_LOOP
		}
		firstInput := outputTokens[tokens[i+1]]
		secondInput := outputTokens[tokens[i+2]]
		position := tokens[i+3]
		switch opCode {
		case OPAdd:
			outputTokens[position] = firstInput + secondInput
		case OPMul:
			outputTokens[position] = firstInput * secondInput
		case OPHalt:
			break MAIN_LOOP
		}
	}

	return strings.Join(parseStrings(outputTokens), ",")
}

func parseInts(ss []string) []int64 {
	var out []int64
	for _, s := range ss {
		i, err := strconv.ParseInt(s, 0, 0)
		if err != nil {
			panic(err)
		}
		out = append(out, i)
	}
	return out
}

func parseStrings(ints []int64) []string {
	var out []string
	for _, i := range ints {
		out = append(out, fmt.Sprintf("%d", i))
	}
	return out
}
