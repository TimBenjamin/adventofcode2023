package day_2

import (
	"adventofcode2023/util"
	"fmt"
	"strconv"
	"strings"
)

var input []string

const MAX_RED = 12
const MAX_GREEN = 13
const MAX_BLUE = 14

func partOne() int {
	total := 0
	for _, game := range input {
		gameIsPossible := true
		halves := strings.Split(game, ": ")
		id, _ := strconv.Atoi(strings.Split(halves[0], " ")[1])
		pulls := strings.Split(halves[1], "; ")
	pullLoop:
		for _, pull := range pulls {
			reveals := strings.Split(pull, ", ")
			for _, reveal := range reveals {
				parts := strings.Split(reveal, " ")
				num, _ := strconv.Atoi(parts[0])
				if parts[1] == "red" && num > MAX_RED || parts[1] == "green" && num > MAX_GREEN || parts[1] == "blue" && num > MAX_BLUE {
					// fmt.Printf("game %v is impossible\n\n", id)
					gameIsPossible = false
					break pullLoop
				}
			}
		}
		if gameIsPossible {
			total += id
			fmt.Printf("game %v is possible! add %v, total is now %v\n\n", id, id, total)
		}
	}
	return total
}

func partTwo() int {
	total := 0
	for _, game := range input {
		halves := strings.Split(game, ": ")
		pulls := strings.Split(halves[1], "; ")
		peakRed, peakGreen, peakBlue := 0, 0, 0
		for _, pull := range pulls {
			reveals := strings.Split(pull, ", ")
			for _, reveal := range reveals {
				parts := strings.Split(reveal, " ")
				num, _ := strconv.Atoi(parts[0])
				if parts[1] == "red" && num > peakRed {
					peakRed = num
				}
				if parts[1] == "green" && num > peakGreen {
					peakGreen = num
				}
				if parts[1] == "blue" && num > peakBlue {
					peakBlue = num
				}
			}
		}
		power := peakRed * peakGreen * peakBlue
		total += power
	}
	return total
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
