package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var (
	part = flag.Int("part", 1, "the part to run")
	file = flag.String("file", "example.txt", "the file to read")

	idRegex   = regexp.MustCompile(`Game (?P<id>\d+)`)
	playRegex = regexp.MustCompile(`(?P<num>\d+) (?P<color>red|green|blue)`)
)

type Game struct {
	ID      int
	Matches []map[string]int
}

type Match struct {
	Blue  int
	Red   int
	Green int
}

func Parse(in []byte) (Game, error) {
	splitColon := bytes.Split(in, []byte(":"))
	rawGame, rawMatches := splitColon[0], splitColon[1]

	rawId := idRegex.FindSubmatch(rawGame)[1]
	id, err := strconv.Atoi(string(rawId))
	if err != nil {
		return Game{}, err
	}

	splitRawMatches := bytes.Split(rawMatches, []byte(`;`))
	var matches []map[string]int
	for i := range splitRawMatches {
		rawMatch := splitRawMatches[i]
		m := map[string]int{}
		rawPlays := bytes.Split(rawMatch, []byte(`,`))
		for j := range rawPlays {
			rawPlay := rawPlays[j]
			matched := playRegex.FindSubmatch(rawPlay)
			if matched == nil {
				panic(fmt.Sprintf("bad play: %s", string(rawPlay)))
			}
			color := string(matched[2])
			value, _ := strconv.Atoi(string(matched[1]))
			m[color] = value
		}
		matches = append(matches, m)
	}
	return Game{ID: id, Matches: matches}, nil
}

func PartOne(input []byte) int {
	var gs []Game
	for _, line := range bytes.Split(input, []byte("\n")) {
		g, err := Parse(line)
		if err != nil {
			panic(err)
		}
		gs = append(gs, g)
	}
	maxes := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
	var ids []int
	for _, g := range gs {
		var impossible bool
	MATCHES:
		for _, m := range g.Matches {
			for k, v := range m {
				if max, ok := maxes[k]; ok && v > max {
					impossible = true
					break MATCHES
				}
			}
		}
		if !impossible {
			ids = append(ids, g.ID)
		}
	}

	var sum int
	for _, id := range ids {
		sum += id
	}

	return sum
}

func PartTwo(input []byte) int {
	var gs []Game
	for _, line := range bytes.Split(input, []byte("\n")) {
		g, err := Parse(line)
		if err != nil {
			panic(err)
		}
		gs = append(gs, g)
	}

	var powers []int
	for _, g := range gs {
		mins := map[string]int{}
		for _, m := range g.Matches {
			for k, v := range m {
				min, ok := mins[k]
				if !ok {
					mins[k] = v
				} else if v > min {
					mins[k] = v
				}
			}
		}
		power := mins["red"] * mins["green"] * mins["blue"]
		powers = append(powers, power)
	}

	var sum int
	for i := range powers {
		sum += powers[i]
	}

	return sum
}

func main() {
	flag.Parse()

	in, err := os.ReadFile(*file)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	switch *part {
	case 1:
		fmt.Println(PartOne(in))
	case 2:
		fmt.Println(PartTwo(in))
	}
}
