package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseCube(t *testing.T) {
	given := ` 
.#.
..#
###
	`
	given = strings.TrimSpace(given)

	want := cube{
		coord{x: 1, y: 0, z: 0}: struct{}{},
		coord{x: 2, y: 1, z: 0}: struct{}{},
		coord{x: 0, y: 2, z: 0}: struct{}{},
		coord{x: 1, y: 2, z: 0}: struct{}{},
		coord{x: 2, y: 2, z: 0}: struct{}{},
	}

	got := parseCube(given)

	assert.Equal(t, want, got)
}

func TestNextState(t *testing.T) {
	// .#.
	// ..#
	// ###
	c := cube{
		coord{x: 1, y: 0, z: 0}: struct{}{},
		coord{x: 2, y: 1, z: 0}: struct{}{},
		coord{x: 0, y: 2, z: 0}: struct{}{},
		coord{x: 1, y: 2, z: 0}: struct{}{},
		coord{x: 2, y: 2, z: 0}: struct{}{},
	}
	testCases := []struct {
		desc  string
		given coord
		want  bool
	}{
		{"Inactive with fewer than three active neighbors remains inactive", coord{x: 0, y: 0, z: 0}, false},
		{"Inactive with three neighbors becomes active", coord{x: 2, y: 2, z: -1}, true},
		{"Active with three active neighbors stays active", coord{x: 1, y: 2, z: 0}, true},
		{"Active with fewer than two active neighbors becomes inactive", coord{x: 0, y: 2, z: 0}, false},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := nextState(tC.given, c)
			assert.Equal(t, tC.want, got)
		})
	}
}

func TestRunBoot(t *testing.T) {
	given := cube{
		coord{x: 1, y: 0, z: 0}: struct{}{},
		coord{x: 2, y: 1, z: 0}: struct{}{},
		coord{x: 0, y: 2, z: 0}: struct{}{},
		coord{x: 1, y: 2, z: 0}: struct{}{},
		coord{x: 2, y: 2, z: 0}: struct{}{},
	}
	numCycles := 6
	want := 112

	got := runBoot(given, numCycles)

	assert.Equal(t, want, got)
}
