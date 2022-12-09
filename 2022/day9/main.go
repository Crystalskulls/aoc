package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.ReadFile("input.txt")
	moves := strings.Split(string(file), "\n")
	knotAmount := []int{2, 10}
	for i, knotCount := range knotAmount {
		knots := make([][]int, knotCount)
		for i := range knots {
			knots[i] = make([]int, 2)
		}
		visitedPositions := make(map[string]struct{})
		visitedPositions[fmt.Sprintf("%d%d", knots[0][0], knots[0][1])] = struct{}{}
		for _, move := range moves {
			tmp := strings.Split(move, " ")
			direction := tmp[0]
			steps, _ := strconv.Atoi(tmp[1])
			step(knots, direction, steps, visitedPositions)
		}
		fmt.Printf("Part %d: %d\n", i+1, len(visitedPositions))
	}
}

func step(knots [][]int, direction string, steps int, visitedPositions map[string]struct{}) {
	for s := 0; s < steps; s++ {
		for i := len(knots) - 1; i > 0; i-- {
			if i == len(knots)-1 {
				if direction == "R" {
					knots[i][0]++
				} else if direction == "L" {
					knots[i][0]--
				} else if direction == "U" {
					knots[i][1]++
				} else {
					knots[i][1]--
				}
			}
			if !isStillNeighbor(knots[i-1], knots[i]) {
				x, y := knots[i-1][0], knots[i-1][1]
				if knots[i][0] > x {
					knots[i-1][0]++
				} else if knots[i][0] < x {
					knots[i-1][0]--
				}

				if knots[i][1] > y {
					knots[i-1][1]++
				} else if knots[i][1] < y {
					knots[i-1][1]--
				}

				if i == 1 {
					visitedPositions[fmt.Sprintf("%d%d", knots[0][0], knots[0][1])] = struct{}{}
				}
			}
		}
	}
}

func isStillNeighbor(tail, head []int) bool {
	x1, y1 := float64(tail[0]), float64(tail[1])
	x2, y2 := float64(head[0]), float64(head[1])
	if math.Abs(y2-y1) > 1 {
		return false
	}
	if math.Abs(x2-x1) > 1 {
		return false
	}
	return true
}
