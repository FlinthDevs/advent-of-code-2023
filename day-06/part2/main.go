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

func getRacesData(lines []string) [2]int {
	re := regexp.MustCompile("[0-9]+")
	racesTimes := re.FindAllString(strings.Split(lines[0], ":")[1], -1)
	racesDistances := re.FindAllString(strings.Split(lines[1], ":")[1], -1)

	timeString, distanceString := "", ""

	for i := 0; i < len(racesTimes); i++ {
		timeString += racesTimes[i]
		distanceString += racesDistances[i]
	}

	time, err := strconv.Atoi(timeString)
	check(err)
	distance, err2 := strconv.Atoi(distanceString)
	check(err2)

	return [2]int{time, distance}
}

func main() {
	lines := getLines("../input.txt")
	race := getRacesData(lines)

	possibleSolutions := 0

	for j := 0; j < race[0]; j++ {
		if j*(race[0]-j) > race[1] {
			possibleSolutions++
		}
	}

	fmt.Printf("Combinations total: %v\n", possibleSolutions)
}
