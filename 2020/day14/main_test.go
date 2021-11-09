package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWorking(t *testing.T) {
	assert.Equal(t, true, true)
}

func TestApplyBitmask(t *testing.T) {
	testCases := []struct {
		desc      string
		inputMask string
		inputVal  int
		want      int
	}{
		// {"11 -> 73", "XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X", 11, 73},
		// {"101 -> 101", "XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X", 101, 101},
		// {"0 -> 64", "XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X", 0, 64},
		// {"370803 -> 233", "10X0110X01100X00111XX00001X011101001", 370803, 233},
		// {"955 -> 233", "10X0110X01100X00111XX00001X011101001", 955, 233},
		// {"316949 -> 233", "10X0110X01100X00111XX00001X011101001", 316949, 233},
		{"903991->8873049408", "0X10X001X00011X1111X111111010100X000", 903991, 8873049408},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := applyMask(tC.inputMask, tC.inputVal)
			assert.Equal(t, tC.want, got)
		})
	}
}

func TestRunProgram(t *testing.T) {
	input := `mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X
mem[8] = 11
mem[7] = 101
mem[8] = 0`
	want := 165

	got, err := runProgram(input)

	assert.Nil(t, err)
	assert.Equal(t, want, got)
}

func TestApplyFloatingBitmask(t *testing.T) {
	mask := "000000000000000000000000000000X1001X"
	got := applyFloatingBitmask(mask, 42)
	want := "000000000000000000000000000000X1101X"
	assert.Equal(t, want, got)
}

func TestGenerateAllMasks(t *testing.T) {
	mask := "000000000000000000000000000000X1101X"
	wants := []int64{26, 27, 58, 59}
	got := generateAllMasks(mask)
	assert.Equal(t, wants, got)

	mask = "00000000000000000000000000000001X0XX"
	wants = []int64{16, 17, 18, 19, 24, 25, 26, 27}
	got = generateAllMasks(mask)
	assert.Equal(t, wants, got)
}

func TestExamplePartTwo(t *testing.T) {
	mask := "000000000000000000000000000000X1001X"
	newMask := applyFloatingBitmask(mask, 42)
	gots := generateAllMasks(newMask)
	wants := []int64{26, 27, 58, 59}
	assert.Equal(t, wants, gots)
}

// func TestGenerateAllBitmasks(t *testing.T) {
// 	mask := "000000000000000000000000000000X1101X"
// 	wants := []string{
// 		"000000000000000000000000000000011010",
// 		"000000000000000000000000000000011011",
// 		"000000000000000000000000000000111010",
// 		"000000000000000000000000000000111011",
// 	}
// 	got := generateAllBitmasks(mask)
// 	assert.Equal(t, wants, got)
// }

// func TestApplyFloatingBitmask(t *testing.T) {
// 	mask := "000000000000000000000000000000X1001X"
// 	location := 42
// 	results := applyFloatingBitmask(mask, location)
// 	assert.Equal(t, []int{26, 27, 58, 59}, results)
// }
