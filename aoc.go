package main

import (
	"adventofcode2023/day_1"
	"adventofcode2023/day_2"
	"adventofcode2023/day_3"
	"adventofcode2023/day_4"
	"fmt"
	"os"
)

// aoc.go <day> <part>
func main() {

	// these exported functions must all have the same return type!
	// therefore any solutions that are ints will be converted to strings
	days := map[string]func(part string, input string) (result string){
		"day_1": day_1.Call,
		"day_2": day_2.Call,
		"day_3": day_3.Call,
		"day_4": day_4.Call,
	}

	var day string
	var part string
	var input string
	if len(os.Args) != 4 {
		fmt.Println("Incorrect number of args: aoc.go <day> <part> <path/to/input.txt>")
	} else {
		day = os.Args[1]
		part = os.Args[2]
		input = os.Args[3]
		f := "day_" + day
		result := days[f](part, input)
		fmt.Println("result:", result)
	}
}
