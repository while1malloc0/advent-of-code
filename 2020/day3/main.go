package main

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/while1malloc0/advent-of-code/2020/challenge"
)

func isTree(char rune) bool {
	return char == '#'
}

func parseSkiMapRow(row string) []rune {
	return []rune(row)
}

func parseSkiMap(rows string) [][]rune {
	var out [][]rune
	r := strings.NewReader(rows)
	s := bufio.NewScanner(r)
	for s.Scan() {
		out = append(out, parseSkiMapRow(s.Text()))
	}
	return out
}

func countSkiMapTrees(in [][]rune, rise, run int) int {
	var count, horizontal, vertical int
	mapLen := len(in[0])
	for vertical < len(in) {
		if isTree(in[vertical][horizontal]) {
			count++
		}
		horizontal += run
		horizontal %= mapLen
		vertical += rise
	}
	return count
}

func main() {
	var rows [][]rune
	challenge.InputScanFunc("input", func(s string) error {
		rows = append(rows, parseSkiMapRow(s))
		return nil
	})

	partOneFunc := func() error {
		count := countSkiMapTrees(rows, 1, 3)
		fmt.Println(count)
		return nil
	}

	partTwoFunc := func() error {
		params := [][]int{
			{1, 1},
			{3, 1},
			{5, 1},
			{7, 1},
			{1, 2},
		}
		result := 1
		for i := range params {
			count := countSkiMapTrees(rows, params[i][1], params[i][0])
			result *= count
		}
		fmt.Println(result)
		return nil
	}

	err := challenge.Run(partOneFunc, partTwoFunc)
	if err != nil {
		panic(err)
	}
}
