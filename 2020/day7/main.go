package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/while1malloc0/advent-of-code/2020/challenge"
)

type edge struct {
	weight int
	value  string
}

type graph struct {
	nodes []string
	edges map[string][]edge
}

type queue struct {
	items []string
}

func (q *queue) enqueue(s string) {
	q.items = append(q.items, s)
}

func (q *queue) dequeue() string {
	item := q.items[0]
	if len(q.items) == 1 {
		q.items = []string{}
	} else {
		q.items = q.items[1:]
	}
	return item
}

func (q *queue) empty() bool {
	return len(q.items) == 0
}

var bagRuleRegex = regexp.MustCompile("(.*) bags contain (.*)")
var numBagsRegex = regexp.MustCompile("(\\d) (.*) bags?")

func parseBagRule(s string) (string, []edge) {
	parsed := bagRuleRegex.FindStringSubmatch(s)
	node := parsed[1]

	if strings.Contains(parsed[2], "no other bags.") {
		return node, []edge{}
	}

	rawEdges := strings.Split(parsed[2], ",")
	var edges []edge
	for _, rawEdge := range rawEdges {
		parsedEdgeStr := numBagsRegex.FindStringSubmatch(rawEdge)
		edgeValue := parsedEdgeStr[2]
		edgeWeight, err := strconv.Atoi(parsedEdgeStr[1])
		if err != nil {
			panic(err)
		}
		edge := edge{value: edgeValue, weight: edgeWeight}
		edges = append(edges, edge)
	}
	return node, edges
}

func hasPath(g *graph, from, to string) bool {
	// We don't consider a thing as having a connection with itself. Jung is //
	// impressed.
	if from == to {
		return false
	}
	visited := map[string]bool{}
	q := &queue{items: []string{}}
	q.enqueue(from)
	for !q.empty() {
		node := q.dequeue()
		if node == to {
			return true
		}
		visited[node] = true
		for _, edge := range g.edges[node] {
			if !visited[edge.value] {
				q.enqueue(edge.value)
			}
		}
	}
	return false
}

func doSumOfWeights(g *graph, node string, sum *int, mult int) {
	for _, edge := range g.edges[node] {
		*sum += edge.weight * mult
		doSumOfWeights(g, edge.value, sum, edge.weight*mult)
	}
}

func sumOfWeights(g *graph, start string) int {
	var sum int
	mult := 1
	doSumOfWeights(g, start, &sum, mult)
	return sum
}

func main() {
	g := &graph{nodes: []string{}, edges: map[string][]edge{}}

	challenge.InputScanFunc("input", func(s string) error {
		node, edges := parseBagRule(s)
		g.nodes = append(g.nodes, node)
		g.edges[node] = edges
		return nil
	})

	partOneFunc := func() error {
		var count int
		for _, node := range g.nodes {
			if hasPath(g, node, "shiny gold") {
				count++
			}
		}
		fmt.Println(count)
		return nil
	}

	partTwoFunc := func() error {
		sum := sumOfWeights(g, "shiny gold")
		fmt.Println(sum)
		return nil
	}

	challenge.Run(partOneFunc, partTwoFunc)
}
