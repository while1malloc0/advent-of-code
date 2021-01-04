package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample(t *testing.T) {
	raw := 389125467
	given := parseLinkedList(raw)

	want := 67384529

	result := playGame(given, 100)
	got := getAllFromOne(result)

	assert.Equal(t, want, got)
}

func TestParseLinkedList(t *testing.T) {
	given := 389125467

	want := "3->8->9->1->2->5->4->6->7->"

	list := parseLinkedList(given)

	got := list.String()

	assert.Equal(t, want, got)
}

func TestGetAllFromOne(t *testing.T) {
	raw := 583741926
	given := parseLinkedList(raw)
	want := 92658374

	got := getAllFromOne(given)

	assert.Equal(t, want, got)
}
