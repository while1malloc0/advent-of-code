package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueueEval(t *testing.T) {
	input := queue{"1", "+", "2"}

	got := input.eval()
	want := 3

	assert.Equal(t, want, got)
}

func TestQueueEval_orderOfOperations(t *testing.T) {
	// in this problem, all operators have equal precendence
	input := queue{"1", "+", "2", "*", "3"}

	got := input.eval()
	want := 9

	assert.Equal(t, want, got)
}

func TestPuzzle(t *testing.T) {
	testCases := []struct {
		input string
		want  int
	}{
		{"1 + 2 * 3 + 4 * 5 + 6", 71},
		{"1 + (2 * 3) + (4 * (5 + 6))", 51},
		{"5 + (8 * 3 + 9 + 3 * 4 * 3)", 437},
		{"5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))", 12240},
		{"((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2", 13632},
	}
	for _, tC := range testCases {
		t.Run("", func(t *testing.T) {
			got := eval(tC.input, evalStrategyNoPrecedence)
			assert.Equal(t, tC.want, got)
		})
	}
}

func TestEvalPrecendence(t *testing.T) {
	testCases := []struct {
		input queue
		want  int
	}{
		{queue{"1", "+", "2", "*", "3"}, 9},
		{queue{"1", "+", "3", "*", "2"}, 8},
	}
	for _, tC := range testCases {
		t.Run("", func(t *testing.T) {
			got := tC.input.evalWithPrecedence()
			assert.Equal(t, tC.want, got)
		})
	}
}

func TestPuzzleWithPrecedence(t *testing.T) {
	testCases := []struct {
		input string
		want  int
	}{
		{"1 + (2 * 3) + (4 * (5 + 6))", 51},
		{"2 * 3 + (4 * 5)", 46},
		{"5 + (8 * 3 + 9 + 3 * 4 * 3)", 1445},
		{"5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))", 669060},
		{"((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2", 23340},
	}
	for _, tC := range testCases {
		t.Run("", func(t *testing.T) {
			got := eval(tC.input, evalStrategyWithPrecedence)
			assert.Equal(t, tC.want, got)
		})
	}
}
