package main

import "testing"

func TestDayTwoPartOne(t *testing.T) {
	cases := []struct {
		name  string
		input string
		want  string
	}{
		{"sample 1", "1,0,0,0,99", "2,0,0,0,99"},
		{"sample 2", "2,3,0,3,99", "2,3,0,6,99"},
		{"sample 3", "2,4,4,5,99,0", "2,4,4,5,99,9801"},
		{"sample 4", "1,1,1,4,99,5,6,0,99", "30,1,1,4,2,5,6,0,99"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := executeString(tc.input)
			if got != tc.want {
				t.Fatalf("[%s]: wanted %s got %s", tc.name, tc.want, got)
			}
		})
	}
}
