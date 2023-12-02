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

	want := 142
	if got != want {
		t.Fatalf("failed, expected %d, got %d", want, got)
	}
}

func TestExtractDigits(t *testing.T) {
	testCases := []struct {
		desc  string
		input string
		want  int
	}{
		{"first and last digit", "1abc2", 12},
		{"padded digits", "pqr3stu8vwx", 38},
		{"multiple digits", "a1b2c3d4e5f", 15},
		{"only one digit", "treb7uchet", 77},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := ExtractDigits(tC.input)
			if tC.want != got {
				t.Fatalf("ExtractDigits failed, expected %d, got %d", tC.want, got)
			}
		})
	}
}

func TestPartTwo(t *testing.T) {
	input, err := os.ReadFile("example-2.txt")
	if err != nil {
		t.Fatal(err)
	}

	got := PartTwo(input)

	want := 281
	if got != want {
		t.Fatalf("failed, expected %d, got %d", want, got)
	}
}

func TestExtractNamedDigits(t *testing.T) {
	testCases := []struct {
		desc  string
		input string
		want  int
	}{
		{"", "two1nine", 29},
		{"", "eightwothree", 83},
		{"", "abcone2threexyz", 13},
		{"", "xtwone3four", 24},
		{"", "4nineeightseven2", 42},
		{"", "zoneight234", 14},
		{"", "7pqrstsixteen", 76},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := ExtractNamedDigits(tC.input)
			if tC.want != got {
				t.Fatalf("ExtractNamedDigits failed, expected %d, got %d", tC.want, got)
			}
		})
	}
}
