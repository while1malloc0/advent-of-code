package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
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
