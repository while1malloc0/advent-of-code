package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type coordinate struct {
	x, y int
}

func main() {
	assert(getAnswerOne("R8,U5,L5,D3", "U7,R6,D4,L4") == 6)
	assert(getAnswerOne("R75,D30,R83,U83,L12,D49,R71,U7,L72", "U62,R66,U55,R34,D71,R55,D58,R83") == 159)
	assert(getAnswerOne("R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51", "U98,R91,D20,R16,D67,R40,U7,R15,U6,R7") == 135)

	f, err := os.Open("input")
	if err != nil {
		panic(err)

	}
	r := bufio.NewReader(f)
	line1, err := r.ReadString(byte('\n'))
	if err != nil {
		panic(err)

	}
	line2, err := r.ReadString(byte('\n'))
	if err != nil {
		panic(err)

	}
	// Part 1
	// fmt.Println(getAnswerOne(line1, line2))
	// Part 2
	assert(getAnswerTwo("R8,U5,L5,D3", "U7,R6,D4,L4") == 30)
	assert(getAnswerTwo("R75,D30,R83,U83,L12,D49,R71,U7,L72", "U62,R66,U55,R34,D71,R55,D58,R83") == 610)
	assert(getAnswerTwo("R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51", "U98,R91,D20,R16,D67,R40,U7,R15,U6,R7") == 410)
	fmt.Println(getAnswerTwo(line1, line2))
}

func assert(predicate bool) {
	if predicate != true {
		panic("Untrue assertion")
	}
}

func getAnswerOne(s1, s2 string) int {
	firstWire := s1
	secondWire := s2

	firstWirePoints := parsePoints(firstWire)
	secondWirePoints := parsePoints(secondWire)
	crosses := crossed(firstWirePoints, secondWirePoints)
	answer := findMinManhattan(crosses)
	return answer
}

func getAnswerTwo(s1, s2 string) int {
	firstWire := s1
	secondWire := s2

	firstWirePoints := parsePoints(firstWire)
	secondWirePoints := parsePoints(secondWire)
	crosses := crossed(firstWirePoints, secondWirePoints)
	answer := findMinStep(firstWirePoints, secondWirePoints, crosses)
	return answer
}

func findMinStep(a, b, cs map[coordinate]int) int {
	lowest := math.MaxInt64
	for k := range cs {
		sum := a[k] + b[k]
		if sum <= lowest {
			lowest = sum
		}
	}
	return lowest
}

func parsePoints(line string) map[coordinate]int {
	line = strings.TrimSpace(line)
	points := map[coordinate]int{}
	x := 0
	y := 0
	step := 0
	for _, l := range strings.Split(line, ",") {
		direction := string(l[0])
		n, err := strconv.Atoi(string(l[1:]))
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
		switch direction {
		case "R":
			for _, coord := range right(x, y, n) {
				if _, exists := points[coord]; exists {
					continue
				}
				points[coord] = step
				x, y = coord.x, coord.y
				step++
			}
		case "L":
			for _, coord := range left(x, y, n) {
				if _, exists := points[coord]; exists {
					continue
				}
				points[coord] = step
				x, y = coord.x, coord.y
				step++
			}
		case "U":
			for _, coord := range up(x, y, n) {
				if _, exists := points[coord]; exists {
					continue
				}
				points[coord] = step
				x, y = coord.x, coord.y
				step++
			}
		case "D":
			for _, coord := range down(x, y, n) {
				if _, exists := points[coord]; exists {
					continue
				}
				points[coord] = step
				x, y = coord.x, coord.y
				step++
			}
		}
	}
	return points
}

func crossed(first, second map[coordinate]int) map[coordinate]int {
	coords := map[coordinate]int{}
	for k, v := range first {
		if contains(k, second) {
			coords[k] = v
		}
	}
	for x := range coords {
		if x.x == 0 && x.y == 0 {
			delete(coords, x)
		}
	}
	return coords
}

func findMinManhattan(candidates map[coordinate]int) int {
	lowest := math.MaxInt64
	for c := range candidates {
		if c.x == 0 && c.y == 0 {
			continue
		}
		sum := int(math.Abs(float64(c.x))) + int(math.Abs(float64(c.y)))
		if sum < lowest {
			lowest = sum
		}
	}
	return lowest
}

func contains(c coordinate, in map[coordinate]int) bool {
	for coord := range in {
		if coord == c {
			return true
		}
	}
	return false
}

func right(x, y, n int) []coordinate {
	points := []coordinate{}
	i := x
	for i <= x+n {
		points = append(points, coordinate{x: i, y: y})
		i++
	}
	return points
}

func left(x, y, n int) []coordinate {
	points := []coordinate{}
	i := x
	for i >= x-n {
		points = append(points, coordinate{x: i, y: y})
		i--
	}
	return points
}

func down(x, y, n int) []coordinate {
	points := []coordinate{}
	i := y
	for i >= y-n {
		points = append(points, coordinate{x: x, y: i})
		i--
	}
	return points
}

func up(x, y, n int) []coordinate {
	points := []coordinate{}
	i := y
	for i <= y+n {
		points = append(points, coordinate{x: x, y: i})
		i++
	}
	return points
}
