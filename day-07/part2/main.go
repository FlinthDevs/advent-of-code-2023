package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getLines(filename string) []string {
	data, err := os.ReadFile(filename)
	check(err)
	return strings.Split(string(data), "\n")
}

// Hand types
const (
	High      = 1 << 0
	OnePair   = 1 << 1
	TwoPair   = 1 << 2
	ThreeKind = 1 << 3
	FullHouse = 1 << 4
	FourKind  = 1 << 5
	FiveKind  = 1 << 6
)

type Hand struct {
	Cards string // Untouched cards
	Type  int    // Hand type in base 2
	Value int    // Hand value in base 10, from cards converted to base13
	Bid   int    // Bid from file
}

// Use in getHandType, calculates the given hand "type" based on current combo.
func getTypeFromCount(count int, previousCount int) int {
	if count == 5 {
		return FiveKind
	} else if count == 4 {
		return FourKind
	} else if count == 3 {
		if previousCount == 2 {
			return FullHouse
		} else {
			return ThreeKind
		}
	} else if count == 2 {
		if previousCount == 3 {
			return FullHouse
		}
		if previousCount == 2 {
			return TwoPair
		} else {
			return OnePair
		}
	}

	return High
}

// Updates handType using joker to maximize the result.
func applyJokers(handType int, jokersCount int) int {
	switch jokersCount {
	case 0:
		return handType
	case 1:
		if handType == High || handType == FourKind {
			return handType << 1
		} else if handType == OnePair || handType == ThreeKind {
			return handType << 2
		}

		return FullHouse
	case 2:
		if handType == OnePair {
			return FourKind
		} else if handType == ThreeKind {
			return FiveKind
		} else if handType == High {
			return ThreeKind
		}
	case 3:
		if handType == High {
			return FourKind
		} else if handType == OnePair {
			return FiveKind
		}
	default:
		return FiveKind
	}

	return handType
}

// Get hand type in base 2
func getHandType(cards string) int {
	sortedCards := strings.Split(cards, "")
	sort.Strings(sortedCards)

	handType := 1
	previousCount, count, jokersCount := 0, 1, 0

	if sortedCards[0] == "J" {
		jokersCount++
	}

	for i := 1; i < len(sortedCards); i++ {
		if sortedCards[i] == "J" {
			jokersCount++
			continue
		}
		// Increase counter when card is the same as previous
		if sortedCards[i] == sortedCards[i-1] {
			count++
		} else if sortedCards[i] != sortedCards[i-1] && count > 1 {
			// If we have a new card and we have a current "combo", updates hand type
			handType = getTypeFromCount(count, previousCount)
			previousCount = count
			count = 1
		}
	}

	// If there is something left in the end, recomputes hand type
	if count > 1 {
		handType = getTypeFromCount(count, previousCount)
	}

	handType = applyJokers(handType, jokersCount)

	return handType
}

// Get card's value in base 13
func getCardValue(card rune) string {
	switch card {
	case 'A':
		return "c"
	case 'K':
		return "b"
	case 'Q':
		return "a"
	case 'J':
		return "0"
	case 'T':
		return "9"
	default:
		return string(card - 1)
	}
}

// Get Base10 hand value from the raw string.
// Doing raw => Base13 => Base10
func getHandValue(cards string) int64 {
	currentHand := ""

	for _, card := range cards {
		currentHand += getCardValue(card)
	}

	value, err := strconv.ParseInt(currentHand, 13, 64)
	check(err)

	return value
}

func getInt(s string) int {
	result, err := strconv.Atoi(s)
	check(err)
	return result
}

// Computes all hands types and values.
func getHands(lines []string) []Hand {
	result := make([]Hand, 0)

	for _, line := range lines {
		values := strings.Split(line, " ")

		result = append(
			result,
			Hand{
				Cards: values[0],
				Type:  getHandType(values[0]),
				Value: int(getHandValue(values[0])),
				Bid:   getInt(values[1]),
			})
	}

	return result
}

// Sort hands by type and value
func getSortedHands(hands []Hand) []Hand {
	sort.Slice(hands, func(i, j int) bool {
		if hands[i].Type == hands[j].Type {
			return hands[i].Value > hands[j].Value
		}
		return hands[i].Type > hands[j].Type
	})

	return hands
}

func main() {
	lines := getLines("../input.txt")
	hands := getHands(lines)

	rankedHands := getSortedHands(hands)
	result := 0
	maxRank := len(rankedHands)

	for i := range rankedHands {
		result += rankedHands[i].Bid * (maxRank - i)
	}

	fmt.Printf("Final result is %v\n", result)
}
