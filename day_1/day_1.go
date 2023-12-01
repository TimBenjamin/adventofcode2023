package day_1

import (
	"adventofcode2023/util"
	"strconv"
)

var input []string

func partOne() int {
	// TODO
	return 0
}

func partTwo() int {
	// TODO
	return 0
}

func Call(part string, inputFile string) string {
	input = util.ParseInputIntoLines(inputFile)
	var r int
	if part == "1" {
		r = partOne()
	} else {
		r = partTwo()
	}
	return strconv.Itoa(r)
}
