package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBoatManhattanDistance(t *testing.T) {
	b := &boat{x: 10, y: 10}
	got := b.manhattan()
	want := 20
	assert.Equal(t, want, got)
}

func TestParseDirection(t *testing.T) {
	testCases := []struct {
		desc    string
		in      string
		want    vector
		wantErr error
	}{
		{"forward 10", "F10", vector{direction: FORWARD, magnitude: 10}, nil},
		{"right 90 degrees", "R90", vector{direction: RIGHT, magnitude: 90}, nil},
		{"north 3", "N3", vector{direction: NORTH, magnitude: 3}, nil},
		{"south 10", "S10", vector{direction: SOUTH, magnitude: 10}, nil},
		{"east 10", "E10", vector{direction: EAST, magnitude: 10}, nil},
		{"west 10", "W10", vector{direction: WEST, magnitude: 10}, nil},
		{"left 90", "L90", vector{direction: LEFT, magnitude: 90}, nil},
		{"not real", "Z10", vector{}, errInvalidDirection},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got, err := parseVector(tC.in)
			assert.Equal(t, tC.wantErr, err)
			assert.Equal(t, tC.want, got, "Expected string %s to parse to vector %#v, got %#v", tC.in, tC.want, got)
		})
	}
}

func TestBoatMove(t *testing.T) {
	testCases := []struct {
		desc string
		in   vector
		want boat
	}{
		{"forward 10", vector{FORWARD, 10}, boat{x: 10, y: 0, bearing: EAST}},
		{"north 10", vector{NORTH, 10}, boat{x: 0, y: -10, bearing: EAST}},
		{"south 10", vector{SOUTH, 10}, boat{x: 0, y: 10, bearing: EAST}},
		{"east 10", vector{EAST, 10}, boat{x: 10, y: 0, bearing: EAST}},
		{"west 10", vector{WEST, 10}, boat{x: -10, y: 0, bearing: EAST}},
		{"right 90", vector{RIGHT, 90}, boat{x: 0, y: 0, bearing: SOUTH}},
		{"left 90", vector{LEFT, 90}, boat{x: 0, y: 0, bearing: NORTH}},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			b := &boat{x: 0, y: 0, bearing: EAST}
			b.move(tC.in)
			assert.Equal(t, tC.want.x, b.x)
			assert.Equal(t, tC.want.y, b.y)
			assert.Equal(t, tC.want.bearing, b.bearing)
		})
	}
}

func TestBoatRotate(t *testing.T) {
	testCases := []struct {
		desc string
		in   int
		want directionType
	}{
		{"90 degrees", 90, SOUTH},
		{"180 degrees", 180, WEST},
		{"270 degrees", 270, NORTH},
		{"360 degrees", 360, EAST},
		{"420 degrees", 420, SOUTH},
		{"90 degrees left", -90, NORTH},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			b := &boat{bearing: EAST, x: 0, y: 0}
			b.rotate(tC.in)
			assert.Equal(t, tC.want, b.bearing)
		})
	}
}

func TestExample(t *testing.T) {
	in := []string{
		"F10",
		"N3",
		"F7",
		"R90",
		"F11",
	}
	var vectors []vector
	for i := range in {
		v, err := parseVector(in[i])
		assert.Nil(t, err)
		vectors = append(vectors, v)
	}
	b := &boat{x: 0, y: 0, bearing: EAST}
	for i := range vectors {
		b.move(vectors[i])
	}
	assert.Equal(t, b.manhattan(), 25)
}

func TestMoveBoatWaypoint(t *testing.T) {
	testCases := []struct {
		desc  string
		input vector
		want  *boat
	}{
		{
			"forward moves towards the waypoint",
			vector{FORWARD, 10},
			&boat{x: 100, y: -10, waypoint: &waypoint{y: -1, x: 10}, strategy: strategyWaypoint},
		},
		{
			"north moves the waypoint north",
			vector{NORTH, 10},
			&boat{x: 0, y: 0, waypoint: &waypoint{y: -11, x: 10}, strategy: strategyWaypoint},
		},
		{
			"south moves the waypoint south",
			vector{SOUTH, 10},
			&boat{x: 0, y: 0, waypoint: &waypoint{y: 9, x: 10}, strategy: strategyWaypoint},
		},
		{
			"east moves the waypoint east",
			vector{EAST, 10},
			&boat{x: 0, y: 0, waypoint: &waypoint{y: -1, x: 20}, strategy: strategyWaypoint},
		},
		{
			"west moves the waypoint west",
			vector{WEST, 10},
			&boat{x: 0, y: 0, waypoint: &waypoint{y: -1, x: 0}, strategy: strategyWaypoint},
		},
		{
			"right rotates the waypoint clockwise",
			vector{RIGHT, 90},
			&boat{x: 0, y: 0, waypoint: &waypoint{y: -10, x: -1}, strategy: strategyWaypoint},
		},
		{
			"left rotates the waypoint counter-clockwise",
			vector{LEFT, 90},
			&boat{x: 0, y: 0, waypoint: &waypoint{y: 10, x: 1}, strategy: strategyWaypoint},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			b := &boat{x: 0, y: 0, waypoint: &waypoint{y: -1, x: 10}, strategy: strategyWaypoint}
			b.move(tC.input)
			assert.Equal(t, tC.want.x, b.x)
			assert.Equal(t, tC.want.y, b.y)
			assert.Equal(t, tC.want.waypoint.x, b.waypoint.x)
			assert.Equal(t, tC.want.waypoint.y, b.waypoint.y)
		})
	}
}

func TestExamplePartTwo(t *testing.T) {
	input := []string{
		"F10",
		"N3",
		"F7",
		"R90",
		"F11",
	}
	b := &boat{x: 0, y: 0, waypoint: &waypoint{x: 10, y: -1}, strategy: strategyWaypoint}
	for i := range input {
		v, err := parseVector(input[i])
		assert.Nil(t, err)
		b.move(v)
	}
	assert.Equal(t, 286, b.manhattan())
}
