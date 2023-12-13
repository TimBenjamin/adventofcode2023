package day_13

import (
	"adventofcode2023/util"
	"fmt"
	"strconv"
	"strings"
)

var grids [][]string

func rotate(grid []string) []string {
	rot := []string{}
	tmp := make([][]string, len(grid[0]))
	for _, line := range grid {
		s := strings.Split(line, "")
		for i, c := range s {
			if len(tmp[i]) == 0 {
				tmp[i] = []string{}
			}
			tmp[i] = append(tmp[i], c)
		}
	}
	for _, t := range tmp {
		rot = append(rot, strings.Join(t, ""))
	}
	return rot
}

func getCountPartOne(grid []string) int {
	// work on rows to find reflections
	// need to work down the rows until an axis of symmetry is found
	// if the axis is before half-way: (e.g. 4 out of 9)
	//  ALL rows above must have reflections until row 1 (1-4 must correspond to 5-8)
	// if the axis is beyond half-way: (e.g. 5 out of 9)
	//  ALL rows below must have reflections until the last row (5-9 must correspond with 4-1)
	// then return the number of rows above this point
	// for _, row := range grid {
	// 	fmt.Println(row)
	// }
	// fmt.Println("-------------------")
	for rowNum, row := range grid {
		if rowNum == 0 {
			continue
		}
		if row == grid[rowNum-1] {
			axis := rowNum - 1
			fmt.Printf("possible axis found at row %v of %v (%v/%v)\n", axis, len(grid), grid[rowNum], grid[axis])
			good := true
			for i := 0; i < len(grid)/2; i++ {
				fmt.Printf("matching rows: %v, %v\n", axis-i, axis+i+1)
				if grid[axis-i] != grid[axis+i+1] {
					good = false
					break
				}
				if axis-i == 0 || axis+i+1 >= len(grid)-1 {
					break
				}
			}
			if good {
				fmt.Printf(" > good! returning %v\n", axis+1)
				fmt.Println(" > matching grid:")
				for _, row := range grid {
					fmt.Println(row)
				}
				fmt.Println("-------------------")
				return rowNum
			}
		}
	}
	fmt.Println(" - no axis found")
	return 0
}

// swap the character at y,x
func generateTestGrid(grid []string, changeY int, changeX int) []string {
	testGrid := []string{}
	for y, row := range grid {
		if y == changeY {
			rowSplit := strings.Split(row, "")
			if rowSplit[changeX] == "." {
				rowSplit[changeX] = "#"
			} else {
				rowSplit[changeX] = "."
			}
			testGrid = append(testGrid, strings.Join(rowSplit, ""))
		} else {
			testGrid = append(testGrid, row)
		}
	}
	return testGrid
}

func getCountPartTwo(grid []string) int {
	// Same algorithm as part 1, but with added fun of swapping characters one by one until the reflection is found
	// Must ignore the original reflection...
	var count int
	originalReflectionCount := getCountPartOne(grid)
	originalReflectionRowsOrColumns := "rows"
	if originalReflectionCount > 0 {
		originalReflectionCount = 100 * originalReflectionCount
	} else {
		colsVersion := rotate(grid)
		originalReflectionCount = getCountPartOne(colsVersion)
		originalReflectionRowsOrColumns = "cols"
	}

	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			testGrid := generateTestGrid(grid, y, x)
			fmt.Printf("new grid after changing %v,%v:\n", y, x)
			for _, row := range testGrid {
				fmt.Println(row)
			}
			fmt.Println("=====")
			f := -1
			if originalReflectionRowsOrColumns == "rows" {
				f = originalReflectionCount
			}
			count = getCountPartTwoHelper(testGrid, f)
			if count > 0 {
				if count*100 == originalReflectionCount {
					// this is the original one, skip
					fmt.Println(" - but this is the original one, skip")
					continue
				}
				return count
			} else {
				testGridColsVersion := rotate(testGrid)
				count = getCountPartTwoHelper(testGridColsVersion, originalReflectionCount)
				if count > 0 {
					if count == originalReflectionCount {
						// original, skip
						fmt.Println(" - but this is the original one, skip")
						continue
					}
					return count
				}
			}
		}
	}
	return 0
}

// identical to getCountPartOne but want to skip "skipRow"
func getCountPartTwoHelper(grid []string, skipRow int) int {
	for rowNum, row := range grid {
		if rowNum == 0 {
			continue
		}
		if row == grid[rowNum-1] {
			if rowNum == skipRow {
				continue
			}
			axis := rowNum - 1
			fmt.Printf("possible axis found at row %v of %v (%v/%v)\n", axis, len(grid), grid[rowNum], grid[axis])
			good := true
			for i := 0; i < len(grid)/2; i++ {
				fmt.Printf("matching rows: %v, %v\n", axis-i, axis+i+1)
				if grid[axis-i] != grid[axis+i+1] {
					good = false
					break
				}
				if axis-i == 0 || axis+i+1 >= len(grid)-1 {
					break
				}
			}
			if good {
				fmt.Printf(" > good! returning %v\n", axis+1)
				fmt.Println(" > matching grid:")
				for _, row := range grid {
					fmt.Println(row)
				}
				fmt.Println("-------------------")
				return rowNum
			}
		}
	}
	fmt.Println(" - no axis found")
	return 0
}

func partOne() int {
	total := 0
	var count int
	for i, rowsVersion := range grids {
		fmt.Printf("Working on grid %v of %v\n", i+1, len(grids))
		count = getCountPartOne(rowsVersion)
		if count > 0 {
			total += (100 * count)
		} else {
			colsVersion := rotate(rowsVersion)
			fmt.Println(" > try rotated")
			count = getCountPartOne(colsVersion)
			if count > 0 {
				total += count
			} else {
				panic("No reflections found")
			}
		}
		fmt.Println()
	}
	return total
}

func partTwo() int {
	// get reflections after changing every character from . to # and vice-versa in turn (i.e. a smudge).
	// do not return original reflections (in fact they should not "work" once a "smudge" is in place?)
	total := 0
	var count int
	for i, rowsVersion := range grids {
		fmt.Printf("Working on grid %v of %v which looks like:\n", i+1, len(grids))
		for _, row := range rowsVersion {
			fmt.Println(row)
		}
		fmt.Println("-=-=-=-=-=-=-=-=-")
		count = getCountPartTwo(rowsVersion)
		if count > 0 {
			total += (100 * count)
		} else {
			colsVersion := rotate(rowsVersion)
			fmt.Println(" > try rotated")
			count = getCountPartTwo(colsVersion)
			if count > 0 {
				total += count
			} else {
				fmt.Printf("Problem on grid %v of %v\n", i+1, len(grids))
				for _, row := range rowsVersion {
					fmt.Println(row)
				}
				panic("No reflections found")
			}
		}
		fmt.Println()
	}
	return total
}

func Call(part string, inputFile string) string {
	in := util.ParseInputIntoLines(inputFile)
	grids = [][]string{}
	grid := []string{}
	for _, line := range in {
		if len(line) > 0 {
			grid = append(grid, line)
		} else {
			grids = append(grids, grid)
			grid = []string{}
		}
	}
	grids = append(grids, grid)
	var r int
	if part == "1" {
		r = partOne()
	} else {
		r = partTwo()
	}
	return strconv.Itoa(r)
}
