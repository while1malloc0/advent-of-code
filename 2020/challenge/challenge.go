package challenge

import (
	"bufio"
	"flag"
	"os"
	"strings"
)

var part *int = flag.Int("part", 1, "Which part of the puzzle to do")

func Setup() int {
	flag.Parse()
	return *part
}

func Run(partOneFn func() error, partTwoFn func() error) error {
	Setup()
	if *part == 2 {
		return partTwoFn()
	}
	return partOneFn()
}

func InputScanFunc(fname string, fn func(string) error) error {
	f, err := os.OpenFile(fname, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return err
	}
	s := bufio.NewScanner(f)
	for s.Scan() {
		line := s.Text()
		line = strings.TrimSpace(line)
		err := fn(line)
		if err != nil {
			return err
		}
	}
	return nil
}

// ActualMod makes up for the fact that Go's mod behavior is suspect
func ActualMod(x, y int) int {
	return ((x % y) + y) % y
}
