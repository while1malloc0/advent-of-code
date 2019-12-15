package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

var part2 = flag.Bool("part2", false, "run part two")

func main() {
	flag.Parse()

	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}

	s := bufio.NewScanner(f)
	seen := map[string]string{}
	for s.Scan() {
		tokens := strings.Split(strings.TrimSpace(s.Text()), ")")
		left, right := tokens[0], tokens[1]
		seen[right] = left
	}
	if *part2 {
		start := seen["YOU"]
		dest := seen["SAN"]
		commonAncestor := getCommonAncestor(start, dest, seen)
		fmt.Println(distanceFrom(start, commonAncestor, seen) + distanceFrom(dest, commonAncestor, seen))
	} else {
		var count int
		for k := range seen {
			kk := k
			for {
				next, ok := seen[kk]
				if !ok {
					break
				}
				count++
				kk = next
			}
		}
		fmt.Println(count)
	}
}

func getCommonAncestor(a, b string, m map[string]string) string {
	candidate := a
	for {
		if isAncestor(candidate, b, m) {
			return candidate
		}
		candidate = m[candidate]
	}
}

func isAncestor(a, b string, m map[string]string) bool {
	current := b
	for {
		if current == a {
			return true
		}
		next, ok := m[current]
		if !ok {
			break
		}
		current = next
	}
	return false
}

func distanceFrom(a, b string, m map[string]string) int {
	var distance int
	current := a
	for {
		if current == b {
			break
		}
		next, ok := m[current]
		if !ok {
			break
		}
		distance++
		current = next
	}
	return distance
}
