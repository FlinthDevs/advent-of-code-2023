package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// Read lines from input.txt.
	data, err := os.ReadFile("../input.txt")
	check(err)

	games := strings.Split(string(data), "\n")

	// Array of ids that will match.
	matches := make([]int, 0)

	re := regexp.MustCompile(`(\d+)`)

	for _, line := range games {
		if len(line) == 0 {
			continue
		}

		maxReds, maxGreens, maxBlues := 0, 0, 0

		// Parses only after "Game X:" and split by ";"
		sets := strings.Split(strings.Split(line, ":")[1], ";")

		// Loop through each "set"
		for _, set := range sets {
			if len(set) == 0 {
				continue
			}

			// Stores each color count value and its start/end indexes.
			colorsCounts := re.FindAll([]byte(set), -1)
			colorsIndex := re.FindAllIndex([]byte(set), -1)

			for cIndex, colorsCount := range colorsCounts {
				count, err := strconv.Atoi(string(colorsCount))
				check(err)

				// Check first letter of color associated with the cound.
				// Index is {end-of-color-count-value-match}+1
				colorFirstLetter := set[colorsIndex[cIndex][1]+1]

				if colorFirstLetter == 'r' && count > maxReds {
					maxReds = count
				} else if colorFirstLetter == 'g' && count > maxGreens {
					maxGreens = count
				} else if colorFirstLetter == 'b' && count > maxBlues {
					maxBlues = count
				}
			}
		}

		matches = append(matches, maxReds*maxGreens*maxBlues)
	}

	result := 0

	for _, match := range matches {
		result += match
	}

	fmt.Printf("\nResult: %v\n", result)
}
