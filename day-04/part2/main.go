package main

import (
	"fmt"
	"os"
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

func main() {
	lines := getLines("../input.txt")
	lineCount := 0
	score := 0

	// List of bonus cards and their count.
	bonusCardsCount := make(map[int]int)

	for l, line := range lines {
		currentMatchesCount := 0
		fmt.Print("---------------------------\nLine ", l, ": ", line, "\n")
		fmt.Printf("This line has %v bonus cards\n", bonusCardsCount[l])

		if line == "" {
			break
		}
		lineCount++

		// Get rid of the "Game XXX" part
		numbers := strings.Split(strings.Split(line, ":")[1], "|")
		winningNumbers := strings.Split(strings.Trim(numbers[0], " "), " ")
		playedNumbers := strings.Split(strings.Trim(numbers[1], " "), " ")

		for _, number := range playedNumbers {
			if number == "" {
				continue
			}

			for _, winningNumber := range winningNumbers {
				if winningNumber == "" {
					continue
				}

				if number == winningNumber {
					currentMatchesCount += 1
				}
			}
		}

		fmt.Printf("Line %v had %v matches\n", l, currentMatchesCount)

		for i := 0; i < currentMatchesCount; i++ {
			if l+i > len(lines)-1 {
				break
			}

			bonusCardsCount[l+i+1] += 1 + bonusCardsCount[l]
			fmt.Printf("Now card %v has %v bonus cards (+%v)\n", l+i+1, bonusCardsCount[l+i+1], bonusCardsCount[l]+1)
		}
	}

	for _, count := range bonusCardsCount {
		score += count
	}

	fmt.Printf("++++++++++++++++++++++++++++++++++++++\nFinal cards count is: %v\n", score+lineCount)
}
