package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	part = flag.Int("part", 1, "the part to run")
	file = flag.String("file", "example.txt", "the file to read")
)

func PartOne(input []byte) int {
	return 0
}

func PartTwo(input []byte) int {
	return 0
}

func main() {
	flag.Parse()

	in, err := os.ReadFile(*file)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	switch *part {
	case 1:
		fmt.Println(PartOne(in))
	case 2:
		fmt.Println(PartTwo(in))
	}
}
