package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAcceptance(t *testing.T) {
	input := `
Player 1:
9
2
6
3
1

Player 2:
5
8
4
7
10
	`
	input = strings.TrimSpace(input)
	want := 306

	got := playGame(input)

	assert.Equal(t, want, got)
}

func TestGetScore(t *testing.T) {
	input := []int{3, 2, 10, 6, 8, 5, 9, 4, 7, 1}
	want := 306

	got := getScore(input)

	assert.Equal(t, want, got)
}

func TestParseDeck(t *testing.T) {
	input := `
Player 1:
9
2
6
3
1

Player 2:
5
8
4
7
10
	`
	input = strings.TrimSpace(input)

	wantPlayerOne := []int{9, 2, 6, 3, 1}
	wantPlayerTwo := []int{5, 8, 4, 7, 10}

	gotPlayerOne, gotPlayerTwo := parseDeck(input)

	assert.Equal(t, wantPlayerOne, gotPlayerOne)
	assert.Equal(t, wantPlayerTwo, gotPlayerTwo)
}

func TestPlayGame(t *testing.T) {
	inputPlayerOne := []int{9, 2, 6, 3, 1}
	inputPlayerTwo := []int{5, 8, 4, 7, 10}

	wantPlayerOne := []int{2, 6, 3, 1, 9, 5}
	wantPlayerTwo := []int{8, 4, 7, 10}

	gotPlayerOne, gotPlayerTwo := playTurn(inputPlayerOne, inputPlayerTwo)

	assert.Equal(t, wantPlayerOne, gotPlayerOne)
	assert.Equal(t, wantPlayerTwo, gotPlayerTwo)
}

func TestAcceptanceRecursive(t *testing.T) {
	input := `
Player 1:
9
2
6
3
1

Player 2:
5
8
4
7
10
	`
	input = strings.TrimSpace(input)
	want := 291

	got := playGameRecursive(input)

	assert.Equal(t, want, got)
}
