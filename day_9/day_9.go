package day_9

import (
	"adventofcode2023/util"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

var input []string

func getDiffs(nums []int) []int {
	diffs := []int{}
	for i, n := range nums {
		if i == 0 {
			continue
		}
		diffs = append(diffs, n-nums[i-1])
	}
	return diffs
}

func partOne() int {
	total := 0
	for _, line := range input {
		nums := []int{}
		for _, c := range strings.Split(line, " ") {
			n, _ := strconv.Atoi(c)
			nums = append(nums, n)
		}
		red := nums
		fmt.Println(red)
		nextDiff := 0 // the sum of the last digits in the reduced list on each iteration; this is the same as the diff to the next number in the original list.
		for {
			red = getDiffs(red)
			fmt.Println(red)
			allZero := true
			for _, r := range red {
				if r != 0 {
					allZero = false
					break
				}
			}
			if allZero {
				break
			}
			nextDiff += red[len(red)-1]
		}
		next := nums[len(nums)-1] + nextDiff
		fmt.Printf("Next in sequence is %v\n", next)
		fmt.Println()
		total += next
	}
	return total
}

func partTwo() int {
	total := 0
	for _, line := range input {
		nums := []int{}
		for _, c := range strings.Split(line, " ") {
			n, _ := strconv.Atoi(c)
			nums = append(nums, n)
		}
		red := nums
		// we can just reverse the sequence from part 1 to "go backwards"
		slices.Reverse(red)
		fmt.Println(red)
		nextDiff := 0
		for {
			red = getDiffs(red)
			fmt.Println(red)
			allZero := true
			for _, r := range red {
				if r != 0 {
					allZero = false
					break
				}
			}
			if allZero {
				break
			}
			nextDiff += red[len(red)-1]
		}
		next := nums[len(nums)-1] + nextDiff
		fmt.Printf("Next in sequence is %v\n", next)
		fmt.Println()
		total += next
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
