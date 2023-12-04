package day_3

import (
	"adventofcode2023/util"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var input []string

func checkForSymbol(snip []string) bool {
	numberRe := regexp.MustCompile("[0-9]")
	for _, c := range snip {
		if c != "." && len(c) > 0 && !numberRe.MatchString(c) {
			return true
		}
	}
	return false
}

func checkForAsterisk(snip []string) bool {
	for _, c := range snip {
		if c == "*" {
			return true
		}
	}
	return false
}

func partOne() int {
	total := 0
	re := regexp.MustCompile("[0-9]+")
	for lineNum, line := range input {
		numberLocations := re.FindAllStringSubmatchIndex(line, -1)
		for _, location := range numberLocations {
			number := line[location[0]:location[1]]
			// check around the number to see if there are symbols
			var charsToCheck []string
			from := location[0]
			to := location[1]
			fmt.Printf("number %v is from: %v to: %v\n", number, from, to)

			// need to get any chars either side of it:
			if from > 0 {
				from--
				charsToCheck = append(charsToCheck, line[from:from+1])
				fmt.Printf("char before: %v\n", line[from:from+1])
			}
			fmt.Printf("len of line: %v\n", len(line))
			if to < len(line) {
				to++
				charsToCheck = append(charsToCheck, line[to-1:to])
				fmt.Printf("char after: %v\n", line[to-1:to])
			}

			// the expanded snippet horizontal positions
			fmt.Printf("expanded: from: %v / to: %v\n", from, to)

			// check line above
			if lineNum > 0 {
				snip := strings.Split(input[lineNum-1][from:to], ".")
				charsToCheck = append(charsToCheck, snip...)
			}
			// check line below
			if lineNum < len(input)-1 {
				snip := strings.Split(input[lineNum+1][from:to], ".")
				charsToCheck = append(charsToCheck, snip...)
			}

			// do the output
			if checkForSymbol(charsToCheck) {
				fmt.Printf("number is a part: %v\n\n", number)
				partNum, _ := strconv.Atoi(number)
				total += partNum
			} else {
				fmt.Printf("number IS NOT a part: %v\n\n", number)
			}
		}
	}
	return total
}

func partTwo() int {
	total := 0

	// find pairs of numbers that are adjacent to the same * character

	// find *'s and then find numbers that are adjacent to them?

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
