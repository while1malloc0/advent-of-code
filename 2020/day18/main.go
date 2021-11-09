package main

import (
	"fmt"
	"strconv"

	"github.com/while1malloc0/advent-of-code/2020/challenge"
)

type evalStrategy int

type intStack []int

func (s intStack) pop() (int, intStack) {
	result := s[len(s)-1]
	s = s[:len(s)-1]
	return result, s
}

func (s intStack) push(elem int) intStack {
	s = append(s, elem)
	return s
}

type opStack []string

func (s opStack) pop() (string, opStack) {
	result := s[len(s)-1]
	s = s[:len(s)-1]
	return result, s
}

func (s opStack) push(elem string) opStack {
	s = append(s, elem)
	return s
}

type queueStack []queue

func (s queueStack) pop() (queue, queueStack) {
	result := s[len(s)-1]
	s = s[:len(s)-1]
	return result, s
}

func (s queueStack) push(elem queue) queueStack {
	s = append(s, elem)
	return s
}

const (
	opAdd  = "+"
	opMult = "*"

	evalStrategyNoPrecedence = iota
	evalStrategyWithPrecedence
)

type queue []string

func (q queue) eval() int {
	ns := intStack{}
	os := opStack{}

	var op string
	var lhs int
	var rhs int
	for _, char := range q {
		if len(ns) > 1 {
			lhs, ns = ns.pop()
			rhs, ns = ns.pop()
			op, os = os.pop()
			if op == opAdd {
				ns = ns.push(lhs + rhs)
			} else if op == opMult {
				ns = ns.push(lhs * rhs)
			} else {
				panic("invalid op found in op stack")
			}
		}
		if char == opAdd || char == opMult {
			os = os.push(char)
		} else {
			n, err := strconv.Atoi(string(char))
			if err != nil {
				panic(err)
			}
			ns = ns.push(n)
		}
	}

	for len(ns) > 1 {
		lhs, ns = ns.pop()
		rhs, ns = ns.pop()
		op, os = os.pop()
		if op == opAdd {
			ns = ns.push(lhs + rhs)
		} else if op == opMult {
			ns = ns.push(lhs * rhs)
		} else {
			panic("invalid op in op stack")
		}
	}

	result, _ := ns.pop()
	return result
}

func (q queue) evalWithPrecedence() int {
	q = q.evalAdd()
	return q.eval()
}

func (q queue) evalAdd() queue {
	out := queue{}
	itsPoppinTime := false
	for _, char := range q {
		if itsPoppinTime {
			res := out[len(out)-1]
			out = out[:len(out)-1]
			num, err := strconv.Atoi(res)
			if err != nil {
				panic(err)
			}
			charNum, err := strconv.Atoi(char)
			if err != nil {
				panic(err)
			}
			result := num + charNum
			out = append(out, strconv.Itoa(result))
			itsPoppinTime = false
			continue
		}
		if char == "+" {
			itsPoppinTime = true
			continue
		}
		out = append(out, char)
	}
	return out
}

func eval(in string, strategy evalStrategy) int {
	qs := queueStack{}
	q := queue{}
	qs = qs.push(q)
	var toEval queue

	for _, c := range in {
		char := string(c)
		if char == " " {
			continue
		}
		if char == "(" {
			q = queue{}
			qs = qs.push(q)
			continue
		}
		if char == ")" {
			toEval, qs = qs.pop()
			var result int
			if strategy == evalStrategyNoPrecedence {
				result = toEval.eval()
			} else if strategy == evalStrategyWithPrecedence {
				result = toEval.evalWithPrecedence()
			} else {
				panic("Invalid eval strategy")
			}
			s := strconv.Itoa(result)
			qs[len(qs)-1] = append(qs[len(qs)-1], s)
			continue
		}
		qs[len(qs)-1] = append(qs[len(qs)-1], char)
	}
	if strategy == evalStrategyNoPrecedence {
		return qs[len(qs)-1].eval()
	}

	return qs[len(qs)-1].evalWithPrecedence()
}

func main() {
	var ins []string

	challenge.InputScanFunc("input", func(s string) error {
		ins = append(ins, s)
		return nil
	})

	partOneFunc := func() error {
		var sum int
		for i := range ins {
			sum += eval(ins[i], evalStrategyNoPrecedence)
		}
		fmt.Println(sum)
		return nil
	}

	partTwoFunc := func() error {
		var sum int
		for i := range ins {
			sum += eval(ins[i], evalStrategyWithPrecedence)
		}
		fmt.Println(sum)
		return nil
	}

	challenge.Run(partOneFunc, partTwoFunc)
}
