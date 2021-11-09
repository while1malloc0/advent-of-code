package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindFirstNonSum(t *testing.T) {
	input := []int64{
		35,
		20,
		15,
		25,
		47,
		40,
		62,
		55,
		65,
		95,
		102,
		117,
		150,
		182,
		127,
		219,
		299,
		277,
		309,
		576,
	}
	got, err := findFirstNonSum(5, input)
	want := int64(127)
	assert.Nil(t, err)
	assert.Equal(t, want, got)
}

func TestFindContiguousRange(t *testing.T) {
	input := []int64{
		35,
		20,
		15,
		25,
		47,
		40,
		62,
		55,
		65,
		95,
		102,
		117,
		150,
		182,
		127,
		219,
		299,
		277,
		309,
		576,
	}
	needle := int64(127)
	want := []int64{15, 25, 47, 40}
	got, err := findContiguousRange(needle, input)
	assert.Nil(t, err)
	assert.Equal(t, want, got)
}
