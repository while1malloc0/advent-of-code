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

type neighborRulesFunc = func(*grid, int, int) []cell

type grid struct {
	width, height int
	cells         [][]cell

	neighborRules neighborRulesFunc
	maxOccupied   int
}

func newGrid(raw string, maxOccupied int, neighborRules neighborRulesFunc) *grid {
	out := &grid{
		neighborRules: neighborRules,
		maxOccupied:   maxOccupied,
	}
	r := strings.NewReader(raw)
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

func (g *grid) neighbors(y, x int) []cell {
	return g.neighborRules(g, y, x)
}

func nearestNeighbors(g *grid, y, x int) []cell {
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

	return cells
}

func lineOfSightNeighbors(g *grid, y, x int) []cell {
	var cells []cell

	// diag up left
	dy, dx := y-1, x-1
	for dy >= 0 && dx >= 0 {
		if g.cells[dy][dx] != FLOOR {
			cells = append(cells, g.cells[dy][dx])
			break
		}
		dy--
		dx--
	}

	// up
	dy = y - 1
	for dy >= 0 {
		if g.cells[dy][x] != FLOOR {
			cells = append(cells, g.cells[dy][x])
			break
		}
		dy--
	}

	// diag up right
	dy, dx = y-1, x+1
	for dy >= 0 && dx < g.width {
		if g.cells[dy][dx] != FLOOR {
			cells = append(cells, g.cells[dy][dx])
			break
		}
		dy--
		dx++
	}

	// left
	dx = x - 1
	for dx >= 0 {
		if g.cells[y][dx] != FLOOR {
			cells = append(cells, g.cells[y][dx])
			break
		}
		dx--
	}

	// right
	dx = x + 1
	for dx < g.width {
		if g.cells[y][dx] != FLOOR {
			cells = append(cells, g.cells[y][dx])
			break
		}
		dx++
	}

	// diag left down
	dy, dx = y+1, x-1
	for dy < g.height && dx >= 0 {
		if g.cells[dy][dx] != FLOOR {
			cells = append(cells, g.cells[dy][dx])
			break
		}
		dy++
		dx--
	}

	// down
	dy = y + 1
	for dy < g.height {
		if g.cells[dy][x] != FLOOR {
			cells = append(cells, g.cells[dy][x])
			break
		}
		dy++
	}

	// diag down right
	dy, dx = y+1, x+1
	for dy < g.height && dx < g.width {
		if g.cells[dy][dx] != FLOOR {
			cells = append(cells, g.cells[dy][dx])
			break
		}
		dy++
		dx++
	}

	return cells
}

func (g *grid) tick() {
	var nextCells [][]cell
	for y := 0; y < g.height; y++ {
		var row []cell
		for x := 0; x < g.width; x++ {
			neighbors := g.neighbors(y, x)
			nextState := getNextCellState(g.cells[y][x], neighbors, g.maxOccupied)
			row = append(row, nextState)
		}
		nextCells = append(nextCells, row)
	}
	g.cells = nextCells
}

func getNextCellState(c cell, neighbors []cell, maxOccupied int) cell {
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
		if numOccupied >= maxOccupied {
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
		g := newGrid(string(in), 4, nearestNeighbors)
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

	partTwoFunc := func() error {
		in, err := ioutil.ReadFile("input")
		if err != nil {
			return err
		}
		g := newGrid(string(in), 5, lineOfSightNeighbors)
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

	challenge.Run(partOneFunc, partTwoFunc)
}
