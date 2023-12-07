package day_7

import (
	"adventofcode2023/util"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

var input []string
var whichPart int

type Hand struct {
	cards []string
	bid   int
	score int
}

func getCardScore(card string) int {
	switch card {
	case "T":
		return 10
	case "J":
		if whichPart == 2 {
			return 1
		}
		return 11
	case "Q":
		return 12
	case "K":
		return 13
	case "A":
		return 14
	}
	score, _ := strconv.Atoi(card)
	return score
}

func handSort(a Hand, b Hand) int {
	if a.score > b.score {
		return 1
	}
	if b.score > a.score {
		return -1
	}
	// tie break!
	for i := 0; i < 5; i++ {
		scoreA := getCardScore(a.cards[i])
		scoreB := getCardScore(b.cards[i])
		if scoreA > scoreB {
			return 1
		}
		if scoreB > scoreA {
			return -1
		}
	}
	fmt.Printf("two hands are identical? %v and %v\n", a.cards, b.cards)
	return 0
}

func getHandScorePartOne(cards []string) int {
	// attempt to classify the hand into one of the types (or none = 0)
	scores := map[string]int{
		"five of a kind":  7,
		"four of a kind":  6,
		"full house":      5,
		"three of a kind": 4,
		"two pair":        3,
		"one pair":        2,
		"high card":       1,
	}
	kind := map[string]int{}
	for _, card := range cards {
		kind[card]++
	}
	pairs := 0
	three := false
	for _, v := range kind {
		if v == 5 {
			return scores["five of a kind"]
		}
		if v == 4 {
			return scores["four of a kind"]
		}
		if v == 3 {
			three = true
		}
		if v == 2 {
			pairs++
		}
	}
	if pairs == 2 {
		return scores["two pair"]
	}
	if pairs == 1 {
		if three {
			return scores["full house"]
		}
		return scores["one pair"]
	}
	if three {
		return scores["three of a kind"]
	}
	if len(kind) == 5 {
		return scores["high card"]
	}
	return 0 // shouldn't be possible...
}

func partOne() int {
	hands := []Hand{}
	for _, line := range input {
		sp := strings.Split(line, " ")
		cards := strings.Split(sp[0], "")
		bid, _ := strconv.Atoi(sp[1])
		hands = append(hands, Hand{cards: cards, bid: bid, score: getHandScorePartOne(cards)})
	}

	slices.SortFunc(hands, handSort)
	total := 0
	for rank, hand := range hands {
		// fmt.Printf("cards: %v / bid: %v / score: %v => rank: %v\n", hand.cards, hand.bid, hand.score, rank+1)
		// fmt.Println()
		total += (rank + 1) * hand.bid
	}
	return total
}

func getHandScorePartTwo(cards []string) int {
	// attempt to classify the hand into one of the types (or none = 0)
	// J can be whatever is going to result in the best outcome
	scores := map[string]int{
		"five of a kind":  7,
		"four of a kind":  6,
		"full house":      5,
		"three of a kind": 4,
		"two pair":        3,
		"one pair":        2,
		"high card":       1,
	}
	kind := map[string]int{}
	for _, card := range cards {
		kind[card]++
	}
	pairs := 0
	three := false
	for _, v := range kind {
		if v == 5 {
			// there is only JJJJJ
			return scores["five of a kind"]
		}
		if kind["J"] == 4 {
			// e.g. J8JJJ
			return scores["five of a kind"]
		}
		if v == 4 {
			if kind["J"] == 1 {
				// e.g. J8888
				return scores["five of a kind"]
			}
			return scores["four of a kind"]
		}
		if v == 3 {
			three = true
		}
		if v == 2 {
			pairs++
		}
	}
	if pairs == 2 {
		if kind["J"] == 2 {
			// two J's count as the other pair
			// e.g. J8QJ8 => 88Q88
			return scores["four of a kind"]
		}
		if kind["J"] == 1 {
			// the spare J makes one pair count as a 3
			// 2233J => 22333
			return scores["full house"]
		}
		// if strings.Contains(strings.Join(cards, ""), "J") {}
		return scores["two pair"]
	}
	if pairs == 1 {
		if three {
			if kind["J"] == 2 || kind["J"] == 3 {
				// e.g. AAAJJ
				// e.g. AAJJJ
				return scores["five of a kind"]
			}
			// e.g. AAAQQ
			return scores["full house"]
		}
		if kind["J"] == 2 {
			// the single pair is JJ
			// the two J's count as any of the other cards
			// the other cards must all be different here
			return scores["three of a kind"]
		}
		if kind["J"] == 1 {
			// could also be "two pair" but "three of a kind" is better
			return scores["three of a kind"]
		}
		return scores["one pair"]
	}
	if three {
		if kind["J"] == 1 {
			// e.g. 222JA
			return scores["four of a kind"]
		}
		if kind["J"] == 3 {
			// e.g. JJJ2A
			return scores["four of a kind"]
		}
		return scores["three of a kind"]
	}
	if len(kind) == 5 {
		if kind["J"] == 1 {
			return scores["one pair"]
		}
		return scores["high card"]
	}
	return 0 // shouldn't be possible...
}

func partTwo() int {
	hands := []Hand{}
	for _, line := range input {
		sp := strings.Split(line, " ")
		cards := strings.Split(sp[0], "")
		bid, _ := strconv.Atoi(sp[1])
		hands = append(hands, Hand{cards: cards, bid: bid, score: getHandScorePartTwo(cards)})
	}

	slices.SortFunc(hands, handSort)
	total := 0
	for rank, hand := range hands {
		// fmt.Printf("cards: %v / bid: %v / score: %v => rank: %v\n", hand.cards, hand.bid, hand.score, rank+1)
		total += (rank + 1) * hand.bid
	}
	return total
}

func Call(part string, inputFile string) string {
	input = util.ParseInputIntoLines(inputFile)
	var r int
	if part == "1" {
		whichPart = 1
		r = partOne()
	} else {
		whichPart = 2
		r = partTwo()
	}
	return strconv.Itoa(r)
}
