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

func getCount(grid []string) int {
	// work on rows to find reflections
	// need to work down the rows until an axis of symmetry is found
	// if the axis is before half-way: (e.g. 4 out of 9)
	//  ALL rows above must have reflections until row 1 (1-4 must correspond to 5-8)
	// if the axis is beyond half-way: (e.g. 5 out of 9)
	//  ALL rows below must have reflections until the last row (5-9 must correspond with 4-1)
	// then return the number of rows above this point
	for _, row := range grid {
		fmt.Println(row)
	}
	fmt.Println("-------------------")
	for rowNum, row := range grid {
		if rowNum == 0 {
			continue
		}
		if row == grid[rowNum-1] {
			axis := rowNum - 1
			fmt.Printf("possible axis found at row %v of %v\n", axis, len(grid))
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
				fmt.Printf(" > good! returning %v\n", axis)
				return axis + 1
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
		count = getCount(rowsVersion)
		if count > 0 {
			total += (100 * count)
		} else {
			colsVersion := rotate(rowsVersion)
			fmt.Println(" > try rotated")
			count = getCount(colsVersion)
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
	return 0
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
