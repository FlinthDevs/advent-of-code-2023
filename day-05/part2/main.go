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

	fmt.Println("Sorted location ranges:")
	for i := range mappings {
		fmt.Printf(" - Dest: %v\n", mappings[i][0])
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

	// If true this means we found a maétch. Avoid more computation until next block.
	// Starts at true because we do not want to parse the location block, we already did.
	skipUntilNextBlock := true
	// found := false
	// for _, locationRange := range lowestLocations {
	// 	for currentLocation := locationRange[0]; currentLocation < locationRange[1]-1; currentLocation++ {
	for currentLocation := 0; currentLocation < highestValue; currentLocation++ {
		value := currentLocation
		//fmt.Printf("\n\nIs location %v valid ?\n", value)

		// Going reverse from bottom-up
		for i := len(lines) - 2; i > 0; i-- {
			//			fmt.Printf("=== Processing line: %v\n", line)
			//			fmt.Printf("Should skip ? %v\n", skipUntilNextBlock)
			// Line found = starts end of current block and start looking at mappings again.
			if lines[i] == "" {
				// fmt.Println(" - ! Skipping empty line...")
				skipUntilNextBlock = false
				continue
				// If we should skip or block is letters, avoid weird parsings.
			} else if skipUntilNextBlock || IsLetter(string(lines[i][0])) {
				// fmt.Println(" - Skipping because we need to!")
				//	if IsLetter(string(line[0])) {
				//	fmt.Printf("= Result after block %v: %v\n", line, value)
				//}

				continue
			}

			// Stores current line as mapping in int array.
			mapRanges := make([]int, 3)

			for i, s := range strings.Split(strings.Trim(lines[i], " "), " ") {
				val, err := strconv.Atoi(s)
				check(err)
				mapRanges[i] = val
			}
			if mapRanges[0] == 1 && mapRanges[1] == 0 && mapRanges[2] == 69 {
				fmt.Println(currentLocation)
				fmt.Println(currentLocation >= mapRanges[0])
				fmt.Println(currentLocation < mapRanges[0]+mapRanges[2])
				fmt.Println(value)
				fmt.Println(mapRanges[1] - mapRanges[0])
			}

			// If current location is in the range of the current mapping, update it.
			if value >= mapRanges[0] && value < mapRanges[0]+mapRanges[2] {
				value += mapRanges[1] - mapRanges[0]
				skipUntilNextBlock = true
			}
		}

		// // Name of the map block, prints and skips to actual values.
		// if skipNext {
		// 	skipNext = false
		// 	continue
		// }

		// mapRanges := make([]int, 3)

		// // Convert strings values to int.
		// for i, s := range strings.Split(strings.Trim(line, " "), " ") {
		// 	val, err := strconv.Atoi(s)
		// 	check(err)
		// 	mapRanges[i] = val
		// }

		// // Update indexes based on map values.
		// for i, v := range indexes {
		// 	if !updateIndexes[i] && v < mapRanges[1]+mapRanges[2] && v >= mapRanges[1] {
		// 		indexes[i] += mapRanges[0] - mapRanges[1]
		// 		updateIndexes[i] = true
		// 	}
		// }

		if isSeedValid(seeds, value) {
			// found = true
			fmt.Printf("Found it ! It's %v (seed: %v)\n", currentLocation, value)
			break
		}
	}

	// fmt.Printf("End values: %v\n", indexes)

	// minVal := indexes[0]
	// for _, v := range indexes {
	// 	if v < minVal {
	// 		minVal = v
	// 	}
	// }
	// if found {
	// 	break
	// }
	// }

	// fmt.Printf("Lowest value is %v\n", minVal)
}
