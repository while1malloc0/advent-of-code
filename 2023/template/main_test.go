package main

import (
	"os"
	"testing"
)

func TestPartOne(t *testing.T) {
	input, err := os.ReadFile("example.txt")
	if err != nil {
		t.Fatal(err)
	}

	got := PartOne(input)

	want := -1
	if got != want {
		t.Fatalf("failed, expected %d, got %d", want, got)
	}
}

func TestPartTwo(t *testing.T) {
	input, err := os.ReadFile("example.txt")
	if err != nil {
		t.Fatal(err)
	}

	got := PartTwo(input)

	want := -1
	if got != want {
		t.Fatalf("failed, expected %d, got %d", want, got)
	}
}
