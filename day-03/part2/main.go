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

func checkLine(reNumbers *regexp.Regexp, line string, gearsIndexes []int) []int {
	numbersIndexesOnTargetLine := reNumbers.FindAllIndex([]byte(line), -1)
	numbersOnTargetLine := reNumbers.FindAll([]byte(line), -1)

	results := make([]int, 0, 2)

	for n, number := range numbersOnTargetLine {
		convertedNumber, err := strconv.Atoi(string(number))
		check(err)

		if numbersIndexesOnTargetLine[n][0] >= gearsIndexes[0]-1 && numbersIndexesOnTargetLine[n][0] <= gearsIndexes[1] {
			fmt.Printf("Number %v matches on previous line (beginning check)\n", convertedNumber)
			results = append(results, convertedNumber)
			continue
		}

		if numbersIndexesOnTargetLine[n][1] >= gearsIndexes[0] && numbersIndexesOnTargetLine[n][1] <= gearsIndexes[1]+1 {
			fmt.Printf("Number %v matches on previous line (end check)\n", convertedNumber)
			results = append(results, convertedNumber)
			continue
		}
	}

	return results
}

func main() {
	data, err := os.ReadFile("../input.txt")
	check(err)

	lines := strings.Split(string(data), "\n")
	lineCount := len(lines)
	gearsRatio := 0

	reNumbers := regexp.MustCompile(`\d+`)
	reGear := regexp.MustCompile(`\*`)

	for i, line := range lines {
		fmt.Printf("\n==================\nLine %v: %s\n\n", i, line)

		if len(line) == 0 {
			break
		}

		gearsMatches := reGear.FindAllString(line, -1)
		gearsIndexes := reGear.FindAllIndex([]byte(line), -1)

		fmt.Printf("Gears: %s\n", gearsMatches)
		fmt.Printf("Indexes: %v\n", gearsIndexes)

		if i > 0 {
			fmt.Printf("\nPrevious line: %s\n", lines[i-1])
		}

		if i < lineCount-1 {
			fmt.Printf("\nNext line: %s\n", lines[i+1])
		}

		for gearMatchIndex := range gearsMatches {
			parts := make([]int, 0, 2)
			numbersIndexesOnSameLine := reNumbers.FindAllIndex([]byte(line), -1)
			numbersOnSameLine := reNumbers.FindAll([]byte(line), -1)

			for n, number := range numbersOnSameLine {
				convertedNumber, err := strconv.Atoi(string(number))
				check(err)

				if numbersIndexesOnSameLine[n][0] == gearsIndexes[gearMatchIndex][1] || numbersIndexesOnSameLine[n][1] == gearsIndexes[gearMatchIndex][0] {
					fmt.Printf("Number %v is on the same line as the gear\n", convertedNumber)
					parts = append(parts, convertedNumber)
					continue
				}
			}

			// Previous line cases
			if i > 0 {
				parts = append(parts, checkLine(reNumbers, lines[i-1], gearsIndexes[gearMatchIndex])...)
			}

			// Next line cases.
			if i < lineCount-1 {
				parts = append(parts, checkLine(reNumbers, lines[i+1], gearsIndexes[gearMatchIndex])...)
			}

			fmt.Printf("--------\n")

			if len(parts) == 2 {
				addedValue := parts[0] * parts[1]
				gearsRatio += addedValue
				fmt.Printf("New ratio is %v (%v added)\n", gearsRatio, addedValue)
			}
		}

	}

	fmt.Printf("Result is now %v\n", gearsRatio)
}
