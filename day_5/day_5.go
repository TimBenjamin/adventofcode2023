package day_5

import (
	"adventofcode2023/util"
	"sort"
	"strconv"
	"strings"
)

var input []string
var seeds []int
var maps map[string][][]int
var chain []string

func parseInput() {
	maps = map[string][][]int{
		"seedToSoil":            {},
		"soilToFertilizer":      {},
		"fertilizerToWater":     {},
		"waterToLight":          {},
		"lightToTemperature":    {},
		"temperatureToHumidity": {},
		"humidityToLocation":    {},
	}
	chain = []string{"seedToSoil", "soilToFertilizer", "fertilizerToWater", "waterToLight", "lightToTemperature", "temperatureToHumidity", "humidityToLocation"}
	var currentMap string
	for _, line := range input {

		if len(line) == 0 {
			continue
		}

		if strings.Contains(line, "seeds: ") {
			for _, s := range strings.Split(line[7:], " ") {
				seed, _ := strconv.Atoi(s)
				seeds = append(seeds, seed)
			}
			continue
		}

		if strings.Contains(line, "seed-to-soil") {
			currentMap = "seedToSoil"
			continue
		} else if strings.Contains(line, "soil-to-fertilizer") {
			currentMap = "soilToFertilizer"
			continue
		} else if strings.Contains(line, "fertilizer-to-water") {
			currentMap = "fertilizerToWater"
			continue
		} else if strings.Contains(line, "water-to-light") {
			currentMap = "waterToLight"
			continue
		} else if strings.Contains(line, "light-to-temperature") {
			currentMap = "lightToTemperature"
			continue
		} else if strings.Contains(line, "temperature-to-humidity") {
			currentMap = "temperatureToHumidity"
			continue
		} else if strings.Contains(line, "humidity-to-location") {
			currentMap = "humidityToLocation"
			continue
		}

		numberRange := []int{}
		for _, s := range strings.Split(line, " ") {
			n, _ := strconv.Atoi(s)
			numberRange = append(numberRange, n)
		}
		maps[currentMap] = append(maps[currentMap], numberRange)
	}

	// fmt.Printf("seeds: %v\n", seeds)
	// for k, v := range maps {
	// 	fmt.Printf("map %v: %v\n", k, v)
	// }
}

func getDestinationFromSource(source int, mapName string) int {
	// The maps have [destinationRangeStart, sourceRangeStart, length]
	// Any source numbers that aren't mapped correspond to the same destination number
	// Given: [50,98,2]
	// Source: 99 should return 51
	// Source: 10 should return 10
	// Tests:
	// seed 99 corresponds to soil 51
	// seed 10 corresponds to soil 10
	// seed 53 corresponds to soil 55
	for _, numberRange := range maps[mapName] {
		if source >= numberRange[1] && source < numberRange[1]+numberRange[2] {
			return numberRange[0] + (source - numberRange[1])
		}
	}
	return source
}

func getSourceFromDestination(destination int, mapName string) int {
	// the inverse of the above function
	// Given: [50,98,2]
	// Destination 51 should return 99
	for _, numberRange := range maps[mapName] {
		if destination >= numberRange[0] && destination < numberRange[0]+numberRange[2] {
			return numberRange[1] + (destination - numberRange[0])
		}
	}
	return destination
}

func partOne() int {
	var locations []int
	for _, seed := range seeds {
		result := seed
		for _, mapName := range chain {
			result = getDestinationFromSource(result, mapName)
		}
		locations = append(locations, result)
	}
	sort.Ints(locations)
	return locations[0]
}

func partTwo() int {
	// turn the seeds list into pairs [start, length]
	seedRanges := [][]int{}
	for i := 0; i < len(seeds); i++ {
		seedList := []int{seeds[i], seeds[i+1]}
		seedRanges = append(seedRanges, seedList)
		i++
	}

	lowestLocation := -1
	// for _, seedRange := range seedRanges {
	// 	for seed := seedRange[0]; seed < seedRange[0]+seedRange[1]; seed++ {
	// 		// This is now slow due to the massive ranges but I don't know how to predict when the jumps in the result are
	// 		result := seed
	// 		for _, mapName := range chain {
	// 			result = getDestinationFromSource(result, mapName)
	// 		}
	// 		if lowestLocation == -1 || result < lowestLocation {
	// 			lowestLocation = result
	// 		}
	// 	}
	// }

	// It's faster to reverse the process
	chain = []string{"humidityToLocation", "temperatureToHumidity", "lightToTemperature", "waterToLight", "fertilizerToWater", "soilToFertilizer", "seedToSoil"}
	for {
		lowestLocation++
		result := lowestLocation
		for _, mapName := range chain {
			result = getSourceFromDestination(result, mapName)
		}
		// fmt.Printf("Got seed: %v for location: %v\n", result, lowestLocation)
		// is the result within the ranges of seeds?
		for _, seedRange := range seedRanges {
			if result >= seedRange[0] && result < seedRange[0]+seedRange[1] {
				return lowestLocation
			}
		}
		// fmt.Printf("Current lowest location: %v\n", lowestLocation)
	}
}

func Call(part string, inputFile string) string {
	input = util.ParseInputIntoLines(inputFile)
	parseInput()
	var r int
	if part == "1" {
		r = partOne()
	} else {
		r = partTwo()
	}
	return strconv.Itoa(r)
}
