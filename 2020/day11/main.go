package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/while1malloc0/advent-of-code/2020/challenge"
)

type cell = rune

const (
	SEAT     cell = 'L'
	FLOOR    cell = '.'
	OCCUPIED cell = '#'
)

type grid struct {
	width, height int
	cells         [][]cell
}

func newGridFromString(in string) *grid {
	out := &grid{}
	r := strings.NewReader(in)
	s := bufio.NewScanner(r)
	for s.Scan() {
		row := s.Text()
		var cellRow []cell
		for _, char := range row {
			cellRow = append(cellRow, char)
		}
		out.cells = append(out.cells, cellRow)
	}
	out.width = len(out.cells[0])
	out.height = len(out.cells)
	return out
}

// Go's mod behavior is different than basically every other languages. There's
// discussions as to why on the internet. They're too pedantic to be linked too.
func actualMod(a, b int) int {
	return ((a % b) + b) % b
}

func (g *grid) neighbors(y, x int) []cell {
	type coordinate struct {
		x, y int
	}

	var cells []cell

	// diag left
	if (y-1 >= 0) && (x-1 >= 0) {
		cells = append(cells, g.cells[y-1][x-1])
	}

	// up
	if y-1 >= 0 {
		cells = append(cells, g.cells[y-1][x])
	}

	// diag up right
	if (y-1 >= 0) && (x+1 < g.width) {
		cells = append(cells, g.cells[y-1][x+1])
	}

	// left
	if x-1 >= 0 {
		cells = append(cells, g.cells[y][x-1])
	}

	// right
	if x+1 < g.width {
		cells = append(cells, g.cells[y][x+1])
	}

	// diag down left
	if (x-1 >= 0) && (y+1 < g.height) {
		cells = append(cells, g.cells[y+1][x-1])
	}

	// down
	if y+1 < g.height {
		cells = append(cells, g.cells[y+1][x])
	}

	// diag down right
	if (y+1 < g.height) && (x+1 < g.width) {
		cells = append(cells, g.cells[y+1][x+1])
	}
	// diagUpLeft := coordinate{x: actualMod(x-1, g.width), y: actualMod(y-1, g.height)}
	// up := coordinate{x: x, y: actualMod(y-1, g.height)}
	// diagUpRight := coordinate{x: actualMod(x+1, g.width), y: actualMod(y-1, g.height)}
	// left := coordinate{x: actualMod(x-1, g.width), y: y}
	// right := coordinate{x: actualMod(x+1, g.width), y: y}
	// diagDownLeft := coordinate{x: actualMod(x-1, g.width), y: actualMod(y+1, g.height)}
	// down := coordinate{x: x, y: actualMod(y+1, g.height)}
	// diagDownRight := coordinate{x: actualMod(x+1, g.width), y: actualMod(y+1, g.height)}

	return cells
	// return []cell{
	// 	g.cells[diagUpLeft.y][diagUpLeft.x],
	// 	g.cells[up.y][up.x],
	// 	g.cells[diagUpRight.y][diagUpRight.x],
	// 	g.cells[left.y][left.x],
	// 	g.cells[right.y][right.x],
	// 	g.cells[diagDownLeft.y][diagDownLeft.x],
	// 	g.cells[down.y][down.x],
	// 	g.cells[diagDownRight.y][diagDownRight.x],
	// }
}

func (g *grid) tick() {
	var nextCells [][]cell
	for y := 0; y < g.height; y++ {
		var row []cell
		for x := 0; x < g.width; x++ {
			neighbors := g.neighbors(y, x)
			nextState := getNextCellState(g.cells[y][x], neighbors)
			row = append(row, nextState)
		}
		nextCells = append(nextCells, row)
	}
	g.cells = nextCells
}

func getNextCellState(c cell, neighbors []cell) cell {
	switch c {
	case FLOOR:
		// Floors don't change
		return FLOOR
	case SEAT:
		// Empty seats with no occupied adjacent seats become occupied,
		// otherwise stay a seat
		var numOccupied int
		for _, neighbor := range neighbors {
			if neighbor == OCCUPIED {
				numOccupied++
			}
		}
		if numOccupied == 0 {
			return OCCUPIED
		}
		return SEAT
	case OCCUPIED:
		var numOccupied int
		for _, neighbor := range neighbors {
			if neighbor == OCCUPIED {
				numOccupied++
			}
		}
		if numOccupied >= 4 {
			return SEAT
		}
		return OCCUPIED
	}
	return FLOOR
}

func (g *grid) String() string {
	var b strings.Builder

	for y := 0; y < g.height; y++ {
		for x := 0; x < g.width; x++ {
			b.WriteRune(g.cells[y][x])
		}
		b.WriteString("\n")
	}

	return strings.TrimSpace(b.String())
}

func main() {
	partOneFunc := func() error {
		in, err := ioutil.ReadFile("input")
		if err != nil {
			return err
		}
		g := newGridFromString(string(in))
		// g.tick()
		// g.tick()
		for {
			lastState := g.String()
			g.tick()
			if g.String() == lastState {
				break
			}
		}
		var count int
		for y := 0; y < g.height; y++ {
			for x := 0; x < g.width; x++ {
				if g.cells[y][x] == OCCUPIED {
					count++
				}
			}
		}
		fmt.Println(count)
		return nil
	}

	challenge.Run(partOneFunc, nil)
}
