package main

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
	"unicode"

	"github.com/while1malloc0/advent-of-code/2020/challenge"
)

func parseRegexMap(in string) map[string]string {
	out := map[string]string{}
	r := strings.NewReader(in)
	s := bufio.NewScanner(r)
	for s.Scan() {
		line := s.Text()
		parts := strings.Split(line, ":")
		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])
		value = strings.ReplaceAll(value, "\"", "")
		out[key] = value
	}
	return out
}

func isTerminalRule(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func parseRegexDeprecated(in string) string {
	m := parseRegexMap(in)
	expr := m["0"]
	for {
		allReplaced := true
		for _, char := range expr {
			if unicode.IsNumber(char) {
				allReplaced = false
				break
			}
		}
		if allReplaced {
			break
		}
		chars := strings.Split(expr, " ")
		for _, s := range chars {
			if s == "(" || s == ")" || s == "|" || s == "" {
				continue
			}
			if !unicode.IsLetter(rune(s[0])) {
				replacement := m[s]
				if isTerminalRule(replacement) {
					expr = strings.ReplaceAll(expr, s, replacement)
				} else {
					expr = strings.ReplaceAll(expr, s, " ( "+m[s]+" ) ")
				}
			}
		}
	}
	expr = strings.ReplaceAll(expr, " ", "")
	expr += "$"
	return expr
}

func parseRegexHelper(m map[string]string, in string) string {
	if in == "|" {
		print(in)
	}
	if in == "a" || in == "b" || in == "|" {
		return in
	}

	var out string
	for _, s := range strings.Split(in, " ") {
		char := strings.TrimSpace(s)
		if char == "|" {
			out += char
			continue
		}
		if char == "" {
			continue
		}
		if !isTerminalRule(m[char]) {
			out += "("
		}
		out += parseRegexHelper(m, m[string(char)])
		if !isTerminalRule(m[char]) {
			out += ")"
		}
	}
	return out
}

func parseRegex(in string) string {
	m := parseRegexMap(in)
	out := "^"
	out += parseRegexHelper(m, m["0"])
	out += "$"
	return out
}

var memo map[string]string

func testMatch(reg string, test string) bool {
	// This should probably take a parsed regex, but it's Saturday, so just
	// memoize instead
	if memo == nil {
		memo = map[string]string{}
	}
	parsed, ok := memo[reg]
	if !ok {
		newParsed := parseRegex(reg)
		memo[reg] = newParsed
		parsed = memo[reg]
	}
	matched, err := regexp.MatchString(parsed, test)
	if err != nil {
		panic(err)
	}
	return matched
}

func main() {
	partOneFunc := func() error {
		in, err := ioutil.ReadFile("input")
		if err != nil {
			return err
		}
		parts := strings.Split(string(in), "\n\n")
		if len(parts) != 2 {
			return errors.New("Did not parse input correctly")
		}
		expression := strings.TrimSpace(parts[0])
		tests := strings.TrimSpace(parts[1])

		var sum int
		for _, s := range strings.Split(tests, "\n") {
			if testMatch(expression, s) {
				sum++
			}
		}
		fmt.Println(sum)

		return nil
	}

	challenge.Run(partOneFunc, nil)
}
