package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/while1malloc0/advent-of-code/2020/challenge"
)

type numberRange struct {
	start int
	end   int
}

func (n *numberRange) contains(guess int) bool {
	return n.start <= guess && guess <= n.end
}

type numberRangeGroup struct {
	name   string
	groups []*numberRange
	memo   map[int]bool
}

func (n *numberRangeGroup) contains(guess int) bool {
	// Is it a good idea to lazy initialize the memo map? Probably not
	if n.memo == nil {
		n.memo = map[int]bool{}
	}

	if seen, ok := n.memo[guess]; ok {
		return seen
	}

	result := false
	for i := range n.groups {
		if n.groups[i].contains(guess) {
			result = true
			break
		}
	}

	n.memo[guess] = result
	return n.memo[guess]
}

type ticket = []int

const NoInvalid = -1

type ticketValidation struct {
	ticket ticket
	ranges *numberRangeGroup
}

func (t *ticketValidation) findInvalid() int {
	for i := range t.ticket {
		if !t.ranges.contains(t.ticket[i]) {
			return t.ticket[i]
		}
	}

	return NoInvalid
}

func parseNumberRangeGroups(in string) []*numberRangeGroup {
	topHalfOfTicket := strings.Split(in, "your ticket:\n")[0]
	topHalfOfTicket = strings.TrimSpace(topHalfOfTicket)
	r := strings.NewReader(topHalfOfTicket)
	s := bufio.NewScanner(r)
	ngs := []*numberRangeGroup{}
	for s.Scan() {
		ng := &numberRangeGroup{groups: []*numberRange{}}
		line := s.Text()
		parts := strings.Split(line, ":")
		ng.name = strings.TrimSpace(parts[0])
		rangesLine := parts[1]
		rangesLine = strings.TrimSpace(rangesLine)
		rangesStr := strings.Split(rangesLine, "or")

		for i := range rangesStr {
			split := strings.Split(rangesStr[i], "-")
			lowerStr, upperStr := split[0], split[1]
			lowerStr = strings.TrimSpace(lowerStr)
			upperStr = strings.TrimSpace(upperStr)
			lower, err := strconv.Atoi(lowerStr)
			if err != nil {
				panic(err)
			}
			upper, err := strconv.Atoi(upperStr)
			if err != nil {
				panic(err)
			}

			ng.groups = append(ng.groups, &numberRange{start: lower, end: upper})
		}

		ngs = append(ngs, ng)
	}

	return ngs
}

func parseTickets(in string) []ticket {
	lowerHalfOfTicket := strings.Split(in, "nearby tickets:\n")[1]
	r := strings.NewReader(lowerHalfOfTicket)
	s := bufio.NewScanner(r)

	var tickets []ticket
	for s.Scan() {
		line := s.Text()
		nums := strings.Split(line, ",")
		t := ticket{}
		for _, numStr := range nums {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				panic(err)
			}
			t = append(t, num)
		}
		tickets = append(tickets, t)
	}

	return tickets
}

func filterInvalidTickets(in []ticket, ng *numberRangeGroup) []ticket {
	var out []ticket
	for i := range in {
		tv := &ticketValidation{ticket: in[i], ranges: ng}
		invalid := tv.findInvalid()
		if invalid == NoInvalid {
			out = append(out, in[i])
		}
	}
	return out
}

func getColumns(in []ticket) [][]int {
	var out [][]int

	for column := 0; column < len(in[0]); column++ {
		var row []int
		for i := range in {
			row = append(row, in[i][column])
		}
		out = append(out, row)
	}

	return out
}

func getFieldPositions(in []ticket, ngs []*numberRangeGroup) []string {
	cols := getColumns(in)

	// Set-ish thing
	positions := map[string]map[int]struct{}{}

	for i, col := range cols {
		t := ticket(col)
		for j := range ngs {
			tv := &ticketValidation{ticket: t, ranges: ngs[j]}
			invalid := tv.findInvalid()
			if invalid == NoInvalid {
				if positions[ngs[j].name] == nil {
					positions[ngs[j].name] = map[int]struct{}{}
				}
				positions[ngs[j].name][i] = struct{}{}
			}
		}
	}

	for {
		foundAll := true
		for i := range positions {
			if len(positions[i]) > 1 {
				foundAll = false
				break
			}
		}
		if foundAll {
			break
		}
		for i := range positions {
			if len(positions[i]) == 1 {
				for val := range positions[i] {
					for j := range positions {
						if j == i {
							continue
						}
						if len(positions[j]) == 1 {
							continue
						}
						delete(positions[j], val)
					}
				}
			}
		}
	}

	out := []string{}
	for range positions {
		out = append(out, "")
	}

	for k := range positions {
		for i := range positions[k] {
			out[i] = k
		}
	}

	return out
}

func parseMyTicket(in string) ticket {
	bottomHalf := strings.Split(in, "your ticket:\n")[1]
	numStr := strings.Split(bottomHalf, "\n")[0]
	numStrs := strings.Split(numStr, ",")
	var t ticket
	for i := range numStrs {
		num, err := strconv.Atoi(numStrs[i])
		if err != nil {
			panic(err)
		}
		t = append(t, num)
	}

	return t
}

func main() {
	partOneFunc := func() error {
		in, err := ioutil.ReadFile("input")
		if err != nil {
			return err
		}
		ngs := parseNumberRangeGroups(string(in))
		tickets := parseTickets(string(in))

		var sum int
		for i := range tickets {
			for j := range ngs {
				tv := &ticketValidation{ranges: ngs[j], ticket: tickets[i]}
				invalid := tv.findInvalid()
				if invalid != NoInvalid {
					sum += invalid
				}
			}
		}

		fmt.Println(sum)

		return nil
	}

	partTwoFunc := func() error {
		in, err := ioutil.ReadFile("input")
		if err != nil {
			return err
		}
		ngs := parseNumberRangeGroups(string(in))
		tickets := parseTickets(string(in))

		var validTickets []ticket

		for i := range tickets {
			var invalidFor int
			for j := range ngs {
				tv := &ticketValidation{ticket: tickets[i], ranges: ngs[j]}
				if tv.findInvalid() != NoInvalid {
					invalidFor++
				}
			}
			if invalidFor == len(ngs) {
				continue
			}
			validTickets = append(validTickets, tickets[i])
		}

		positions := getFieldPositions(validTickets, ngs)

		myTicket := parseMyTicket(string(in))

		result := 1
		for i := range positions {
			if strings.HasPrefix(positions[i], "departure") {
				result *= myTicket[i]
			}
		}
		fmt.Println(result)
		return nil
	}
	challenge.Run(partOneFunc, partTwoFunc)
}
