package main

import (
	"fmt"
	"sort"

	"github.com/while1malloc0/advent-of-code/2020/challenge"
)

func getSeatID(row, col int) int {
	return (row * 8) + col
}

func getRow(s string) int {
	var seats []int
	max := 127
	for i := 0; i <= max; i++ {
		seats = append(seats, i)
	}
	var pivot int
	for _, char := range s {
		pivot = len(seats) / 2
		if char == 'F' {
			seats = seats[0:pivot]
		} else if char == 'B' {
			seats = seats[pivot:len(seats)]
		} else {
			panic("Uh, not a number")
		}
	}
	return seats[0]
}

func getColumn(s string) int {
	var columns []int
	max := 7
	for i := 0; i <= max; i++ {
		columns = append(columns, i)
	}
	var pivot int
	for _, char := range s {
		pivot = len(columns) / 2
		if char == 'L' {
			columns = columns[0:pivot]
		} else if char == 'R' {
			columns = columns[pivot:len(columns)]
		} else {
			panic("Uh, not a column direction")
		}
	}
	return columns[0]
}

func seatIDFromBoardingPass(s string) int {
	rowDirection := s[0:7]
	columnDirection := s[7:len(s)]
	row := getRow(rowDirection)
	col := getColumn(columnDirection)
	return getSeatID(row, col)
}

func main() {
	var ins []string
	challenge.InputScanFunc("input", func(s string) error {
		ins = append(ins, s)
		return nil
	})
	partOneFunc := func() error {
		var max int
		for _, boardingPass := range ins {
			seatID := seatIDFromBoardingPass(boardingPass)
			if seatID > max {
				max = seatID
			}
		}
		fmt.Println(max)
		return nil
	}

	partTwoFunc := func() error {
		var seatIDs []int
		for _, boardingPass := range ins {
			seatIDs = append(seatIDs, seatIDFromBoardingPass(boardingPass))
		}
		sort.Ints(seatIDs)
		for i := 1; i < len(seatIDs)-1; i++ {
			if seatIDs[i+1]-1 != seatIDs[i] {
				fmt.Println(seatIDs[i] + 1)
			}
		}
		return nil
	}
	challenge.Run(partOneFunc, partTwoFunc)
}
