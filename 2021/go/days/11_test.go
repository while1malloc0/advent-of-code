package days

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDay11Part1E2E(t *testing.T) {
	input := `
5483143223
2745854711
5264556173
6141336146
6357385478
4167524645
2176841721
6882881134
4846848554
5283751526
	`
	input = strings.TrimSpace(input)

	got, err := Day11Part1(input)
	require.Nil(t, err)

	want := 1656
	assert.Equal(t, got, want)
}

func TestParseCoords(t *testing.T) {
	input := `
5483143223
2745854711
5264556173
6141336146
6357385478
4167524645
2176841721
6882881134
4846848554
5283751526
	`
	input = strings.TrimSpace(input)

	got, err := parseCoords(input)
	require.Nil(t, err)

	assert.Equal(t, got[[2]int{0, 0}], 5)
}

func TestGetNeighbors(t *testing.T) {
	got := neighbors(0, 0)
	assert.Equal(t, got, []Coord{
		{1, 0},
		{0, 1},
		{1, 1},
	})
}

func TestOctopusMapInc(t *testing.T) {
	subject := OctopusMap{
		{0, 0}: 1,
		{0, 1}: 2,
		{1, 0}: 3,
	}
	subject.Inc()
	assert.Equal(t, subject.Get(0, 0), 2)
	assert.Equal(t, subject.Get(0, 1), 3)
	assert.Equal(t, subject.Get(1, 0), 4)
}

func TestFlash(t *testing.T) {
	input := `
11111
19991
19191
19991
11111
	`
	input = strings.TrimSpace(input)
	subject, err := parseCoords(input)
	require.Nil(t, err)

	subject.Inc()
	subject.Flash()

	next := `
34543
40004
50005
40004
34543
	`
	next = strings.TrimSpace(next)
	want, err := parseCoords(next)
	require.Nil(t, err)

	assert.Equal(t, want, subject)
}
