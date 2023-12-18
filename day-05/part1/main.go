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

func main() {
	lines := getLines("../input.txt")
	skipNext := false
	indexes := make([]int, 0)

	for _, seed := range strings.Split(strings.Trim(strings.Split(lines[0], ":")[1], " "), " ") {
		val, err := strconv.Atoi(seed)
		check(err)
		indexes = append(indexes, val)
	}

	updateIndexes := make([]bool, len(indexes))

	fmt.Printf("Seeds being processed: %v\n", indexes)

	for _, line := range lines[1:] {
		// Line is empty means next is the map name and can be skipped, and stop updating current indexes
		if line == "" {
			skipNext = true
			continue
		}

		// Name of the map block, prints and skips to actual values.
		if skipNext {
			skipNext = false
			for i := range updateIndexes {
				updateIndexes[i] = false
			}

			continue
		}

		mapRanges := make([]int, 3)

		// Convert strings values to int.
		for i, s := range strings.Split(strings.Trim(line, " "), " ") {
			val, err := strconv.Atoi(s)
			check(err)
			mapRanges[i] = val
		}

		// Update indexes based on map values.
		for i, v := range indexes {
			if !updateIndexes[i] && v < mapRanges[1]+mapRanges[2] && v >= mapRanges[1] {
				indexes[i] += mapRanges[0] - mapRanges[1]
				updateIndexes[i] = true
			}
		}
	}

	fmt.Printf("End values: %v\n", indexes)

	minVal := indexes[0]
	for _, v := range indexes {
		if v < minVal {
			minVal = v
		}
	}

	fmt.Printf("Lowest value is %v\n", minVal)
}
