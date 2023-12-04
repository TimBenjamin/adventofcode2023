package day_4

import (
	"adventofcode2023/util"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var input []string

func partOne() int {
	total := 0
	for _, line := range input {
		nums := strings.Split(line, ": ")
		cards := strings.Split(nums[1], " | ")
		re := regexp.MustCompile(`\s+`)
		winningCardNums := re.Split(strings.TrimSpace(cards[0]), -1)
		winningCard := make(map[string]struct{})
		for _, num := range winningCardNums {
			winningCard[num] = struct{}{}
		}
		myNumbers := re.Split(strings.TrimSpace(cards[1]), -1)
		numberOfMatches := 0
		for _, num := range myNumbers {
			if _, ok := winningCard[num]; ok {
				numberOfMatches++
			}
		}
		if numberOfMatches > 0 {
			total += 1 << (numberOfMatches - 1)
		}
	}
	return total
}

func partTwo() int {
	total := 0
	cardCopies := make(map[int]int)
	for i := 1; i <= len(input); i++ {
		cardCopies[i] = 1
	}
	fmt.Println(cardCopies)
	fmt.Println()
	for _, line := range input {
		nums := strings.Split(line, ":")
		re := regexp.MustCompile(`\s+`)
		cardNumber, _ := strconv.Atoi(re.Split(nums[0], -1)[1])
		cards := strings.Split(nums[1], " | ")
		winningCardNums := re.Split(strings.TrimSpace(cards[0]), -1)
		winningCard := make(map[string]struct{})
		for _, num := range winningCardNums {
			winningCard[num] = struct{}{}
		}
		myNumbers := re.Split(strings.TrimSpace(cards[1]), -1)
		numberOfMatches := 0
		for _, num := range myNumbers {
			if _, ok := winningCard[num]; ok {
				numberOfMatches++
			}
		}
		fmt.Printf("card %v has %v matches\n", cardNumber, numberOfMatches)

		for i := 1; i <= (numberOfMatches); i++ {
			copies := cardCopies[cardNumber]
			fmt.Printf("card %v gets %v copies\n", (cardNumber + i), copies)
			cardCopies[cardNumber+i] += copies
		}
		fmt.Println(cardCopies)
		fmt.Println()
	}
	fmt.Println(cardCopies)
	for _, v := range cardCopies {
		// fmt.Printf("card: %v / copies: %v\n", k, v)
		total += v
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
