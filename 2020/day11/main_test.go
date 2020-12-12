package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGridString(t *testing.T) {
	input := &grid{
		width:  4,
		height: 4,
		cells: [][]cell{
			{SEAT, FLOOR, SEAT, SEAT},
			{SEAT, FLOOR, SEAT, SEAT},
			{SEAT, OCCUPIED, SEAT, FLOOR},
			{SEAT, OCCUPIED, SEAT, FLOOR},
		},
	}
	want := `L.LL
L.LL
L#L.
L#L.
`
	got := input.String()
	assert.Equal(t, want, got, "Expected grid %#v to print to:\n%s\nBut got\n\n%s", input, want, got)
}

func TestGetNextCellState(t *testing.T) {
	testCases := []struct {
		desc      string
		cell      cell
		neighbors []cell
		want      cell
	}{
		{
			"floors don't change",
			FLOOR,
			[]cell{SEAT, FLOOR, OCCUPIED, SEAT},
			FLOOR,
		},
		{
			"empty seats with no occupied seats become occupied",
			SEAT,
			[]cell{SEAT, FLOOR, SEAT, SEAT},
			OCCUPIED,
		},
		{
			"empty seats with 1+ occupied seats stay the same",
			SEAT,
			[]cell{SEAT, OCCUPIED, SEAT, SEAT},
			SEAT,
		},
		{
			"occupied seats with 4+ occupied seats become empty",
			OCCUPIED,
			[]cell{OCCUPIED, OCCUPIED, OCCUPIED, OCCUPIED},
			SEAT,
		},
		{
			"occupied seats with 3 or fewer occupied seats stay occupied",
			OCCUPIED,
			[]cell{OCCUPIED, OCCUPIED, OCCUPIED, SEAT},
			OCCUPIED,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := getNextCellState(tC.cell, tC.neighbors)
			assert.Equal(t, tC.want, got)
		})
	}
}

func TestNewGridFromString(t *testing.T) {
	input := `L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL
`
	got := newGridFromString(input)
	want := &grid{
		width:  10,
		height: 10,
		cells: [][]cell{
			{SEAT, FLOOR, SEAT, SEAT, FLOOR, SEAT, SEAT, FLOOR, SEAT, SEAT},
			{SEAT, SEAT, SEAT, SEAT, SEAT, SEAT, SEAT, FLOOR, SEAT, SEAT},
			{SEAT, FLOOR, SEAT, FLOOR, SEAT, FLOOR, FLOOR, SEAT, FLOOR, FLOOR},
			{SEAT, SEAT, SEAT, SEAT, FLOOR, SEAT, SEAT, FLOOR, SEAT, SEAT},
			// Rest of cells elided for my sanity
		},
	}
	assert.Equal(t, got.height, want.height)
	assert.Equal(t, got.width, want.width)
	assert.Equal(t, got.cells[0], want.cells[0])
	assert.Equal(t, got.cells[1], want.cells[1])
	assert.Equal(t, got.cells[2], want.cells[2])
	assert.Equal(t, got.cells[3], want.cells[3])
}

func TestGridGetNeighbors(t *testing.T) {
	input := `L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL
`
	g := newGridFromString(input)
	got := g.neighbors(0, 0)
	want := []cell{SEAT, SEAT, FLOOR, SEAT, FLOOR, SEAT, SEAT, SEAT}
	assert.Equal(t, want, got)
}

// Test to see if compensating for Go's broken mod function broke things
func TestGridGetNeighbors_middle(t *testing.T) {
	input := `L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL
`
	g := newGridFromString(input)
	got := g.neighbors(1, 5)
	want := []cell{FLOOR, SEAT, SEAT, SEAT, SEAT, SEAT, FLOOR, FLOOR}
	assert.Equal(t, want, got)
}

func TestGridTick(t *testing.T) {
	input := `L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL`
	// 46 == .
	// 76 == L
	// 35 == #

	// L.LL.LL.LL
	// LLLLLLL.LL
	// L.L.L..L..
	// LLLL.LL.LL
	// L.LL.LL.LL
	// L.LLLLL.LL
	// ..L.L.....
	// LLLLLLLLLL
	// L.LLLLLL.L
	// L.LLLLL.LL

	g := newGridFromString(input)
	g.tick()
	got := g.String()
	want := `#.##.##.##
#######.##
#.#.#..#..
####.##.##
#.##.##.##
#.#####.##
..#.#.....
##########
#.######.#
#.#####.##
`
	assert.Equal(t, want, got)
}

func TestGridTick_two(t *testing.T) {
	input := `#.##.##.##
#######.##
#.#.#..#..
####.##.##
#.##.##.##
#.#####.##
..#.#.....
##########
#.######.#
#.#####.##`

	// 46 == .
	// 76 == L
	// 35 == #

	// L.LL.LL.LL
	// LLLLLLL.LL
	// L.L.L..L..
	// LLLL.LL.LL
	// L.LL.LL.LL
	// L.LLLLL.LL
	// ..L.L.....
	// LLLLLLLLLL
	// L.LLLLLL.L
	// L.LLLLL.LL

	g := newGridFromString(input)
	g.tick()
	got := g.String()
	want := `#.LL.L#.##
#LLLLLL.L#
L.L.L..L..
#LLL.LL.L#
#.LL.LL.LL
#.LLLL#.##
..L.L.....
#LLLLLLLL#
#.LLLLLL.L
#.#LLLL.##`
	assert.Equal(t, want, got)
}
