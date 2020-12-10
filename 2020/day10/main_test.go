package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample(t *testing.T) {
	input := []int{
		16,
		10,
		15,
		5,
		1,
		11,
		7,
		19,
		6,
		12,
		4,
	}

	gotOnes, gotThrees := countOnesAndThrees(input)

	wantOnes, wantThrees := 7, 5
	assert.Equal(t, wantOnes, gotOnes)
	assert.Equal(t, wantThrees, gotThrees)
}

func TestExampleTwo(t *testing.T) {
	input := []int{
		28,
		33,
		18,
		42,
		31,
		14,
		46,
		20,
		48,
		47,
		24,
		23,
		49,
		45,
		19,
		38,
		39,
		11,
		1,
		32,
		25,
		35,
		8,
		17,
		7,
		9,
		4,
		2,
		34,
		10,
		3,
	}

	gotOnes, gotThrees := countOnesAndThrees(input)

	wantOnes, wantThrees := 22, 10
	assert.Equal(t, wantOnes, gotOnes)
	assert.Equal(t, wantThrees, gotThrees)
}
