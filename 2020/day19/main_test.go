package main

import (
	"regexp"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseRegexMap(t *testing.T) {
	input := `
0: 1 2
1: a
2: 1 3 | 3 1
3: b
	`
	input = strings.TrimSpace(input)

	got := parseRegexMap(input)

	want := map[string]string{
		"0": "1 2",
		"1": "a",
		"2": "1 3 | 3 1",
		"3": "b",
	}

	assert.Equal(t, want, got)
}

func TestParseExpression(t *testing.T) {
	input := `
0: 1 2
1: a
2: 1 3 | 3 1
3: b
	`
	input = strings.TrimSpace(input)
	got := parseRegex(input)
	want := "^a(ab|ba)$"
	assert.Equal(t, want, got)
}

func TestParseExpression_complex(t *testing.T) {
	input := `
0: 4 1 5
1: 2 3 | 3 2
2: 4 4 | 5 5
3: 4 5 | 5 4
4: "a"
5: "b"
	`
	input = strings.TrimSpace(input)
	got := parseRegex(input)
	want := "^a((aa|bb)(ab|ba)|(ab|ba)(aa|bb))b$"
	_, err := regexp.Compile(got)

	assert.Equal(t, want, got)
	assert.Nil(t, err)
}

func TestMatches(t *testing.T) {
	testCases := []struct {
		input string
		want  bool
	}{
		{"ababbb", true},
		{"abbbab", true},
		{"bababa", false},
		{"aaabbb", false},
		{"aaaabbb", false},
	}
	reg := `
0: 4 1 5
1: 2 3 | 3 2
2: 4 4 | 5 5
3: 4 5 | 5 4
4: "a"
5: "b"
	`
	reg = strings.TrimSpace(reg)
	for _, tC := range testCases {
		t.Run("", func(t *testing.T) {
			got := testMatch(reg, tC.input)
			assert.Equal(t, tC.want, got)
		})
	}
}
