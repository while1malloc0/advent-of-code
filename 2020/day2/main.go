package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/while1malloc0/advent-of-code/2020/challenge"
)

type pwEntry struct {
	min  int
	max  int
	char rune
	pw   string
}

// thanks regex101.com
var pwEntryRegex = regexp.MustCompile("(\\d+)-(\\d+)\\s([a-z]):\\s([a-z]+)")

func parsePWEntry(in string) (pwEntry, error) {
	matches := pwEntryRegex.FindStringSubmatch(in)
	if len(matches) < 1 {
		return pwEntry{}, fmt.Errorf("Could not parse pw entry from [%s]", in)
	}
	// first match (match[0]) is the whole string, submatches are 1+
	min, err := strconv.Atoi(matches[1])
	if err != nil {
		return pwEntry{}, err
	}
	max, err := strconv.Atoi(matches[2])
	if err != nil {
		return pwEntry{}, err
	}
	char := []rune(matches[3])[0]
	pw := matches[4]
	return pwEntry{min: min, max: max, char: char, pw: pw}, nil
}

func minMaxCharAllowed(policy pwEntry) bool {
	var count int
	for _, c := range policy.pw {
		if c == policy.char {
			count++
		}
		if count > policy.max {
			return false
		}
	}
	return count >= policy.min
}

// a xor b <-> a != b
func positionXOr(policy pwEntry) bool {
	firstPositionMatch := rune(policy.pw[policy.min-1]) == policy.char
	secondPositionMatch := rune(policy.pw[policy.max-1]) == policy.char
	return firstPositionMatch != secondPositionMatch
}

func passwordIsValid(policy pwEntry, validatorFunc func(pwEntry) bool) bool {
	return validatorFunc(policy)
}

func main() {
	var ins []pwEntry
	challenge.InputScanFunc("input", func(s string) error {
		param, err := parsePWEntry(s)
		if err != nil {
			return err
		}
		ins = append(ins, param)
		return nil
	})
	partOneFn := func() error {
		var count int
		for _, entry := range ins {
			if passwordIsValid(entry, minMaxCharAllowed) {
				count++
			}
		}
		fmt.Println(count)
		return nil
	}
	partTwoFn := func() error {
		var count int
		for _, entry := range ins {
			if passwordIsValid(entry, positionXOr) {
				count++
			}
		}
		fmt.Println(count)
		return nil
	}
	challenge.Run(partOneFn, partTwoFn)
}
