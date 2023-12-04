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
			fmt.Printf("found a symbol: %v\n", c)
			return true
		}
	}
	return false
}

func partOne() int {
	total := 0
	re := regexp.MustCompile("[^0-9]+")
	for lineNum, line := range input {
		numbers := re.Split(line, -1)
		for _, number := range numbers {

			// check around the number to see if there are symbols
			// Bug here - the same number can appear twice in a line :-(
			idx := strings.Index(line, number)
			if len(number) == 0 {
				continue
			}
			fmt.Printf("checking number: %v\n", number)

			var charsToCheck []string
			from := idx
			to := idx + len(number) - 1
			if from > 0 {
				from--
				charsToCheck = append(charsToCheck, line[from:from+1])
				fmt.Printf("from it is: %v\n", line[from:from+1])
			}
			if to < len(line)-1 {
				to++
				charsToCheck = append(charsToCheck, line[to:to+1])
				fmt.Printf("to it is: %v\n", line[to:to+1])
			}
			fmt.Printf("from: %v / to: %v\n", from, to)

			// check line above
			if lineNum > 0 {
				snip := strings.Split(input[lineNum-1][from:to+1], ".")
				charsToCheck = append(charsToCheck, snip...)
				fmt.Printf("linB: %v\n", input[lineNum-1])
			}
			fmt.Printf("line: %v\n", line)
			// check line below
			if lineNum < len(input)-1 {
				snip := strings.Split(input[lineNum+1][from:to+1], ".")
				charsToCheck = append(charsToCheck, snip...)
				fmt.Printf("linA: %v\n", input[lineNum+1])
			}

			partNum, _ := strconv.Atoi(number)
			if checkForSymbol(charsToCheck) {
				fmt.Printf("number is a part: %v\n\n", partNum)
				total += partNum
			} else {
				fmt.Printf("number IS NOT a part: %v\n\n", partNum)
			}
		}

	}

	return total
}

func partTwo() int {
	total := 0
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
