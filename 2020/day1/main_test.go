package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain_findSum(t *testing.T) {
	ins := []int{
		1721,
		979,
		366,
		299,
		675,
		1456,
	}
	first, second, err := findSum(2020, ins)
	assert.Nil(t, err)
	assert.Equal(t, first, 1721)
	assert.Equal(t, second, 299)
}

func TestMain_findSumThree(t *testing.T) {
	ins := []int{
		1721,
		979,
		366,
		299,
		675,
		1456,
	}
	first, second, third, err := findSumThree(2020, ins)
	assert.Nil(t, err)
	assert.Equal(t, first, 979)
	assert.Equal(t, second, 366)
	assert.Equal(t, third, 675)
}
