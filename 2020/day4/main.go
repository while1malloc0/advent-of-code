package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"

	"github.com/while1malloc0/advent-of-code/2020/challenge"
)

type passport = map[string]string

type passportValidationFuncs = map[string]func(string) bool

var existenceValidatorFuncs = passportValidationFuncs{
	"byr": func(string) bool { return true },
	"iyr": func(string) bool { return true },
	"eyr": func(string) bool { return true },
	"hgt": func(string) bool { return true },
	"hcl": func(string) bool { return true },
	"ecl": func(string) bool { return true },
	"pid": func(string) bool { return true },
}

var partTwoValidatorFuncs = passportValidationFuncs{
	"byr": func(s string) bool {
		parsed, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		return len(s) == 4 && parsed >= 1920 && parsed <= 2002
	},
	"iyr": func(s string) bool {
		parsed, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		return len(s) == 4 && parsed >= 2010 && parsed <= 2020
	},
	"eyr": func(s string) bool {
		parsed, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		return len(s) == 4 && parsed >= 2020 && parsed <= 2030
	},
	"hgt": func(s string) bool {
		if strings.HasSuffix(s, "cm") {
			parts := strings.Split(s, "cm")
			num, err := strconv.Atoi(parts[0])
			if err != nil {
				panic(err)
			}
			return num >= 150 && num <= 193
		}
		if strings.HasSuffix(s, "in") {
			parts := strings.Split(s, "in")
			num, err := strconv.Atoi(parts[0])
			if err != nil {
				panic(err)
			}
			return num >= 59 && num <= 76
		}
		return false
	},
	"hcl": func(s string) bool {
		return regexp.MustCompile("#[0-9a-f]{6}").MatchString(s)
	},
	"ecl": func(s string) bool {
		allowedEyeColors := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
		for _, color := range allowedEyeColors {
			if s == color {
				return true
			}
		}
		return false
	},
	"pid": func(s string) bool {
		return regexp.MustCompile("[0-9]{9}").MatchString(s)
	},
}

func passportValid(p passport, validators passportValidationFuncs) bool {
	requiredKeys := map[string]bool{}
	for k := range validators {
		requiredKeys[k] = false
	}
	for k := range p {
		requiredKeys[k] = true
		_, ok := validators[k]
		if !ok {
			continue
		}
		if !validators[k](p[k]) {
			return false
		}
	}
	for _, set := range requiredKeys {
		if !set {
			return false
		}
	}
	return true
}

func parseBatchFile(batchFile string) []passport {
	// This returns an extra passport in some cases. :shrug:
	var out []passport

	currentPassport := passport{}

	r := strings.NewReader(batchFile)
	s := bufio.NewScanner(r)
	for s.Scan() {
		row := s.Text()
		if row == "" {
			out = append(out, currentPassport)
			currentPassport = passport{}
			continue
		}
		entries := strings.Split(row, " ")
		for _, entry := range entries {
			parts := strings.Split(entry, ":")
			field, value := parts[0], parts[1]
			currentPassport[field] = value
		}
	}
	out = append(out, currentPassport)

	return out
}

func main() {
	f, err := ioutil.ReadFile("input")
	if err != nil {
		panic(err)
	}
	in := strings.TrimSpace(string(f))
	partOneFunc := func() error {
		var count int
		passports := parseBatchFile(in)
		for _, p := range passports {
			if passportValid(p, existenceValidatorFuncs) {
				count++
			}
		}
		fmt.Println(count)
		return nil
	}
	partTwoFunc := func() error {
		var count int
		passports := parseBatchFile(in)
		for _, p := range passports {
			if passportValid(p, partTwoValidatorFuncs) {
				count++
			}
		}
		fmt.Println(count)
		return nil
	}

	err = challenge.Run(partOneFunc, partTwoFunc)
	if err != nil {
		panic(err)
	}
}
