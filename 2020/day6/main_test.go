package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGroupCountAny(t *testing.T) {
	testCases := []struct {
		input group
		want  int
	}{
		{group{answers: []string{"abc"}}, 3},
		{group{answers: []string{"a", "b", "c"}}, 3},
		{group{answers: []string{"ab", "bc"}}, 3},
		{group{answers: []string{"a", "a", "a", "a"}}, 1},
	}
	for _, tC := range testCases {
		t.Run("", func(t *testing.T) {
			got := tC.input.countAny()
			assert.Equal(t, tC.want, got)
		})
	}
}

func TestGroupCountAll(t *testing.T) {
	testCases := []struct {
		input group
		want  int
	}{
		{group{answers: []string{"abc"}}, 3},
		{group{answers: []string{"a", "b", "c"}}, 0},
		{group{answers: []string{"ab", "ac"}}, 1},
		{group{answers: []string{"a", "a", "a", "a"}}, 1},
	}
	for _, tC := range testCases {
		t.Run("", func(t *testing.T) {
			got := tC.input.countAll()
			assert.Equal(t, tC.want, got)
		})
	}
}

func TestParseGroups(t *testing.T) {
	input := `abc

a
b
c

ab
ac

a
a
a
a

b
`
	want := []group{
		{answers: []string{"abc"}},
		{answers: []string{"a", "b", "c"}},
		{answers: []string{"ab", "ac"}},
		{answers: []string{"a", "a", "a", "a"}},
		{answers: []string{"b"}},
	}
	got := parseGroups(input)
	assert.Equal(t, got, want)
}
