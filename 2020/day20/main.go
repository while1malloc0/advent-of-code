package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/while1malloc0/advent-of-code/2020/challenge"
)

type tile struct {
	id      int
	content [][]rune
	edges   []int64
}

func reverseString(s string) string {
	out := ""
	for i := len(s) - 1; i >= 0; i-- {
		out += string(s[i])
	}
	return out
}

func parseTile(in string) *tile {
	parts := strings.Split(in, ":")
	idParts := strings.Split(parts[0], " ")
	id, err := strconv.Atoi(idParts[1])

	if err != nil {
		panic(err)
	}

	tileStr := strings.TrimSpace(parts[1])
	r := strings.NewReader(tileStr)
	s := bufio.NewScanner(r)

	var rows [][]rune
	for s.Scan() {
		var row []rune
		line := s.Text()
		for _, char := range line {
			row = append(row, char)
		}
		rows = append(rows, row)
	}

	topEdgeStr := ""
	for i := range rows {
		topEdgeStr += string(rows[i][0])
	}
	topEdge := getEdge(topEdgeStr)
	topEdgeRev := getEdge(reverseString(topEdgeStr))

	rightEdgeStr := ""
	for i := range rows {
		rightEdgeStr += string(rows[len(rows)-1][i])
	}
	rightEdge := getEdge(rightEdgeStr)
	rightEdgeRev := getEdge(reverseString(rightEdgeStr))

	bottomEdgeStr := ""
	for i := range rows {
		bottomEdgeStr += string(rows[i][len(rows[i])-1])
	}
	bottomEdge := getEdge(bottomEdgeStr)
	bottomEdgeRev := getEdge(reverseString(bottomEdgeStr))

	leftEdgeStr := ""
	for i := range rows {
		leftEdgeStr += string(rows[0][i])
	}
	leftEdge := getEdge(leftEdgeStr)
	leftEdgeRev := getEdge(reverseString(leftEdgeStr))

	return &tile{id: id, content: rows, edges: []int64{topEdge, topEdgeRev, rightEdge, rightEdgeRev, bottomEdge, bottomEdgeRev, leftEdge, leftEdgeRev}}
}

func getEdge(row string) int64 {
	sb := &strings.Builder{}
	for _, char := range row {
		switch char {
		case '.':
			sb.WriteString("0")
		case '#':
			sb.WriteString("1")
		default:
			panic(fmt.Sprintf("Unknown character value: %s", string(char)))
		}
	}
	bitstring := sb.String()
	n, err := strconv.ParseInt(bitstring, 2, 64)
	if err != nil {
		panic(err)
	}
	return n
}

func isNeighbor(lhs, rhs []int64) bool {
	for i := range lhs {
		for j := range rhs {
			if lhs[i] == rhs[j] {
				return true
			}
		}
	}
	return false
}

func runExample(in string) int {
	var tiles []*tile

	parts := strings.Split(in, "\n\n")

	for i := range parts {
		tileStr := strings.TrimSpace(parts[i])
		tile := parseTile(tileStr)
		tiles = append(tiles, tile)
	}

	var corners []*tile
	for i := range tiles {
		candidate := tiles[i]
		var numNeighbors int
		for j := range tiles {
			other := tiles[j]
			if candidate.id == other.id {
				continue
			}
			if isNeighbor(candidate.edges, other.edges) {
				numNeighbors++
			}
		}
		if numNeighbors == 2 {
			corners = append(corners, candidate)
		}
	}

	result := 1
	for i := range corners {
		result *= corners[i].id
	}

	return result
}

func main() {
	partOneFunc := func() error {
		f, err := ioutil.ReadFile("input")
		if err != nil {
			return err
		}
		result := runExample(string(f))
		fmt.Println(result)
		return nil
	}

	challenge.Run(partOneFunc, nil)
}
