package main

import "testing"

func TestCalcFuelWhilePositive(t *testing.T) {
	cases := []struct {
		name  string
		input int64
		want  int64
	}{
		{"14", 14, 2},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			got := calcFuelWhilePositive(tc.input)
			if got != tc.want {
				t.Fatalf("[%s]: Got: %d, Want: %d", tc.name, got, tc.want)
			}
		})
	}
}
