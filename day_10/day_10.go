package day_10

import (
	"adventofcode2023/util"
	"fmt"
	"strconv"
	"strings"
)

var input []string
var grid [][]string
var startCoords []int
var currentDirection string

func getNewCoords(currentCoords []int) []int {
	// NB, coords are always y, x

	// this function moves one step based on the current shape and current direction
	// e.g. | and north => move one space up

	// if we are on the start shape, determine:
	// if space above is | then move up and set dir to north
	// if space above is F then move up and set dir to east
	// if space above is 7 then move up and set dir to west
	// if space left is - then move left and set dir to west
	// etc

	var apply []int // the transform to apply based on the shape and current direction

	shape := grid[currentCoords[0]][currentCoords[1]] // the shape at our current location
	// fmt.Printf("Shape at coords %v is %v\n", currentCoords, shape)

	if shape == "S" {
		// special case for the start

		// look north, then east, then south. No need to look west.
		if grid[currentCoords[0]][currentCoords[1]+1] == "-" {
			apply = []int{0, 1}
			currentDirection = "east"
		} else if grid[currentCoords[0]][currentCoords[1]+1] == "J" {
			apply = []int{0, 1}
			currentDirection = "east"
		} else if grid[currentCoords[0]][currentCoords[1]+1] == "7" {
			apply = []int{0, 1}
			currentDirection = "east"
		} else if grid[currentCoords[0]+1][currentCoords[1]] == "|" {
			apply = []int{1, 0}
			currentDirection = "south"
		} else if grid[currentCoords[0]+1][currentCoords[1]] == "J" {
			apply = []int{1, 0}
			currentDirection = "south"
		} else if grid[currentCoords[0]+1][currentCoords[1]] == "L" {
			apply = []int{1, 0}
			currentDirection = "south"
		} else if grid[currentCoords[0]-1][currentCoords[1]] == "|" {
			apply = []int{-1, 0}
			currentDirection = "north"
		} else if grid[currentCoords[0]-1][currentCoords[1]] == "F" {
			apply = []int{-1, 0}
			currentDirection = "north"
		} else if grid[currentCoords[0]-1][currentCoords[1]] == "7" {
			apply = []int{-1, 0}
			currentDirection = "north"
		} else {
			fmt.Println("ERROR: cannot move from starting location")
		}
	} else {
		switch shape {
		case "|":
			{
				if currentDirection == "north" {
					apply = []int{-1, 0}
				} else if currentDirection == "south" {
					apply = []int{1, 0}
				} else {
					fmt.Printf("ERROR: cannot have a vertical pipe without N/S direction")
				}
			}
		case "-":
			{
				if currentDirection == "east" {
					apply = []int{0, 1}
				} else if currentDirection == "west" {
					apply = []int{0, -1}
				} else {
					fmt.Printf("ERROR: cannot have a horizontal pipe without E/W direction")
				}
			}
		case "L":
			{
				if currentDirection == "south" {
					apply = []int{0, 1}
					currentDirection = "east"
				}
				if currentDirection == "west" {
					currentDirection = "north"
					apply = []int{-1, 0}
				}
			}
		case "J":
			{
				if currentDirection == "east" {
					currentDirection = "north"
					apply = []int{-1, 0}
				}
				if currentDirection == "south" {
					currentDirection = "west"
					apply = []int{0, -1}
				}
			}
		case "7":
			{
				if currentDirection == "east" {
					currentDirection = "south"
					apply = []int{1, 0}
				}
				if currentDirection == "north" {
					currentDirection = "west"
					apply = []int{0, -1}
				}
			}
		case "F":
			{
				if currentDirection == "north" {
					currentDirection = "east"
					apply = []int{0, 1}
				}
				if currentDirection == "west" {
					currentDirection = "south"
					apply = []int{1, 0}
				}
			}
		case ".":
			{
				fmt.Println("ERROR: this tile contains no pipe")
			}
		}
	}
	// fmt.Printf("Found shape %v, applying transform %v, current direction is %v\n", shape, apply, currentDirection)
	currentCoords[0] += apply[0]
	currentCoords[1] += apply[1]
	return currentCoords
}

func partOne() int {
	startX := startCoords[1]
	startY := startCoords[0]
	currentCoords := startCoords
	steps := [][]int{}
	for {
		steps = append(steps, []int{currentCoords[0], currentCoords[1]})
		currentCoords = getNewCoords(currentCoords)
		if len(steps) > 1 && currentCoords[0] == startY && currentCoords[1] == startX {
			fmt.Printf("Finished loop after %v steps\n", len(steps))
			break
		}
	}

	// max distance is half (the length of steps)
	return len(steps) / 2
}

func stepsContains(coord []int, steps [][]int) bool {
	for _, step := range steps {
		if step[0] == coord[0] && step[1] == coord[1] {
			return true
		}
	}
	return false
}

