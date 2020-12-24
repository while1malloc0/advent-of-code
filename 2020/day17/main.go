package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/while1malloc0/advent-of-code/2020/challenge"
)

type coord struct {
	x, y, z, w int
}

type cube = map[coord]struct{}

func parseCube(in string) cube {
	out := cube{}

	for y, line := range strings.Split(in, "\n") {
		for x, char := range line {
			if char == '#' {
				c := coord{x: x, y: y, z: 0}
				out[c] = struct{}{}
			}
		}
	}

	return out
}

func nextState(location coord, c cube) bool {
	// Active
	if _, ok := c[location]; ok {
		var numActive int
		for dw := -1; dw <= 1; dw++ {
			for dz := -1; dz <= 1; dz++ {
				for dy := -1; dy <= 1; dy++ {
					for dx := -1; dx <= 1; dx++ {
						// On the current coordinate
						if dx == 0 && dy == 0 && dz == 0 && dw == 0 {
							continue
						}
						test := coord{x: location.x + dx, y: location.y + dy, z: location.z + dz, w: location.w + dw}
						if _, found := c[test]; found {
							numActive++
						}
					}
				}
			}
		}
		return numActive == 2 || numActive == 3
	}

	// Inactive
	var numActive int
	for dw := -1; dw <= 1; dw++ {
		for dz := -1; dz <= 1; dz++ {
			for dy := -1; dy <= 1; dy++ {
				for dx := -1; dx <= 1; dx++ {
					// On the current coordinate
					if dx == 0 && dy == 0 && dz == 0 && dw == 0 {
						continue
					}
					test := coord{x: location.x + dx, y: location.y + dy, z: location.z + dz, w: location.w + dw}
					if _, found := c[test]; found {
						numActive++
					}
				}
			}
		}
	}
	return numActive == 3
}

func nextCubeState(c cube) cube {
	dc := cube{}
	for k, v := range dc {
		dc[k] = v
	}

	for w := -20; w <= 20; w++ {
		for z := -20; z <= 20; z++ {
			for y := -20; y <= 20; y++ {
				for x := -20; x <= 20; x++ {
					location := coord{x: x, y: y, z: z, w: w}
					active := nextState(location, c)
					if active {
						dc[location] = struct{}{}
					} else {
						delete(dc, location)
					}
				}
			}

		}
	}

	return dc
}

func runBoot(c cube, cycles int) int {
	dc := cube{}
	for k, v := range c {
		dc[k] = v
	}

	for i := 0; i < cycles; i++ {
		dc = nextCubeState(dc)
	}

	return len(dc)
}

func main() {
	// Part two is the same as part one, just with more dimensions. This logic
	// works for both.
	partOneFunc := func() error {
		in, err := ioutil.ReadFile("input")
		if err != nil {
			return err
		}
		c := parseCube(string(in))

		result := runBoot(c, 6)
		fmt.Println(result)

		return nil
	}

	challenge.Run(partOneFunc, nil)
}
