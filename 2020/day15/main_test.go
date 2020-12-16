package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRunGame(t *testing.T) {
	testCases := []struct {
		input []int
		want  int
		turns int
	}{
		{[]int{0, 3, 6}, 436, 2020},
		{[]int{2, 1, 3}, 10, 2020},
		{[]int{1, 2, 3}, 27, 2020},
		{[]int{2, 3, 1}, 78, 2020},
		{[]int{0, 3, 6}, 175594, 30000000},
	}
	for _, tC := range testCases {
		t.Run("Test", func(t *testing.T) {
			got := runGame(tC.input, tC.turns)
			assert.Equal(t, tC.want, got)
		})
	}
}
