package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsTree(t *testing.T) {
	testCases := []struct {
		desc string
		char rune
		want bool
	}{
		{"nope", '.', false},
		{"yup", '#', true},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := isTree(tC.char)
			assert.Equal(t, got, tC.want, "Expected isTree(%#v) to be %v but got %v", tC.char, tC.want, got)
		})
	}
}

func TestParseSkiMap(t *testing.T) {
	input := "..##......."
	got := parseSkiMapRow(input)
	want := []rune{'.', '.', '#', '#', '.', '.', '.', '.', '.', '.', '.'}
	assert.Equal(t, got, want, "Expected first row of parseSkiMap(%#v) to be %v, got %v", input, want, got)
}

func TestCountSkiMapTrees(t *testing.T) {
	input := `..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#
`
	parsed := parseSkiMap(input)
	rise := 1
	run := 3
	got := countSkiMapTrees(parsed, rise, run)
	want := 7
	assert.Equal(t, want, got, "Expected to find %v trees in ski map [%#v], but found %v", want, input, got)
}
