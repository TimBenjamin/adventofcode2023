package day_1

import (
	"adventofcode2023/util"
	"strconv"
	"strings"
)

var input []string

func partOne() int {
	total := 0
	for _, line := range input {
		lineSlice := strings.Split(line, "")
		var nums []int
		for _, c := range lineSlice {
			num, _ := strconv.Atoi(c)
			if num > 0 {
				nums = append(nums, num)
			}
		}
		total += (nums[0] * 10) + nums[len(nums)-1]
	}
	return total
}

func partTwo() int {
	total := 0
	for _, line := range input {
		var nums []int
		num := 0
		rest := line
		for len(rest) > 0 {
			num, _ = strconv.Atoi(strings.Split(rest, "")[0])
			if num > 0 {
				nums = append(nums, num)
			} else {
				numbers := map[string]int{
					"one":   1,
					"two":   2,
					"three": 3,
					"four":  4,
					"five":  5,
					"six":   6,
					"seven": 7,
					"eight": 8,
					"nine":  9,
				}
				for stringNumber, intNumber := range numbers {
					if strings.Index(rest, stringNumber) == 0 {
						nums = append(nums, intNumber)
					}
				}
			}
			rest = rest[1:]
		}
		total += ((nums[0] * 10) + nums[len(nums)-1])
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
