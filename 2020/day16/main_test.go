package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumberRangeContains(t *testing.T) {
	testCases := []struct {
		numberRange *numberRange
		input       int
		want        bool
	}{
		{&numberRange{start: 1, end: 3}, 7, false},
		{&numberRange{start: 5, end: 7}, 7, true},
		{&numberRange{start: 1, end: 3}, 3, true},
	}
	for _, tC := range testCases {
		t.Run("", func(t *testing.T) {
			got := tC.numberRange.contains(tC.input)
			assert.Equal(t, tC.want, got)
		})
	}
}

func TestNumberRangeGroupContains(t *testing.T) {
	testCases := []struct {
		inputGroup *numberRangeGroup
		input      int
		want       bool
	}{
		{
			&numberRangeGroup{groups: []*numberRange{{start: 1, end: 3}, {start: 5, end: 7}}},
			7,
			true,
		},
		{
			&numberRangeGroup{groups: []*numberRange{{start: 1, end: 3}, {start: 5, end: 7}}},
			4,
			false,
		},
	}
	for _, tC := range testCases {
		t.Run("", func(t *testing.T) {
			got := tC.inputGroup.contains(tC.input)
			assert.Equal(t, tC.want, got)
		})
	}
}

// Make sure that memoizing returns the same answer
func TestNumberRangeGroupContains_checkMemo(t *testing.T) {
	ng := &numberRangeGroup{groups: []*numberRange{{start: 1, end: 3}, {start: 5, end: 7}}}
	guess := 7
	want := true

	// first time, no memo
	got := ng.contains(guess)
	assert.Equal(t, want, got)

	// second time, avec memo
	got = ng.contains(guess)
	assert.Equal(t, want, got)
}

func TestTicketValid(t *testing.T) {
	testCases := []struct {
		input *ticketValidation
		want  int
	}{
		{
			&ticketValidation{
				ticket: ticket{7, 3, 47},
				ranges: &numberRangeGroup{groups: []*numberRange{{1, 3}, {5, 7}, {6, 11}, {33, 44}, {13, 40}, {45, 50}}},
			},
			NoInvalid,
		},
		{
			&ticketValidation{
				ticket: ticket{40, 4, 50},
				ranges: &numberRangeGroup{groups: []*numberRange{{1, 3}, {5, 7}, {6, 11}, {33, 44}, {13, 40}, {45, 50}}},
			},
			4,
		},
		{
			&ticketValidation{
				ticket: ticket{55, 2, 20},
				ranges: &numberRangeGroup{groups: []*numberRange{{1, 3}, {5, 7}, {6, 11}, {33, 44}, {13, 40}, {45, 50}}},
			},
			55,
		},
		{
			&ticketValidation{
				ticket: ticket{38, 6, 12},
				ranges: &numberRangeGroup{groups: []*numberRange{{1, 3}, {5, 7}, {6, 11}, {33, 44}, {13, 40}, {45, 50}}},
			},
			12,
		},
	}
	for _, tC := range testCases {
		t.Run("", func(t *testing.T) {
			got := tC.input.findInvalid()
			assert.Equal(t, tC.want, got)
		})
	}
}

func TestParseNumberRangeGroup(t *testing.T) {
	input := `class: 1-3 or 5-7
row: 6-11 or 33-44
seat: 13-40 or 45-50

your ticket:
7,1,14

nearby tickets:
7,3,47
40,4,50
55,2,20
38,6,12`

	want := []*numberRangeGroup{
		{name: "class", groups: []*numberRange{{1, 3}, {5, 7}}},
		{name: "row", groups: []*numberRange{{6, 11}, {33, 44}}},
		{name: "seat", groups: []*numberRange{{13, 40}, {45, 50}}},
	}

	got := parseNumberRangeGroups(input)

	assert.Equal(t, want, got)
}

func TestParseTickets(t *testing.T) {
	input := `class: 1-3 or 5-7
row: 6-11 or 33-44
seat: 13-40 or 45-50

your ticket:
7,1,14

nearby tickets:
7,3,47
40,4,50
55,2,20
38,6,12`
	want := []ticket{{7, 3, 47}, {40, 4, 50}, {55, 2, 20}, {38, 6, 12}}

	got := parseTickets(input)

	assert.Equal(t, want, got)
}

func TestFilterInvalid(t *testing.T) {
	input := []ticket{{7, 3, 47}, {40, 4, 50}, {55, 2, 20}, {38, 6, 12}}
	ng := &numberRangeGroup{
		groups: []*numberRange{{1, 3}, {5, 7}, {6, 11}, {33, 44}, {13, 40}, {45, 50}},
	}

	got := filterInvalidTickets(input, ng)

	want := []ticket{{7, 3, 47}}

	assert.Equal(t, want, got)
}

func TestGetColumns(t *testing.T) {
	input := []ticket{{3, 9, 18}, {15, 1, 5}, {5, 14, 9}}
	want := [][]int{{3, 15, 5}, {9, 1, 14}, {18, 5, 9}}

	got := getColumns(input)

	assert.Equal(t, want, got)
}

func TestGetFieldPositions(t *testing.T) {
	input := []ticket{{3, 9, 18}, {15, 1, 5}, {5, 14, 9}}
	ngs := []*numberRangeGroup{
		&numberRangeGroup{
			name:   "class",
			groups: []*numberRange{{0, 1}, {4, 19}},
		},
		&numberRangeGroup{
			name:   "row",
			groups: []*numberRange{{0, 5}, {8, 19}},
		},
		&numberRangeGroup{
			name:   "seat",
			groups: []*numberRange{{0, 13}, {16, 19}},
		},
	}
	want := []string{"row", "class", "seat"}

	got := getFieldPositions(input, ngs)

	assert.Equal(t, want, got)
}
