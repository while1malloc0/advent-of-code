package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/while1malloc0/advent-of-code/2021/days"
)

var (
	day  = flag.Int("day", 0, "the day to run")
	part = flag.Int("part", 1, "the part of the day to run")
)

func main() {
	flag.Parse()

	if day == nil || *day == 0 || part == nil {
		usage()
	}

	answer, err := days.Run(*day, *part)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(answer)
}

func usage() {
	fmt.Println("Usage: go run main.go -day DAY -part PART")
	os.Exit(1)
}