func partTwo() int {
	startX := startCoords[1]
	startY := startCoords[0]
	currentCoords := startCoords
	steps := [][]int{}
	for {
		steps = append(steps, []int{currentCoords[0], currentCoords[1]})
		currentCoords = getNewCoords(currentCoords)
		if len(steps) > 1 && currentCoords[0] == startY && currentCoords[1] == startX {
			fmt.Printf("Finished loop after %v steps\n", len(steps))
			break
		}
	}

	// we now have a list of all the coords we have passed through
	fmt.Println(steps)

	// any coord in the grid that is NOT on an edge and IS a "." is potentially enclosed
	possiblyEnclosed := [][]int{}
	for y, row := range grid {
		// if y == 0 || y == len(grid)-1 {
		// 	continue // ignore edges
		// }
		for x, _ := range row {
			// if x == 0 || x == len(row)-1 {
			// 	continue // ignore edges
			// }
			// if !stepsContains([]int{y, x}, steps) {
			possiblyEnclosed = append(possiblyEnclosed, []int{y, x})
			// }
		}
	}
	fmt.Printf("possibly enclosed locations: %v\n", possiblyEnclosed)

	// test
	// possiblyEnclosed = [][]int{}
	// possiblyEnclosed = append(possiblyEnclosed, []int{4, 4})
	// test3.txt should give 4
	// test4.txt should give 8
	// test4 start shape is F
	grid[startY][startX] = "F"
	// test5.txt should give 10
	// test5 start shape is 7
	// grid[startY][startX] = "7"
	// test 5: [6,13] and [6,14] should be inside
	// [7,1] should NOT be inside

	// TODO: we need to know the shape of the starting position

	// From each of these points, "cast a ray" East
	// if the ray crosses the loop (i.e. the coordinate is in `steps`) an odd number of times then it is inside
	// however I am not sure how to define "crossing" when it comes to bending steps F, J, 7, L
	numberInside := 0

	for y, row := range grid {
		for x, shape := range row {
			testCoord := []int{y, x}
			if stepsContains(testCoord, steps) {
				continue
			}
			fmt.Printf("Test coord: %v\n", testCoord)

			// go east from this coord and look for crossings
			for i := x; i < len(row); i++ {
				coord := []int{y, i}
				fmt.Printf("  found shape %v\n", shape)
				crossings := 0
				dir := ""
				if stepsContains(coord, steps) {
					fmt.Printf("   it is on the loop\n")
					if shape == "|" {
						crossings++
						dir = ""
					}
					if shape == "J" || shape == "L" {
						if dir == "down" {
							crossings++
							dir = ""
						} else {
							dir = "up"
						}
					} else if shape == "7" || shape == "F" {
						if dir == "up" {
							crossings++
							dir = ""
						} else {
							dir = "down"
						}
					}
					fmt.Printf(" current dir is: %v\n", dir)
				}
				fmt.Printf(" > crossings: %v\n", crossings)
				if crossings == 0 || crossings%2 == 0 {
					fmt.Println("  > even number of E crossings, point is not inside")
				} else {
					fmt.Printf(" >> point %v is inside!\n", coord)
					numberInside++
				}
			}
		}

		// n := 0
		// s := 0

		// for i := coord[1]; i < len(grid[0])-1; i++ {
		// 	shape := grid[coord[0]][i]
		// 	shapeCoord := []int{coord[0], i}
		// 	if shapeCoord[0] == coord[0] && shapeCoord[1] == coord[1] {
		// 		continue
		// 	}
		// 	fmt.Printf("  found shape %v\n", shape)
		// 	if stepsContains(shapeCoord, steps) {
		// 		if shape == "|" {
		// 			n++
		// 			s++
		// 		}
		// 		if shape == "J" || shape == "L" {
		// 			n++
		// 		}
		// 		if shape == "F" || shape == "7" {
		// 			s++
		// 		}
		// 	}
		// }

		// u := s
		// if n < s {
		// 	u = n
		// }
		// if u > 0 && u%2 == 1 {
		// 	fmt.Printf(" >> point %v is inside!\n", coord)
		// 	numberInside++
		// } else {
		// 	fmt.Printf(" >> point %v is outside\n", coord)
		// }

	}

	return numberInside
}

func Call(part string, inputFile string) string {
	input = util.ParseInputIntoLines(inputFile)
	for lineNum, line := range input {
		grid = append(grid, strings.Split(line, ""))
		if strings.Contains(line, "S") {
			startCoords = []int{lineNum, strings.Index(line, "S")}
		}
	}
	var r int
	if part == "1" {
		r = partOne()
	} else {
		r = partTwo()
	}
	return strconv.Itoa(r)
}
