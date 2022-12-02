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
	for _, r := range rounds {
		shapes := strings.Fields(r)
		opponent, me := shapeScore[shapes[0]], shapeScore[shapes[1]]
		if opponent == me {
			score += me + 3
		} else if (me == 2 && opponent == 1) || (me == 1 && opponent == 3) || (me == 3 && opponent == 2) {
			score += me + 6
		} else {
			score += me
		}
	}
	fmt.Println(score)
}
