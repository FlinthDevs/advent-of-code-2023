package main

import (
	"fmt"
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

type Node struct {
	Name         string    // Raw name of the node (just in case)
	Id           int       // Base 26 value
	Neighbors    [2]string // Possible next nodes names
	NeighborsIds [2]int    // Possibles next nodes Ids
}

// Get nodes with computed data
func getNodes(lines []string) []Node {
	nodes := make([]Node, len(lines))

	// Firt init array with all values and neighbors named
	for i, line := range lines {
		values := strings.Split(line, " = ")
		neighbors := strings.Split(values[1][1:len(values[1])-1], ", ")

		nodes[i] =
			Node{
				Name: values[0],
				Neighbors: [2]string{
					neighbors[0],
					neighbors[1],
				},
			}
	}

	// Now computes indexes of neighbors
	for i := range nodes {
		nodes[i].NeighborsIds[0] = getNodeIdFromName(nodes, nodes[i].Neighbors[0])
		nodes[i].NeighborsIds[1] = getNodeIdFromName(nodes, nodes[i].Neighbors[1])
	}

	return nodes
}

// Search in all nodes the node matching given name and return its id
func getNodeIdFromName(nodes []Node, name string) int {
	for i := range nodes {
		if nodes[i].Name == name {
			return i
		}
	}

	return -1
}

func getAllNodesEndingWith(nodes []Node, letter rune) []int {
	result := make([]int, 0)

	for i := range nodes {
		if nodes[i].Name[2] == byte(letter) {
			result = append(result, i)
		}
	}

	return result
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

// Calculate all required steps for finding the way.
func getStepsRequired(instructions string, nodes []Node) int {
	startingNodes := getAllNodesEndingWith(nodes, 'A')
	instructionsLength := len(instructions)
	stepCounts := make([]int, len(startingNodes))

	for nodeIndex := range startingNodes {
		for currentNode := startingNodes[nodeIndex]; nodes[currentNode].Name[2] != 'Z'; {
			step := instructions[stepCounts[nodeIndex]%instructionsLength]
			targetNeighbor := 0

			if step == 'R' {
				targetNeighbor = 1
			}

			currentNode = nodes[currentNode].NeighborsIds[targetNeighbor]
			stepCounts[nodeIndex]++
		}
	}

	return LCM(stepCounts[0], stepCounts[1], stepCounts...)
}

func main() {
	lines := getLines("../input.txt")
	instructions := lines[0]
	nodes := getNodes(lines[2:])
	steps := getStepsRequired(instructions, nodes)

	fmt.Printf("Numbers of steps: %v\n", steps)
}
