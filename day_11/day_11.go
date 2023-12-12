package day_11

import (
	"adventofcode2023/util"
	"fmt"
	"math"
	"strconv"
)

var grid [][]string

func expand(grid [][]string) [][]string {
	expandedLineGrid := [][]string{}
	// first double empty rows
	for _, row := range grid {
		expandedLineGrid = append(expandedLineGrid, row)
		isEmpty := true
		for _, c := range row {
			if c != "." {
				isEmpty = false
				break
			}
		}
		if isEmpty {
			emptyRow := []string{}
			for i := 0; i < len(row); i++ {
				emptyRow = append(emptyRow, "M")
			}
			expandedLineGrid = append(expandedLineGrid, emptyRow)
		}
	}
	// now columns...
	emptyColumns := make([]int, len(grid))
	for _, row := range expandedLineGrid {
		for colNum, c := range row {
			if c == "#" {
				emptyColumns[colNum] = 1
			}
		}
	}
	expandedGrid := [][]string{}
	for _, row := range expandedLineGrid {
		newRow := []string{}
		for colNum, c := range row {
			newRow = append(newRow, c)
			if emptyColumns[colNum] == 0 {
				newRow = append(newRow, "M")
			}
		}
		expandedGrid = append(expandedGrid, newRow)
	}
	return expandedGrid
}

func getGalaxyLocations(grid [][]string) [][]int {
	locations := [][]int{}
	for y, row := range grid {
		for x, c := range row {
			if c == "#" {
				locations = append(locations, []int{y, x})
			}
		}
	}
	return locations
}

func calculateManhattanDistance(a, b []int) int {
	if len(a) != len(b) || len(a) != 2 {
		panic("Invalid input: Both points must have exactly two coordinates.")
	}

	return int(math.Abs(float64(a[0]-b[0])) + math.Abs(float64(a[1]-b[1])))
}

func partOne() int {
	grid = expand(grid)
	galaxies := getGalaxyLocations(grid)
	// need the manhattan distance between each pair
	// can do them all then halve the result (combinations not permutations)
	distanceTotal := 0
	for i, galaxyA := range galaxies {
		for j, galaxyB := range galaxies {
			if i == j {
				continue
			}
			md := calculateManhattanDistance(galaxyA, galaxyB)
			// fmt.Printf("distance from galaxy %v (%v) to galaxy %v (%v) is %v\n", i+1, galaxyA, j+1, galaxyB, md)
			distanceTotal += md
		}
	}
	return distanceTotal / 2
}

func partTwo() int {
	grid = expand(grid)
	for _, row := range grid {
		fmt.Println(row)
	}
	galaxies := getGalaxyLocations(grid)
	// get the row and col numbers that have M at the start
	expandedRows := []int{}
	expandedCols := []int{}
	for y, row := range grid {
		if row[0] == "M" {
			expandedRows = append(expandedRows, y)
		}
	}
	for x, c := range grid[0] {
		if c == "M" {
			expandedCols = append(expandedCols, x)
		}
	}
	fmt.Printf("expanded rows are: %v\n", expandedRows)
	fmt.Printf("expanded cols are: %v\n", expandedCols)
	// get the manhattan distance as before
	// for every row and col between A and B that has an M, add 1M to the distance
	distanceTotal := 0
	for i, galaxyA := range galaxies {
		for j, galaxyB := range galaxies {
			if i == j {
				continue
			}
			md := calculateManhattanDistance(galaxyA, galaxyB)
			for _, y := range expandedRows {
				if (galaxyA[0] < y && galaxyB[0] > y) || (galaxyB[0] < y && galaxyA[0] > y) {
					md += 1000000 - 2
				}
			}
			for _, x := range expandedCols {
				if (galaxyA[1] < x && galaxyB[1] > x) || (galaxyB[1] < x && galaxyA[1] > x) {
					md += 1000000 - 2
				}
			}
			distanceTotal += md
		}
	}
	return distanceTotal / 2
}

func Call(part string, inputFile string) string {
	grid = util.ParseInputIntoStringGrid(inputFile)
	var r int
	if part == "1" {
		r = partOne()
	} else {
		r = partTwo()
	}
	return strconv.Itoa(r)
}
