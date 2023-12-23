package main

import (
	"fmt"
	"os"
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

func getInt(s string) int {
	v, e := strconv.Atoi(s)
	check(e)
	return v
}

func getHistories(lines []string) [][]int {
	result := make([][]int, len(lines))

	for l, line := range lines {
		values := strings.Split(line, " ")
		valuesInt := make([]int, len(values))

		for i := range values {
			valuesInt[i] = getInt(values[i])
		}
		result[l] = valuesInt
	}

	return result
}

func isAllZeroes(toCheck []int) bool {
	for _, n := range toCheck {
		if n != 0 {
			return false
		}
	}

	return true
}

func getExtrapolates(histories [][]int) []int {
	result := make([]int, len(histories))

	// Computes all slices of relations histories.
	for h, history := range histories {
		relations := make([][]int, 1)
		relations[0] = history
		lastRelation := relations[:1][0]

		// Add new line until the last line added is all zeroes.
		for i := 1; !isAllZeroes(lastRelation); i++ {
			lastRelation = relations[i-1]
			newRel := make([]int, len(lastRelation)-1)

			for j := 0; j < len(lastRelation)-1; j++ {
				newRel[j] = lastRelation[j+1] - lastRelation[j]
			}

			relations = append(relations, newRel)
			lastRelation = newRel
		}

		// Now extrapolating the last value
		for i := len(relations) - 1; i >= 1; i-- {
			lastCellIndex := len(relations[i]) - 1
			relations[i-1] = append(relations[i-1], relations[i][lastCellIndex]+relations[i-1][lastCellIndex])
		}

		result[h] = relations[0][len(relations[0])-1]
	}

	return result
}

func main() {
	lines := getLines("../input.txt")
	histories := getHistories(lines)
	extrapolates := getExtrapolates(histories)

	sum := 0
	for _, e := range extrapolates {
		sum += e
	}

	fmt.Printf("Sum of extrapolate values: %v\n", sum)
}
