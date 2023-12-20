package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
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

func getLowestLocations(lines []string) [][]int {
	mappings := make([][]int, 0)
	skipNext := false

	for _, line := range lines {
		// Line is empty means next is the map name and can be skipped, and stop updating current indexes
		if line == "" {
			skipNext = true
			continue
		}

		// Name of the map block, prints and skips to actual values.
		if skipNext {
			skipNext = false
			mappings = make([][]int, 0)
			continue
		}

		mapRanges := make([]int, 3)

		// Convert strings values to int.
		for i, s := range strings.Split(strings.Trim(line, " "), " ") {
			val, err := strconv.Atoi(s)
			check(err)
			mapRanges[i] = val
		}

		mappings = append(mappings, mapRanges)
	}

	for i := range mappings {
		for j := range mappings {
			if mappings[i][0] < mappings[j][0] {
				mapRanges := mappings[j]
				mappings[j] = mappings[i]
				mappings[i] = mapRanges
			}
		}

	}

	return mappings
}

func isSeedValid(seedsRanges [][]int, target int) bool {
	for _, seedsRange := range seedsRanges {
		if target >= seedsRange[0] && target < seedsRange[0]+seedsRange[1] {
			return true
		}
	}

	return false
}

func main() {
	lines := getLines("../input.txt")
	rawSeeds := strings.Split(strings.Trim(strings.Split(lines[0], ":")[1], " "), " ")
	seeds := make([][]int, 0)
	IsLetter := regexp.MustCompile(`^[a-zA-Z]+$`).MatchString

	lowestLocations := getLowestLocations(lines[1:])

	highestValue := lowestLocations[len(lowestLocations)-1][0] + lowestLocations[len(lowestLocations)-1][2]
	fmt.Printf("Highest value is %v\n", highestValue)

	// First loop to get the total length of the array.
	for i, seed := range rawSeeds {
		val, err := strconv.Atoi(seed)
		check(err)

		if i%2 == 0 {
			val2, err2 := strconv.Atoi(rawSeeds[i+1])
			check(err2)
			seeds = append(seeds, []int{val, val2})
		}
	}

	// Sorts seeds from lowest to highest
	for i := range seeds {
		for j := range seeds {
			if seeds[i][0] < seeds[j][0] {
				mapRanges := seeds[j]
				seeds[j] = seeds[i]
				seeds[i] = mapRanges
			}
		}
	}

	// If true this means we found a match. Avoid more computation until next block.
	// Starts at true because we do not want to parse the location block, we already did.
	startTime := time.Now()
	stepTime := startTime

	mappings := make([][][]int, 0)
	currentMapping := make([][]int, 0)
	currentBlock := 0

	// Put all mappings in precalculated array.
	for i := len(lines) - 2; i > 0; i-- {
		if lines[i] == "" || IsLetter(string(lines[i][0])) {
			if lines[i] == "" {
				currentBlock++
			}

			mappings = append(mappings, currentMapping)
			currentMapping = make([][]int, 0)
			continue
		}

		mappingRanges := make([]int, 3)

		for j, s := range strings.Split(strings.Trim(lines[i], " "), " ") {
			val, err := strconv.Atoi(s)
			check(err)
			mappingRanges[j] = val
		}

		currentMapping = append(currentMapping, mappingRanges)
	}

	step := 100000

	fmt.Printf("Going to process %v lines (%v steps)\n", highestValue, highestValue/step)

	for currentLocation := 0; currentLocation < highestValue; currentLocation++ {
		if currentLocation%step == 0 {
			fmt.Printf(" - %v done (step took %v)\n", currentLocation, time.Since(stepTime))
			stepTime = time.Now()
		}

		value := currentLocation

		for i := 0; i < len(mappings); {
			for j := 0; j < len(mappings[i]); j++ {
				if value >= mappings[i][j][0] && value < mappings[i][j][0]+mappings[i][j][2] {
					//					fmt.Printf("Value %v matched %v\n", value, mappings[i][j])
					value += mappings[i][j][1] - mappings[i][j][0]
					i++
					break
				}
			}

			i++
		}

		if isSeedValid(seeds, value) {
			fmt.Printf("Found it ! It's %v (seed: %v)\n", currentLocation, value)
			break
		}
	}

	fmt.Printf("Done in %v\n", time.Since(startTime))
}
