package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/while1malloc0/advent-of-code/2020/challenge"
)

type group struct {
	answers []string
}

func (g *group) countAny() int {
	seen := map[byte]struct{}{}
	for i := range g.answers {
		for j := range g.answers[i] {
			seen[g.answers[i][j]] = struct{}{}
		}
	}
	return len(seen)
}

func (g *group) countAll() int {
	seen := map[byte]int{}
	for i := range g.answers {
		for j := range g.answers[i] {
			seen[g.answers[i][j]]++
		}
	}
	numAnswers := len(g.answers)
	var count int
	for _, v := range seen {
		if v == numAnswers {
			count++
		}
	}
	return count
}

func parseGroups(groupsraw string) []group {
	var groups []group
	s := bufio.NewScanner(strings.NewReader(groupsraw))
	g := group{answers: []string{}}
	for s.Scan() {
		line := s.Text()
		if line == "" {
			groups = append(groups, g)
			g = group{answers: []string{}}
			continue
		}
		g.answers = append(g.answers, line)
	}
	groups = append(groups, g)
	return groups
}

func main() {
	in, err := ioutil.ReadFile("input")
	if err != nil {
		panic(err)
	}
	groups := parseGroups(string(in))
	partOneFunc := func() error {
		var sum int
		for i := range groups {
			sum += groups[i].countAny()
		}
		fmt.Println(sum)
		return nil
	}

	partTwoFunc := func() error {
		var sum int
		for i := range groups {
			sum += groups[i].countAll()
		}
		fmt.Println(sum)
		return nil
	}

	challenge.Run(partOneFunc, partTwoFunc)
}
