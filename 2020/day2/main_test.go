package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinMaxCharAllowed(t *testing.T) {
	testCases := []struct {
		desc   string
		policy pwEntry
		want   bool
	}{
		{"valid", pwEntry{1, 3, 'a', "abcde"}, true},
		{"invalid: not enough b", pwEntry{1, 3, 'b', "cdefg"}, false},
		{"valid", pwEntry{2, 9, 'c', "ccccccccc"}, true},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := minMaxCharAllowed(tC.policy)
			assert.Equal(t, tC.want, got, fmt.Sprintf("Expected validity of [%s] to be %v, got %v", tC.policy.pw, tC.want, got))
		})
	}
}

func TestPositionXOr(t *testing.T) {
	testCases := []struct {
		desc   string
		policy pwEntry
		want   bool
	}{
		{"valid", pwEntry{min: 1, max: 3, char: 'a', pw: "abcde"}, true},
		{"invalid: no b", pwEntry{min: 1, max: 3, char: 'b', pw: "cdefg"}, false},
		{"invalid: both positions", pwEntry{min: 2, max: 9, char: 'c', pw: "ccccccccc"}, false},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := positionXOr(tC.policy)
			assert.Equal(t, tC.want, got, fmt.Sprintf("Expected [%#v] to be %v, got %v", tC.policy, tC.want, got))
		})
	}
}

func TestParsePWEntry(t *testing.T) {
	input := "1-3 a: abcde"
	want := pwEntry{min: 1, max: 3, char: 'a', pw: "abcde"}
	got, err := parsePWEntry(input)
	assert.Nil(t, err)
	assert.Equal(t, want, got, "Expected [%s] to parse to [%#v], got [%#v]", input, want, got)
}
