package day_12

import (
	"adventofcode2023/util"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Record struct {
	springs string
	spec    []int
}

var records []Record

func check(pattern string, spec []int) bool {
	pattern = regexp.MustCompile(`(^\.+|\.+$)`).ReplaceAllString(pattern, "")
	pattern = regexp.MustCompile(`\.+`).ReplaceAllString(pattern, ".")
	t := []string{}
	for _, s := range spec {
		x := ""
		for i := 0; i < s; i++ {
			x += "#"
		}
		t = append(t, x)
	}
	return pattern == strings.Join(t, ".")
}

func generateVariations(input string, spec []int) []string {
	// Find the index of the first "?" in the string
	index := strings.Index(input, "?")

	// If no "?" is found, return the original string
	if index == -1 {
		// If the string doesn't meet the basic spec in terms of groups, return empty
		s := regexp.MustCompile(`(^\.+|\.+$)`).ReplaceAllString(input, "")
		sp := regexp.MustCompile(`\.+`).Split(s, -1)
		if len(sp) != len(spec) {
			// fmt.Printf("%v (%v) does not meet spec %v\n", input, sp, specLength)
			return []string{}
		}
		return []string{input}
	}

	// Generate variations by replacing "?" with "." and "#"
	var variations []string
	variations = append(variations, generateVariations(input[:index]+"."+input[index+1:], spec)...)
	variations = append(variations, generateVariations(input[:index]+"#"+input[index+1:], spec)...)

	return variations
}

func partOne() int {
	total := 0
	for _, record := range records {
		memo := map[string]bool{}

		fmt.Printf("Springs: %v / Spec: %v\n", record.springs, record.spec)

		variations := generateVariations(record.springs, record.spec)
		fmt.Printf("size of variations: %v\n", len(variations))

		ways := 0
		for _, v := range variations {
			if ok := memo[v]; ok {
				ways++
			} else if check(v, record.spec) {
				ways++
				memo[v] = true
			} else {
				memo[v] = false
			}
		}
		total += ways
		fmt.Printf(" > %v ways\n", ways)
		fmt.Println()
	}
	return total
}

func getUnfoldedRecord(record Record) Record {
	// generate the "unfolded" version
	s := []string{}
	spec := []int{}
	for i := 0; i < 5; i++ {
		s = append(s, record.springs)
		spec = append(spec, record.spec...)
	}
	springs := strings.Join(s, "?")
	return Record{springs: springs, spec: spec}
}

func partTwo() int {
	total := 0
	for _, record := range records {
		memo := map[string]bool{}

		unfoldedRecord := getUnfoldedRecord(record)
		fmt.Printf("springs: %v / spec: %v\n", unfoldedRecord.springs, unfoldedRecord.spec)

		variations := generateVariations(unfoldedRecord.springs, unfoldedRecord.spec)
		fmt.Printf("size of variations: %v\n", len(variations))

		ways := 0
		for _, v := range variations {
			if ok := memo[v]; ok {
				ways++
			} else if check(v, unfoldedRecord.spec) {
				ways++
				memo[v] = true
			} else {
				memo[v] = false
			}
		}
		total += ways
		fmt.Printf(" > %v ways\n", ways)
		fmt.Println()
	}
	return 0
}

func Call(part string, inputFile string) string {
	lines := util.ParseInputIntoLines(inputFile)
	for _, line := range lines {
		s := strings.Split(line, " ")
		r := strings.Split(s[1], ",")
		springs := regexp.MustCompile(`(^\.+|\.+$)`).ReplaceAllString(s[0], "")
		spec := []int{} // the list of numbers
		for _, n := range r {
			num, _ := strconv.Atoi(n)
			spec = append(spec, num)
		}
		records = append(records, Record{
			springs: springs, spec: spec,
		})
	}
	var r int
	if part == "1" {
		r = partOne()
	} else {
		r = partTwo()
	}
	return strconv.Itoa(r)
}
