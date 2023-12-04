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

func addPart(parts *[]int, part string) {
	value, err := strconv.Atoi(string(part))
	check(err)
	*parts = append(*parts, value)
}

func main() {
	data, err := os.ReadFile("../input.txt")
	check(err)

	lines := strings.Split(string(data), "\n")
	lineCount := len(lines)
	parts := make([]int, 0)
	re := regexp.MustCompile(`\d+`)

	// Regex for all special characters except dot
	reSpecialChars := regexp.MustCompile(`[^\d\w\s\.]`)

	for i, line := range lines {
		if len(line) == 0 {
			continue
		}

		numbersMatch := re.FindAllString(line, -1)
		numbersIndexes := re.FindAllIndex([]byte(line), -1)

		for matchIndex, match := range numbersMatch {
			// Same line cases
			if numbersIndexes[matchIndex][0] > 0 && reSpecialChars.Match([]byte{line[numbersIndexes[matchIndex][0]-1]}) {
				addPart(&parts, match)
				continue
			}

			if numbersIndexes[matchIndex][1] < len(line)-1 && reSpecialChars.Match([]byte{line[numbersIndexes[matchIndex][1]]}) {
				addPart(&parts, match)
				continue
			}

			// Previous line cases
			if i > 0 {
				previousLineStartIndex := 0
				previousLineEndIndex := len(lines[i-1]) - 1

				if numbersIndexes[matchIndex][0] > 0 {
					previousLineStartIndex = numbersIndexes[matchIndex][0] - 1
				}

				if numbersIndexes[matchIndex][1] < len(lines[i-1])-2 {
					previousLineEndIndex = numbersIndexes[matchIndex][1] + 1
				}

				previousLineSegment := lines[i-1][previousLineStartIndex:previousLineEndIndex]

				if reSpecialChars.Match([]byte(previousLineSegment)) {
					addPart(&parts, match)
					continue
				}
			}

			// Next line cases.
			if i < lineCount-1 {
				nextLineStartIndex := 0
				nextLineEndIndex := len(lines[i+1]) - 1

				if numbersIndexes[matchIndex][0] > 0 {
					nextLineStartIndex = numbersIndexes[matchIndex][0] - 1
				}

				if numbersIndexes[matchIndex][1] < len(lines[i+1])-2 {

					nextLineEndIndex = numbersIndexes[matchIndex][1] + 1
				}

				nextLineSegment := lines[i+1][nextLineStartIndex:nextLineEndIndex]

				if reSpecialChars.Match([]byte(nextLineSegment)) {
					addPart(&parts, match)
					continue
				}
			}
		}
	}

	sum := 0
	for _, part := range parts {
		sum += part
	}

	fmt.Printf("Result is %v\n", sum)
}
