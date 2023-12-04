package main

import (
	"fmt"
	"math"
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
	score := 0

	for _, line := range lines {
		currentScore := 0

		if line == "" {
			break
		}

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
					currentScore++
				}
			}
		}

		if currentScore > 0 {
			score += int(math.Pow(2, float64(currentScore-1)))
		}
	}

	fmt.Printf("Final score is: %v\n", score)
}
