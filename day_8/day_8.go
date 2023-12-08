package day_8

import (
	"adventofcode2023/util"
	"strconv"
	"strings"
)

var input []string

type Node struct {
	L string
	R string
}

var nodeMap map[string]Node
var instructions []string

func partOne() int {
	instructionsIndex := 0
	currentNode := "AAA"
	steps := 0
	for {
		steps++
		if instructions[instructionsIndex] == "L" {
			currentNode = nodeMap[currentNode].L
		} else {
			currentNode = nodeMap[currentNode].R
		}
		if currentNode == "ZZZ" {
			break
		}
		instructionsIndex++
		if instructionsIndex == len(instructions) {
			instructionsIndex = 0
		}
	}
	return steps
}

func getSteps(currentNode string) int {
	instructionsIndex := 0
	steps := 0
	for {
		steps++
		if instructions[instructionsIndex] == "L" {
			currentNode = nodeMap[currentNode].L
		} else {
			currentNode = nodeMap[currentNode].R
		}
		if currentNode[2] == 'Z' {
			return steps
		}
		instructionsIndex++
		if instructionsIndex == len(instructions) {
			instructionsIndex = 0
		}
	}
}

func findGCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
func findLCM(a, b int) int {
	return (a * b) / findGCD(a, b)
}

func partTwo() int {
	currentNodes := []string{}
	for k := range nodeMap {
		if k[2] == 'A' {
			currentNodes = append(currentNodes, k)
		}
	}

	// find out how many steps it takes for each of the starting nodes to end in a node that ends in Z
	// then find the least common multiple
	steps := []int{}
	for _, currentNode := range currentNodes {
		steps = append(steps, getSteps(currentNode))
	}

	product := 1
	for _, s := range steps {
		product *= s
	}

	lcm := steps[0]
	for _, num := range steps[1:] {
		lcm = findLCM(lcm, num)
	}

	return lcm
}

func Call(part string, inputFile string) string {
	input = util.ParseInputIntoLines(inputFile)
	instructions = strings.Split(input[0], "")
	nodeMap = map[string]Node{}
	for _, row := range input[2:] {
		parts := strings.Split(row, " = ")
		locations := strings.Split(parts[1][1:9], ", ")
		nodeMap[parts[0]] = Node{L: locations[0], R: locations[1]}
	}
	var r int
	if part == "1" {
		r = partOne()
	} else {
		r = partTwo()
	}
	return strconv.Itoa(r)
}
