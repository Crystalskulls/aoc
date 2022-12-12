package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

type Node struct {
	costs     *int
	elevation rune
	neighbors []*Node
}

func main() {
	heightMap, myCurrentPosition, startNodes, endNode := getHeightMap()
	var fewestSteps *int
	for _, startNode := range startNodes {
		resetCosts(heightMap)
		startNode.costs = new(int)
		queue := make([]*Node, 0)
		queue = append(queue, startNode)

		for len(queue) > 0 {
			currentNode := queue[0]
			queue = queue[1:]
			for _, neighbor := range currentNode.neighbors {
				if neighbor.elevation <= currentNode.elevation+1 {
					if neighbor.costs == nil {
						neighbor.costs = new(int)
						*neighbor.costs = *currentNode.costs + 1
						queue = appendToQueue(queue, neighbor)
					} else {
						if *currentNode.costs+1 < *neighbor.costs {
							*neighbor.costs = *currentNode.costs + 1
						}
					}
				}
			}
		}
		if endNode.costs == nil {
			continue
		}
		if fewestSteps == nil {
			fewestSteps = new(int)
			*fewestSteps = *endNode.costs
		} else if *endNode.costs < *fewestSteps {
			*fewestSteps = *endNode.costs
		}
		if startNode == myCurrentPosition {
			fmt.Printf("Part One: %d\n", *endNode.costs)
		}
	}

	fmt.Printf("Part Two: %d\n", *fewestSteps)
}

func resetCosts(heightmap [][]*Node) {
	for _, line := range heightmap {
		for _, node := range line {
			node.costs = nil
		}
	}
}

func appendToQueue(queue []*Node, newNode *Node) []*Node {
	index := 0
	for i, node := range queue {
		if newNode.elevation > node.elevation {
			index = i
		}
	}
	if len(queue) == index {
		queue = append(queue, newNode)
		return queue
	}
	queue = append(queue[:index+1], queue[index:]...)
	queue[index] = newNode
	return queue
}

func getHeightMap() (heightmap [][]*Node, myCurrentPosition *Node, startNodes []*Node, endNode *Node) {
	file, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(file), "\n")
	heightmap = make([][]*Node, len(lines))
	startNodes = make([]*Node, 0)
	for i, line := range lines {
		heightmap[i] = make([]*Node, len(line))
		for j, r := range line {
			node := new(Node)
			node.elevation = r
			if string(r) == "a" {
				startNodes = append(startNodes, node)
			}
			if string(r) == "S" {
				node.elevation = rune(97)
				myCurrentPosition = node
				startNodes = append(startNodes, node)
			}
			if string(r) == "E" {
				node.elevation = rune(122)
				endNode = node
			}
			heightmap[i][j] = node
		}
	}

	for i, line := range heightmap {
		for j, node := range line {
			node.neighbors = getNeighbors(heightmap, float64(i), float64(j))
		}
	}
	return
}

func getNeighbors(heightmap [][]*Node, i, j float64) (neighbors []*Node) {
	neighbors = make([]*Node, 0)
	rowLimit, columnLimit := float64(len(heightmap)-1), float64(len(heightmap[0])-1)
	for x := math.Max(0, i-1); x <= math.Min(i+1, rowLimit); x++ {
		for y := math.Max(0, j-1); y <= math.Min(j+1, columnLimit); y++ {
			if (x != i || y != j) && (x == i || y == j) {
				neighbors = append(neighbors, heightmap[int(x)][int(y)])
			}
		}
	}
	return neighbors
}
