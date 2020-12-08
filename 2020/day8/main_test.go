package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// nop +0
// acc +1
// jmp +4
// acc +3
// jmp -3
// acc -99
// acc +1
// jmp -4
// acc +6

func TestParseProgram(t *testing.T) {
	testCases := []struct {
		desc  string
		input string
		want  *program
	}{
		{"nop instruction", "nop +0", &program{instructions: []instruction{{op: opNop, arg: "+0"}}}},
		{"acc instruction", "acc +1", &program{instructions: []instruction{{op: opAcc, arg: "+1"}}}},
		{"jmp instruction", "jmp +4", &program{instructions: []instruction{{op: opJmp, arg: "+4"}}}},
		{"multiple instructions", "nop +0\nacc +1\njmp +4", &program{instructions: []instruction{{op: opNop, arg: "+0"}, {op: opAcc, arg: "+1"}, {op: opJmp, arg: "+4"}}}},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := parseProgram(tC.input)
			assert.Equal(t, tC.want, got, "Expected raw program %s to parse to %#v, got %#v", tC.input, tC.want, got)
		})
	}
}

func TestRunProgram(t *testing.T) {
	testCases := []struct {
		desc    string
		input   *program
		wantAcc int
		wantPC  int
		wantErr error
	}{
		{
			"nop does nothing but increment pc",
			&program{instructions: []instruction{{op: opNop, arg: "+0"}}},
			0,
			1,
			nil,
		},
		{
			"positive acc increases accumulator",
			&program{instructions: []instruction{{op: opAcc, arg: "+1"}}},
			1,
			1,
			nil,
		},
		{
			"negative acc decreases accumulator",
			&program{instructions: []instruction{{op: opAcc, arg: "-1"}}},
			-1,
			1,
			nil,
		},
		{
			"jumps set program counter",
			&program{instructions: []instruction{{op: opJmp, arg: "+4"}}},
			0,
			4,
			nil,
		},
		{
			"infinite loops cause errors",
			&program{instructions: []instruction{
				{"nop", "+0"},
				{"acc", "+1"},
				{"jmp", "+4"},
				{"acc", "+3"},
				{"jmp", "-3"},
				{"acc", "-99"},
				{"acc", "+1"},
				{"jmp", "-4"},
				{"acc", "+6"},
			}},
			0,
			4,
			errInfiniteLoop,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			gotErr := tC.input.run()
			if tC.wantErr != nil {
				assert.Equal(t, tC.wantErr, gotErr)
				return
			}
			assert.Equal(t, tC.wantAcc, tC.input.acc)
			assert.Equal(t, tC.wantPC, tC.input.pc)
		})
	}
}

func TestFixProgram(t *testing.T) {
	inputRaw := `nop +0
acc +1
jmp +4
acc +3
jmp -3
acc -99
acc +1
jmp -4
acc +6
`
	inputProgram := parseProgram(inputRaw)

	fixed := fixProgram(inputProgram)
	got := fixed.acc
	want := 8

	assert.Equal(t, want, got)
}
