package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

type Area struct {
	trees [][]Tree
}

type Tree bool

type Slope struct {
	right int
	down  int
}

func main() {
	records := readCSV("./input.csv")
	area := parseToArea(records)

	slopes := []Slope{
		{right: 1, down: 1},
		{right: 3, down: 1},
		{right: 5, down: 1},
		{right: 7, down: 1},
		{right: 1, down: 2},
	}

	s := 1
	for _, slope := range slopes {
		c := 0
		j := slope.right
		for i := slope.down; i < len(area.trees); i += slope.down {
			if area.trees[i][j] {
				c++
			}
			j += slope.right
			max := len(area.trees[i])
			if j >= max {
				j = j - max
			}
		}
		s *= c
		fmt.Printf("slope %+v - tree count: %d\n", slope, c)
	}
	fmt.Printf("multiplied together: %d\n", s)
}

func readCSV(path string) [][]string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Can't open file %s\n", path)
	}
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("Can't read file %s\n", path)
	}
	return records
}

func parseToArea(records [][]string) *Area {
	area := new(Area)
	area.trees = make([][]Tree, len(records))

	for i, record := range records {
		area.trees[i] = make([]Tree, len(record[0]))
		for j, b := range record[0] {
			s := rune(b)
			if s == '.' {
				area.trees[i][j] = false
				continue
			}
			area.trees[i][j] = true
		}
	}
	return area
}
