package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSeatID(t *testing.T) {
	testCases := []struct {
		row  int
		col  int
		want int
	}{
		{44, 5, 357},
		{70, 7, 567},
		{14, 7, 119},
		{102, 4, 820},
	}
	for _, tC := range testCases {
		t.Run("", func(t *testing.T) {
			got := getSeatID(tC.row, tC.col)
			assert.Equal(t, tC.want, got, "Expected %v, but got %v", tC.want, got)
		})
	}
}

func TestGetRow(t *testing.T) {
	testCases := []struct {
		input string
		want  int
	}{
		{"FBFBBFF", 44},
		{"BFFFBBF", 70},
		{"FFFBBBF", 14},
		{"BBFFBBF", 102},
	}
	for _, tC := range testCases {
		t.Run("", func(t *testing.T) {
			got := getRow(tC.input)
			assert.Equal(t, tC.want, got, "Expected %v, got %v", tC.want, got)
		})
	}
}

func TestGetColumn(t *testing.T) {
	testCases := []struct {
		input string
		want  int
	}{
		{"RLR", 5},
		{"RRR", 7},
		{"RLL", 4},
	}
	for _, tC := range testCases {
		t.Run("", func(t *testing.T) {
			got := getColumn(tC.input)
			assert.Equal(t, tC.want, got, "Expected %v got %v", tC.want, got)
		})
	}
}

func TestSeatIDFromBoardingPass(t *testing.T) {
	testCases := []struct {
		input string
		want  int
	}{
		{"FBFBBFFRLR", 357},
		{"BFFFBBFRRR", 567},
		{"FFFBBBFRRR", 119},
		{"BBFFBBFRLL", 820},
	}
	for _, tC := range testCases {
		t.Run("", func(t *testing.T) {
			got := seatIDFromBoardingPass(tC.input)
			assert.Equal(t, tC.want, got, "Expected %v got %v", tC.want, got)
		})
	}
}
