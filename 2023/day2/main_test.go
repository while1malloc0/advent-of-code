package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartOne(t *testing.T) {
	input, err := os.ReadFile("example.txt")
	if err != nil {
		t.Fatal(err)
	}

	got := PartOne(input)

	want := 8
	if got != want {
		t.Fatalf("failed, expected %d, got %d", want, got)
	}
}

func TestParse(t *testing.T) {
	want := Game{
		ID: 1,
		Matches: []map[string]int{
			{"blue": 3, "red": 4},
			{"red": 1, "green": 2, "blue": 6},
			{"green": 2},
		},
	}
	input := []byte(`Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green`)
	got, err := Parse(input)
	assert.Nil(t, err)
	assert.Equal(t, want, got)
}

func TestPartTwo(t *testing.T) {
	input, err := os.ReadFile("example.txt")
	if err != nil {
		t.Fatal(err)
	}

	got := PartTwo(input)

	want := 2286
	if got != want {
		t.Fatalf("failed, expected %d, got %d", want, got)
	}
}
