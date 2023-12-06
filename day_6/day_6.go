package day_6

import (
	"adventofcode2023/util"
	"fmt"
	"regexp"
	"strconv"
)

var input []string

type Race struct {
	time     int
	distance int
}

var races []Race

func partOne() int {
	times := PartOneProcessInput(input[0])
	distances := PartOneProcessInput(input[1])
	for i := 0; i < len(times); i++ {
		races = append(races, Race{time: times[i], distance: distances[i]})
	}
	recordTotals := []int{}
	for _, race := range races {
		fmt.Printf("doing race with time: %v and distance: %v\n", race.time, race.distance)
		numRecords := 0
		for hold := 0; hold <= race.time; hold++ {
			distance := hold * (race.time - hold)
			if distance > race.distance {
				numRecords++
			}
		}
		fmt.Printf(" => there are %v ways to beat the record\n\n", numRecords)
		recordTotals = append(recordTotals, numRecords)
	}
	total := 1
	for _, w := range recordTotals {
		total *= w
	}
	return total
}

func partTwo() int {
	race := Race{time: PartTwoProcessInput(input[0]), distance: PartTwoProcessInput(input[1])}
	numRecords := 0
	for hold := 0; hold <= race.time; hold++ {
		distance := hold * (race.time - hold)
		if distance > race.distance {
			numRecords++
		}
	}
	return numRecords
}

func PartOneProcessInput(line string) []int {
	digits := []int{}
	digitString := regexp.MustCompile(`\D+:\s+`).ReplaceAllString(line, "")
	digitStringSplit := regexp.MustCompile(`\s+`).Split(digitString, -1)
	for _, s := range digitStringSplit {
		d, _ := strconv.Atoi(s)
		digits = append(digits, d)
	}
	return digits
}

func PartTwoProcessInput(line string) int {
	digitString := regexp.MustCompile(`\D+:\s+`).ReplaceAllString(line, "")
	digitStringCombined := regexp.MustCompile(`\s+`).ReplaceAllString(digitString, "")
	j, _ := strconv.Atoi(digitStringCombined)
	return j
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
