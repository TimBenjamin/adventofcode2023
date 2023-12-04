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

func partOne() int {
	total := 0
	re := regexp.MustCompile(`[0-9]+`)
	for lineNum, line := range input {
		numberLocations := re.FindAllStringSubmatchIndex(line, -1)
		for _, location := range numberLocations {
			number := line[location[0]:location[1]]
			// check around the number to see if there are symbols
			var charsToCheck []string
			from := location[0]
			to := location[1]
			// fmt.Printf("number %v is from: %v to: %v\n", number, from, to)

			// need to get any chars either side of it:
			if from > 0 {
				from--
				charsToCheck = append(charsToCheck, line[from:from+1])
			}
			if to < len(line) {
				to++
				charsToCheck = append(charsToCheck, line[to-1:to])
			}

			// the expanded snippet horizontal positions
			// fmt.Printf("expanded: from: %v / to: %v\n", from, to)

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

func getPartNumbersAt(lineNum int, position int) []int {
	re := regexp.MustCompile(`[0-9]+`)
	numberLocations := re.FindAllStringSubmatchIndex(input[lineNum], -1)
	partNumbers := []int{}
	for _, loc := range numberLocations {
		if position >= loc[0]-1 && position <= loc[1] {
			partNumber, _ := strconv.Atoi(input[lineNum][loc[0]:loc[1]])
			partNumbers = append(partNumbers, partNumber)
		}
	}
	return partNumbers
}

func partTwo() int {
	total := 0

	// now find all asterisks and see if exactly two numbers are adjacent
	re := regexp.MustCompile(`[0-9]+`)
	for lineNum, line := range input {
		for pos, c := range line {
			if c == '*' {
				from, to := pos, pos
				var partNumbersFound []int
				if pos > 0 {
					if re.MatchString(string(line[pos-1])) {
						partNumbersFound = append(partNumbersFound, getPartNumbersAt(lineNum, pos-1)...)
					}
					from--
				}
				if pos < len(line)-1 {
					if re.MatchString(string(line[pos+1])) {
						partNumbersFound = append(partNumbersFound, getPartNumbersAt(lineNum, pos+1)...)
					}
					to++
				}
				if lineNum > 0 {
					if re.MatchString(input[lineNum-1][from : to+1]) {
						partNumbersFound = append(partNumbersFound, getPartNumbersAt(lineNum-1, pos)...)
					}
				}
				if lineNum < len(input)-1 {
					if re.MatchString(input[lineNum+1][from : to+1]) {
						partNumbersFound = append(partNumbersFound, getPartNumbersAt(lineNum+1, pos)...)
					}
				}
				if len(partNumbersFound) == 2 {
					// fmt.Printf("found a gear at pos: %v in line: %v\n", pos, lineNum)
					// fmt.Printf("part numbers: %v\n", partNumbersFound)
					// fmt.Printf("ratio: %v\n", partNumbersFound[0]*partNumbersFound[1])
					total += (partNumbersFound[0] * partNumbersFound[1])
				}
			}
		}
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
