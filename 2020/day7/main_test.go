package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseBagGraph(t *testing.T) {
	testCases := []struct {
		input     string
		wantNode  string
		wantEdges []edge
	}{
		{
			"light red bags contain 1 bright white bag, 2 muted yellow bags.",
			"light red",
			[]edge{{weight: 1, value: "bright white"}, {weight: 2, value: "muted yellow"}},
		},
		{
			"dark orange bags contain 3 bright white bags, 4 muted yellow bags.",
			"dark orange",
			[]edge{{weight: 3, value: "bright white"}, {weight: 4, value: "muted yellow"}},
		},
		{
			"bright white bags contain 1 shiny gold bag.",
			"bright white",
			[]edge{{weight: 1, value: "shiny gold"}},
		},
		{
			"faded blue bags contain no other bags.",
			"faded blue",
			[]edge{},
		},
	}
	for _, tC := range testCases {
		t.Run("", func(t *testing.T) {
			gotNode, gotEdges := parseBagRule(tC.input)
			assert.Equal(t, gotNode, tC.wantNode, "Expected [%s] to parse to node %s, got %s", tC.input, tC.wantNode, gotNode)
			assert.Equal(t, gotEdges, tC.wantEdges, "Expected [%s] to parse to edges %v, got %v", tC.input, tC.wantEdges, gotEdges)
		})
	}
}

func TestHasPath(t *testing.T) {
	g := &graph{
		nodes: []string{"light red", "bright white", "muted yellow", "faded blue", "shiny gold"},
		edges: map[string][]edge{
			"light red":    []edge{{value: "muted yellow"}, {value: "bright white"}},
			"bright white": []edge{{value: "shiny gold"}},
			"muted yellow": []edge{{value: "faded blue"}, {value: "shiny gold"}},
			"faded blue":   []edge{},
			"shiny gold":   []edge{},
		},
	}
	testCases := []struct {
		desc string
		from string
		to   string
		want bool
	}{
		{
			"first order connection",
			"bright white",
			"shiny gold",
			true,
		},
		{
			"higher order connection",
			"light red",
			"shiny gold",
			true,
		},
		{
			"no connection",
			"faded blue",
			"shiny gold",
			false,
		},
		{
			"no connection to self",
			"shiny gold",
			"shiny gold",
			false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := hasPath(g, tC.from, tC.to)
			assert.Equal(t, tC.want, got, "Expected path from %s to %s in graph [%v] to %v got %v", tC.from, tC.to, g, tC.want, got)
		})
	}
}

func TestSumOfWeights(t *testing.T) {
	testCases := []struct {
		desc  string
		graph *graph
		want  int
		start string
	}{
		{"", &graph{nodes: []string{"first"}, edges: map[string][]edge{}}, 0, "first"},
		{"", &graph{nodes: []string{"first", "second"}, edges: map[string][]edge{"first": []edge{{weight: 1, value: "second"}}}}, 1, "first"},
		{"", &graph{nodes: []string{"first", "second", "third"}, edges: map[string][]edge{"first": []edge{{weight: 1, value: "second"}, {weight: 2, value: "third"}}}}, 3, "first"},
		{"", &graph{nodes: []string{"first", "second", "third"}, edges: map[string][]edge{"first": []edge{{weight: 1, value: "second"}}, "second": []edge{{weight: 2, value: "third"}}}}, 3, "first"},
		{"", &graph{nodes: []string{"first", "second", "third"}, edges: map[string][]edge{"first": []edge{{weight: 2, value: "second"}}, "second": []edge{{weight: 2, value: "third"}}}}, 4, "first"},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := sumOfWeights(tC.graph, tC.start)
			assert.Equal(t, tC.want, got, "Expected sum of weights for graph [%#v], node %v to be %v, got %v", tC.graph, tC.start, tC.want, got)
		})
	}
}
