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

func getRacesData(lines []string) [][]int {
	re := regexp.MustCompile("[0-9]+")
	racesTimes := re.FindAllString(strings.Split(lines[0], ":")[1], -1)
	racesDistances := re.FindAllString(strings.Split(lines[1], ":")[1], -1)

	racesData := make([][]int, len(racesTimes))

	for i := 0; i < len(racesTimes); i++ {
		racesData[i] = make([]int, 2)
		time, errTime := strconv.Atoi(racesTimes[i])
		distance, errDistance := strconv.Atoi(racesDistances[i])
		check(errDistance)
		check(errTime)

		racesData[i][0] = time
		racesData[i][1] = distance
	}

	return racesData
}

func main() {
	lines := getLines("../input.txt")
	races := getRacesData(lines)
	result := 1

	for i := 0; i < len(races); i++ {
		possibleSolutions := 0

		for j := 0; j < races[i][0]; j++ {
			if j*(races[i][0]-j) > races[i][1] {
				possibleSolutions++
			}
		}

		result *= possibleSolutions
	}

	fmt.Printf("Combinations total: %v\n", result)
}
