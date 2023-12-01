package util

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func ParseInputIntoLines(inputFile string) (input []string) {
	f, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var line string
	for scanner.Scan() {
		line = scanner.Text()
		input = append(input, line)
	}

	return input
}

func ParseSingleLineInput(inputFile string) string {
	f, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	scanner.Scan()
	return scanner.Text()
}

func ParseSingleLineInputIntoInts(inputFile string) []int {
	s := ParseSingleLineInput(inputFile)
	words := strings.Fields(s)
	nums := []int{}
	for _, w := range words {
		i, _ := strconv.Atoi(w)
		nums = append(nums, i)
	}
	return nums
}
