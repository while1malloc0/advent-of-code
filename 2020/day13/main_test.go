package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetNearestMultipleOf(t *testing.T) {
	testCases := []struct {
		target int
		input  int
		want   int
	}{
		{939, 7, 945},
		{939, 13, 949},
		{939, 59, 944},
	}
	for _, tC := range testCases {
		t.Run("", func(t *testing.T) {
			got := getNearestMultipleOf(tC.input, tC.target)
			assert.Equal(t, tC.want, got)
		})
	}
}

func TestExample(t *testing.T) {
	target := 939
	nums := []int{
		7, 13, 59, 31, 19,
	}
	id, departure := getNearestBusID(target, nums)
	assert.Equal(t, 59, id)
	assert.Equal(t, 944, departure)
}

func TestGetMagicTimestamp(t *testing.T) {
	nums := []int{
		7, 13, SKIP, SKIP, 59, SKIP, 31, 19,
	}
	result := getMagicTimestamp(nums)
	assert.Equal(t, 1068781, result)
}
