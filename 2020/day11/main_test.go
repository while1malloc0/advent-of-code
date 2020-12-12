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
L#L.`
	got := input.String()
	assert.Equal(t, want, got, "Expected grid %#v to print to:\n%s\nBut got\n\n%s", input, want, got)
}

func TestGetNextCellState(t *testing.T) {
	testCases := []struct {
		desc        string
		cell        cell
		neighbors   []cell
		maxOccupied int
		want        cell
	}{
		{
			"floors don't change",
			FLOOR,
			[]cell{SEAT, FLOOR, OCCUPIED, SEAT},
			4,
			FLOOR,
		},
		{
			"empty seats with no occupied seats become occupied",
			SEAT,
			[]cell{SEAT, FLOOR, SEAT, SEAT},
			4,
			OCCUPIED,
		},
		{
			"empty seats with 1+ occupied seats stay the same",
			SEAT,
			[]cell{SEAT, OCCUPIED, SEAT, SEAT},
			4,
			SEAT,
		},
		{
			"occupied seats with N+ occupied seats become empty. N = min occupied",
			OCCUPIED,
			[]cell{OCCUPIED, OCCUPIED, OCCUPIED, OCCUPIED},
			4,
			SEAT,
		},
		{
			"occupied seats with N+ occupied seats become empty. N = min occupied",
			OCCUPIED,
			[]cell{OCCUPIED, OCCUPIED, OCCUPIED, OCCUPIED},
			5,
			OCCUPIED,
		},
		{
			"occupied seats with N+ occupied seats become empty. N = min occupied",
			OCCUPIED,
			[]cell{OCCUPIED, OCCUPIED, OCCUPIED, OCCUPIED, OCCUPIED},
			5,
			SEAT,
		},
		{
			"occupied seats with 3 or fewer occupied seats stay occupied",
			OCCUPIED,
			[]cell{OCCUPIED, OCCUPIED, OCCUPIED, SEAT},
			4,
			OCCUPIED,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := getNextCellState(tC.cell, tC.neighbors, tC.maxOccupied)
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
	got := newGrid(input, 4, nearestNeighbors)
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
L.LLLLL.LL`
	g := newGrid(input, 4, nearestNeighbors)
	got := g.neighbors(0, 0)
	want := []cell{FLOOR, SEAT, SEAT}
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
L.LLLLL.LL`
	g := newGrid(input, 4, nearestNeighbors)
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

	g := newGrid(input, 4, nearestNeighbors)
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
#.#####.##`
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

	g := newGrid(input, 4, nearestNeighbors)
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

func TestLineOfSightNeighbors(t *testing.T) {
	input := `.......#.
...#.....
.#.......
.........
..#L....#
....#....
.........
#........
...#.....`
	g := newGrid(input, 4, lineOfSightNeighbors)
	got := g.neighbors(4, 3)
	want := []cell{OCCUPIED, OCCUPIED, OCCUPIED, OCCUPIED, OCCUPIED, OCCUPIED, OCCUPIED, OCCUPIED}
	assert.Equal(t, want, got)
}

func TestLineOfSightNeighbors_empty(t *testing.T) {
	input := `.##.##.
#.#.#.#
##...##
...L...
##...##
#.#.#.#
.##.##.`
	g := newGrid(input, 4, lineOfSightNeighbors)
	got := g.neighbors(3, 3)
	var want []cell
	assert.Equal(t, want, got)
}

func TestTickMaxOccupied(t *testing.T) {
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
	first := `#.##.##.##
#######.##
#.#.#..#..
####.##.##
#.##.##.##
#.#####.##
..#.#.....
##########
#.######.#
#.#####.##`

	second := `#.LL.LL.L#
#LLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLL#
#.LLLLLL.L
#.LLLLL.L#`

	g := newGrid(input, 5, lineOfSightNeighbors)
	g.tick()
	assert.Equal(t, first, g.String())
	g.tick()
	assert.Equal(t, second, g.String())
}
