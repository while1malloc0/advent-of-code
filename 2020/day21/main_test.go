package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseIngredientList(t *testing.T) {
	input := `
mxmxvkd kfcds sqjhc nhms (contains dairy, fish)
trh fvjkl sbzzf mxmxvkd (contains dairy)
sqjhc fvjkl (contains soy)
sqjhc mxmxvkd sbzzf (contains fish)
	`
	input = strings.TrimSpace(input)

	want := map[string][][]string{
		"dairy": {{"mxmxvkd", "kfcds", "sqjhc", "nhms"}, {"trh", "fvjkl", "sbzzf", "mxmxvkd"}},
		"fish":  {{"mxmxvkd", "kfcds", "sqjhc", "nhms"}, {"sqjhc", "mxmxvkd", "sbzzf"}},
		"soy":   {{"sqjhc", "fvjkl"}},
	}

	got := parseIngredientList(input)

	assert.Equal(t, want, got)
}

func TestIngredientListToCandidateList(t *testing.T) {
	input := map[string][][]string{
		"dairy": {{"mxmxvdk", "kfcds", "sqjhc", "nhms"}, {"trh", "fvjkl", "sbzzf", "mxmxvdk"}},
		"fish":  {{"mxmxvdk", "kfcds", "sqjhc", "nhms"}, {"sqjhc", "mxmxvdk", "sbzzf"}},
		"soy":   {{"sqjhc", "fujkl"}},
	}

	want := map[string][]string{
		"dairy": {"mxmxvdk"},
		"fish":  {"mxmxvdk", "sqjhc"},
		"soy":   {"sqjhc", "fujkl"},
	}

	got := ingredientListToCandidates(input)

	assert.Equal(t, want, got)
}
