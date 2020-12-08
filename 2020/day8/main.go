package main

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/while1malloc0/advent-of-code/2020/challenge"
)

type opType string

const (
	opAcc = "acc"
	opNop = "nop"
	opJmp = "jmp"
)

var errInfiniteLoop = errors.New("Infinite loop detected in program")

type instruction struct {
	op  opType
	arg string
}

type program struct {
	// accumulator
	acc int
	// program counter
	pc           int
	instructions []instruction
}

func (p *program) run() error {
	seen := map[int]bool{}
	for {
		// usually we'd segfault or something, but just breaking works for now
		if p.pc > len(p.instructions)-1 {
			break
		}
		if seen, ok := seen[p.pc]; ok && seen {
			return errInfiniteLoop
		}
		seen[p.pc] = true
		inst := p.instructions[p.pc]
		switch inst.op {
		case opNop:
			p.pc++
		case opAcc:
			val, err := strconv.Atoi(inst.arg)
			if err != nil {
				// C-style error handling, aka just blow the whole program up
				panic(fmt.Sprintf("Invalid arg for op acc: %v", err))
			}
			p.acc += val
			p.pc++
		case opJmp:
			val, err := strconv.Atoi(inst.arg)
			if err != nil {
				panic(fmt.Sprintf("Invalid arg for op jmp: %v", err))
			}
			// jmp 0 is "do this forever"
			if val == 0 {
				return errInfiniteLoop
			}
			p.pc += val
		default:
			panic(fmt.Sprintf("Invalid op for program: %s", inst.op))
		}
	}
	return nil
}

func parseProgram(in string) *program {
	r := strings.NewReader(in)
	s := bufio.NewScanner(r)
	var instructions []instruction
	for s.Scan() {
		txt := s.Text()
		parts := strings.Split(txt, " ")
		opRaw := strings.TrimSpace(parts[0])
		argRaw := strings.TrimSpace(parts[1])

		var op opType
		switch opRaw {
		case "acc":
			op = opAcc
		case "nop":
			op = opNop
		case "jmp":
			op = opJmp
		default:
			panic(fmt.Sprintf("Unsupported instruction found in program: %s", opRaw))
		}

		instructions = append(instructions, instruction{arg: argRaw, op: op})
	}
	return &program{instructions: instructions}
}

func flipInstruction(p *program, pos int) (*program, error) {
	flipped := &program{
		acc:          p.acc,
		pc:           p.pc,
		instructions: []instruction{},
	}
	for i := range p.instructions {
		flipped.instructions = append(flipped.instructions, p.instructions[i])
	}
	inst := flipped.instructions[pos]
	switch inst.op {
	case opJmp:
		flipped.instructions[pos] = instruction{op: opNop, arg: inst.arg}
	case opNop:
		flipped.instructions[pos] = instruction{op: opJmp, arg: inst.arg}
	default:
		return nil, errors.New("Not a flippable instruction")
	}
	return flipped, nil
}

func fixProgram(p *program) *program {
	var pos int
	for {
		attempt, err := flipInstruction(p, pos)
		if err != nil {
			pos++
			continue
		}
		err = attempt.run()
		if err != nil {
			pos++
			continue
		}
		return attempt
	}
}

func main() {
	partOneFunc := func() error {
		rawProgram, err := ioutil.ReadFile("input")
		if err != nil {
			return err
		}
		p := parseProgram(string(rawProgram))
		err = p.run()
		if err == errInfiniteLoop {
			fmt.Println(p.acc)
			return nil
		}
		return errors.New("Did not find expected inf loop")
	}

	partTwoFunc := func() error {
		rawProgram, err := ioutil.ReadFile("input")
		if err != nil {
			return err
		}
		p := parseProgram(string(rawProgram))
		fixed := fixProgram(p)
		err = fixed.run()
		if err != nil {
			return err
		}
		fmt.Println(fixed.acc)
		return nil
	}
	challenge.Run(partOneFunc, partTwoFunc)
}
