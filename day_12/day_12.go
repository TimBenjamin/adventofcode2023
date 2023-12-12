package day_12

import (
	"adventofcode2023/util"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Record struct {
	springs    string
	springSets []string
	report     []int
}

var records []Record

func permutations(n, r int) int {
	if n < r {
		return 0
	}
	result := 1
	for i := 0; i < r; i++ {
		result *= (n - i)
	}
	return result
}

func countWays(stringLength, containerLength int) int {
	if stringLength > containerLength {
		return 0
	}
	if stringLength == containerLength {
		return 1
	}
	// Calculate permutations: containerLength(n) P stringLength(r)
	ways := permutations(containerLength, stringLength)

	return ways
}

func partOne() int {
	for _, record := range records {
		// go through the springs to find groups of #
		// count the number of permutations that the corresponding number in the report has in the set of #

		fmt.Println(record.springSets)
		ways := 0
		if len(record.springSets) == len(record.report) {
			for i, set := range record.springSets {
				if regexp.MustCompile(`^#+$`).MatchString(set) {
					// these ones can be ignored as there are no alternatives - counts as 0
					// e.g. ### (3)
					// fmt.Printf("set %v is entirely made from #, of length %v and the corresponding report number is %v\n", set, len(set), record.report[i])
				} else if len(set) == record.report[i] {
					// this is a combination of ?# same length as number of springs
					// e.g. ???? (4) or #?# (3)
					// it counts as 1 not 0
					fmt.Printf("set %v of length %v can only take report number %v in one way\n", set, len(set), record.report[i])
					ways++
				} else {
					fmt.Printf("set %v is of length %v and the corresponding report number is %v\n", set, len(set), record.report[i])
					ways += countWays(record.report[i], len(set))
				}
			}
		} else {
			// there is going to be multiple springSets in one set of ?#
			fmt.Printf("we have springSets: %v and report: %v\n", record.springSets, record.report)
			// ???
		}
		fmt.Printf(" > Counted %v ways\n\n", ways)
	}
	return 0
}

func partTwo() int {
	return 0
}

func Call(part string, inputFile string) string {
	lines := util.ParseInputIntoLines(inputFile)
	for _, line := range lines {
		s := strings.Split(line, " ")
		r := strings.Split(s[1], ",")
		report := []int{}
		for _, n := range r {
			num, _ := strconv.Atoi(n)
			report = append(report, num)
		}
		re := regexp.MustCompile(`\.+`)
		sprSplit := re.Split(s[0], -1)
		springSets := []string{}
		for _, x := range sprSplit {
			if len(x) > 0 {
				springSets = append(springSets, x)
			}
		}
		records = append(records, Record{springs: s[0], springSets: springSets, report: report})
	}
	var r int
	if part == "1" {
		r = partOne()
	} else {
		r = partTwo()
	}
	return strconv.Itoa(r)
}
