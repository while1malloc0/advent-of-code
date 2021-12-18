package days

import (
	"github.com/while1malloc0/advent-of-code/2021/util"
)

type dayFunc = func(string) (interface{}, error)

var funcs map[[2]int]dayFunc = map[[2]int]dayFunc{
	{11, 1}: Day11Part1,
}

func Run(day, part int) (interface{}, error) {
	input, err := util.ReadDay(day)
	if err != nil {
		return nil, err
	}
	return funcs[[2]int{day, part}](input)
}
