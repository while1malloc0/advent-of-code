package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var part2 = flag.Bool("part2", false, "run part two")

const width, height = 25, 6 // lol

func main() {
	flag.Parse()
	content, err := ioutil.ReadFile("input")
	if err != nil {
		panic(err)
	}
	input := string(content)
	input = strings.TrimSpace(input)
	layers := [][height][width]int{}
	s := bufio.NewScanner(strings.NewReader(input))
	s.Split(bufio.ScanRunes)
OUTER:
	for {
		layer := [height][width]int{}
		for i := 0; i < height; i++ {
			for j := 0; j < width; j++ {
				if !s.Scan() {
					break OUTER
				}
				num, err := strconv.Atoi(s.Text())
				if err != nil {
					panic(err)
				}
				layer[i][j] = num
			}
		}
		layers = append(layers, layer)
	}
	if !*part2 {
		counts := map[int]int{}
		for i := range layers {
			counts[i] = getNumZeros(layers[i])
		}
		var lowestNumZeroes int = 99999999
		lowestNumZeroIdx := 0
		for idx, numZeros := range counts {
			if numZeros < lowestNumZeroes {
				lowestNumZeroIdx = idx
				lowestNumZeroes = numZeros
			}
		}
		theLayer := layers[lowestNumZeroIdx]
		answer := getMultOnesAndTwos(theLayer)
		fmt.Println(answer)
		return
	}
	image := renderLayers(layers)
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			pixel := image[i][j]
			switch pixel {
			case "B":
				fmt.Printf("\u001b[32m*\u001b[0m")
			case "W":
				fmt.Printf("\u001b[37m*\u001b[0m")
			default:
				fmt.Printf("")
			}
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}

func renderLayers(layers [][height][width]int) [height][width]string {
	out := [height][width]string{}
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			for z := 0; z < len(layers); z++ {
				color := getColor(layers[z][i][j])
				if color == "B" || color == "W" {
					out[i][j] = color
					break
				}
			}
		}
	}
	return out
}

func getColor(i int) string {
	switch i {
	case 0:
		return "B"
	case 1:
		return "W"
	default: // lol input validation
		return "T"
	}
}

func getNumZeros(in [height][width]int) int {
	var numZeros int
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if in[i][j] == 0 {
				numZeros++
			}
		}
	}
	return numZeros
}

func getMultOnesAndTwos(in [height][width]int) int {
	var numOnes, numTwos int
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if in[i][j] == 1 {
				numOnes++
			} else if in[i][j] == 2 {
				numTwos++
			}
		}
	}
	return numOnes * numTwos
}
