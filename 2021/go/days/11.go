package days

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Coord [2]int

func (c Coord) X() int {
	return c[0]
}

func (c Coord) Y() int {
	return c[1]
}

type OctopusMap map[Coord]int

func (o OctopusMap) Get(x, y int) int {
	return o[Coord{x, y}]
}

func (o OctopusMap) Set(x, y, val int) {
	o[Coord{x, y}] = val
}

func (o OctopusMap) Inc() {
	for k, v := range o {
		o[k] = v + 1
	}
}

func (o OctopusMap) Flash() {
	stack := []Coord{}
	visited := []Coord{}

	for k, v := range o {
		if v > 9 {
			stack = append(stack, k)
		}
	}
}

func Day11Part1(input string) (interface{}, error) {
	coords, err := parseCoords(input)
	if err != nil {
		return nil, err
	}

	fmt.Println(coords)
	return 0, errors.New("not finished")
}

func parseCoords(input string) (OctopusMap, error) {
	coords := make(OctopusMap)
	for y, line := range strings.Split(input, "\n") {
		for x, val := range line {
			parsed, err := strconv.Atoi(string(val))
			if err != nil {
				return nil, err
			}
			coords[[2]int{x, y}] = parsed
		}
	}
	return coords, nil
}

func neighbors(x, y int) []Coord {
	candidates := []Coord{
		// top left
		{x - 1, y - 1},
		// up
		{x, y - 1},
		// top right
		{x + 1, y - 1},
		// left
		{x - 1, y},
		// right
		{x + 1, y},
		// bottom left
		{x - 1, y + 1},
		// bottom
		{x, y + 1},
		// bottom right
		{x + 1, y + 1},
	}

	result := []Coord{}
	for _, c := range candidates {
		if c.X() < 0 || c.Y() < 0 || c.X() > 9 || c.Y() > 9 {
			continue
		}
		result = append(result, c)
	}
	return result
}
