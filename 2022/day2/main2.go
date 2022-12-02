package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	data, _ := os.ReadFile("input.csv")
	rounds := strings.Split(string(data), "\n")
	score := 0
	shapeScore := map[string]int{
		"A": 1,
		"X": 1,
		"B": 2,
		"Y": 2,
		"C": 3,
		"Z": 3,
	}
	winningShapes := map[string]string{
		"A": "Y",
		"B": "Z",
		"C": "X",
	}
	loosingShapes := map[string]string{
		"A": "Z",
		"B": "X",
		"C": "Y",
	}
	for _, r := range rounds {
		shapes := strings.Fields(r)
		opponent, result := shapes[0], shapes[1]
		if result == "X" {
			score += shapeScore[loosingShapes[opponent]]
		} else if result == "Y" {
			score += shapeScore[opponent] + 3
		} else {
			score += shapeScore[winningShapes[opponent]] + 6
		}
	}
	fmt.Println(score)
}
