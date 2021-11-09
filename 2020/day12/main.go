package main

import (
	"errors"
	"fmt"
	"math"
	"strconv"

	"github.com/while1malloc0/advent-of-code/2020/challenge"
)

type directionType = string

const (
	FORWARD directionType = "F"
	RIGHT   directionType = "R"
	LEFT    directionType = "L"

	NORTH directionType = "N"
	SOUTH directionType = "S"
	EAST  directionType = "E"
	WEST  directionType = "W"
)

var errInvalidDirection = errors.New("Invalid direction")

type vector struct {
	direction directionType
	magnitude int
}

func parseVector(raw string) (vector, error) {
	rawDirection := raw[0]
	var direction directionType
	switch rawDirection {
	case 'F':
		direction = FORWARD
	case 'R':
		direction = RIGHT
	case 'L':
		direction = LEFT
	case 'N':
		direction = NORTH
	case 'S':
		direction = SOUTH
	case 'E':
		direction = EAST
	case 'W':
		direction = WEST
	default:
		return vector{}, errInvalidDirection
	}

	rawMagnitude := string(raw[1:])
	magnitude, err := strconv.Atoi(rawMagnitude)
	if err != nil {
		return vector{}, err
	}

	return vector{direction: direction, magnitude: magnitude}, nil
}

type waypoint struct {
	x int
	y int
}

type strategyType = int

const (
	strategySelf strategyType = iota
	strategyWaypoint
)

type boat struct {
	y        int
	x        int
	bearing  directionType
	waypoint *waypoint

	strategy strategyType
}

func (b *boat) manhattan() int {
	return int(math.Abs(float64(b.x)) + math.Abs(float64(b.y)))
}

func (b *boat) rotate(degrees int) {
	delta := int(math.Abs(float64(degrees)))
	directions := []directionType{EAST, SOUTH, WEST, NORTH}
	var index int
	for i := range directions {
		if b.bearing == directions[i] {
			index = i
		}
	}

	for delta > 0 {
		// rotate right if positive degrees, otherwise left
		if degrees >= 0 {
			index++
		} else {
			index--
		}
		index = challenge.ActualMod(index, len(directions))
		delta -= 90
		b.bearing = directions[index]
	}
}

func (b *boat) moveSelf(instruction vector) {
	dir, mag := instruction.direction, instruction.magnitude
	switch dir {
	case FORWARD:
		b.move(vector{direction: b.bearing, magnitude: mag})
	case RIGHT:
		b.rotate(mag)
	case LEFT:
		b.rotate(-mag)
	case NORTH:
		b.y -= mag
	case SOUTH:
		b.y += mag
	case EAST:
		b.x += mag
	case WEST:
		b.x -= mag
	}
}

func (b *boat) moveTowardWaypoint(num int) {
	for i := 0; i < num; i++ {
		b.x += b.waypoint.x
		b.y += b.waypoint.y
	}
}

func (b *boat) rotateWaypoint(degrees int) {
	delta := int(math.Abs(float64(degrees)))

	for delta > 0 {
		// rotate right if positive, else left
		if degrees > 0 {
			b.waypoint.x, b.waypoint.y = b.waypoint.y, b.waypoint.x
			b.waypoint.x *= -1
		} else {
			b.waypoint.x, b.waypoint.y = b.waypoint.y, b.waypoint.x
			b.waypoint.y *= -1
		}
		delta -= 90
	}
}

func (b *boat) moveRelativeToWaypoint(instruction vector) {
	dir, mag := instruction.direction, instruction.magnitude
	switch dir {
	case FORWARD:
		b.moveTowardWaypoint(mag)
	case NORTH:
		b.waypoint.y -= mag
	case SOUTH:
		b.waypoint.y += mag
	case EAST:
		b.waypoint.x += mag
	case WEST:
		b.waypoint.x -= mag
	case RIGHT:
		b.rotateWaypoint(mag)
	case LEFT:
		b.rotateWaypoint(-mag)
	}
}

func (b *boat) move(instruction vector) {
	if b.strategy == strategySelf {
		b.moveSelf(instruction)
	} else if b.strategy == strategyWaypoint {
		b.moveRelativeToWaypoint(instruction)
	}
}

func main() {
	var vectors []vector
	challenge.InputScanFunc("input", func(s string) error {
		v, err := parseVector(s)
		if err != nil {
			return err
		}
		vectors = append(vectors, v)
		return nil
	})

	partOneFunc := func() error {
		b := &boat{x: 0, y: 0, bearing: EAST}
		for i := range vectors {
			b.move(vectors[i])
		}
		fmt.Println(b.manhattan())
		return nil
	}

	partTwoFunc := func() error {
		b := &boat{x: 0, y: 0, waypoint: &waypoint{x: 10, y: -1}, strategy: strategyWaypoint}
		for i := range vectors {
			b.move(vectors[i])
		}
		fmt.Println(b.manhattan())
		return nil
	}

	challenge.Run(partOneFunc, partTwoFunc)
}
