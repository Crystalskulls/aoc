package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

type RiskLevel struct {
	neighbors []*RiskLevel
	value int
	entered bool
}

type Cavern struct {
	riskLevels [][]*RiskLevel
}

func main() {
	cavern := mapCavern(openCsvFile("input.csv"))
	fmt.Println(cavern)
}

func (cavern *Cavern) findPath(

func mapCavern(data [][]string) *Cavern {
	cavern := new(Cavern)
	for _, line := range data {
		var riskLevels []*RiskLevel
		for _, v := range line[0] {
			riskLevel := new(RiskLevel)
			riskLevel.value = parseToInt(v)
			riskLevels = append(riskLevels, riskLevel)
		}
		cavern.riskLevels = append(cavern.riskLevels, riskLevels)
	}

	for i, line := range cavern.riskLevels {
		for j, riskLevel := range line {
			riskLevel.findNeighbors(i, j, cavern)
		}
	}
	return cavern
}

func (riskLevel *RiskLevel) findNeighbors(row, column int, cavern *Cavern) {
	lastIndex := len(cavern.riskLevels)-1
	if column == 0 && row == 0 {
		// top-left
		riskLevel.neighbors = append(riskLevel.neighbors, cavern.riskLevels[row][column+1])
		riskLevel.neighbors = append(riskLevel.neighbors, cavern.riskLevels[row+1][column])
	} else if column == 0 && row == lastIndex {
		// bottom-left
		riskLevel.neighbors = append(riskLevel.neighbors, cavern.riskLevels[row][column+1])
		riskLevel.neighbors = append(riskLevel.neighbors, cavern.riskLevels[row-1][column])
	} else if column == lastIndex && row == 0 {
		// top-right
		riskLevel.neighbors = append(riskLevel.neighbors, cavern.riskLevels[row][column-1])
		riskLevel.neighbors = append(riskLevel.neighbors, cavern.riskLevels[row+1][column])
	} else if column == lastIndex && row == lastIndex {
		// bottom-right
		riskLevel.neighbors = append(riskLevel.neighbors, cavern.riskLevels[row][column-1])
		riskLevel.neighbors = append(riskLevel.neighbors, cavern.riskLevels[row-1][column])
	} else if column == 0 {
		// edge-left
		riskLevel.neighbors = append(riskLevel.neighbors, cavern.riskLevels[row][column+1])
		riskLevel.neighbors = append(riskLevel.neighbors, cavern.riskLevels[row-1][column])
		riskLevel.neighbors = append(riskLevel.neighbors, cavern.riskLevels[row+1][column])
	} else if column == lastIndex {
		// edge-right
		riskLevel.neighbors = append(riskLevel.neighbors, cavern.riskLevels[row][column-1])
		riskLevel.neighbors = append(riskLevel.neighbors, cavern.riskLevels[row-1][column])
		riskLevel.neighbors = append(riskLevel.neighbors, cavern.riskLevels[row+1][column])
	} else if row == 0 {
		// edge-top
		riskLevel.neighbors = append(riskLevel.neighbors, cavern.riskLevels[row][column+1])
		riskLevel.neighbors = append(riskLevel.neighbors, cavern.riskLevels[row][column-1])
		riskLevel.neighbors = append(riskLevel.neighbors, cavern.riskLevels[row+1][column])
	} else if row == lastIndex {
		// edge-bottom
		riskLevel.neighbors = append(riskLevel.neighbors, cavern.riskLevels[row][column+1])
		riskLevel.neighbors = append(riskLevel.neighbors, cavern.riskLevels[row][column-1])
		riskLevel.neighbors = append(riskLevel.neighbors, cavern.riskLevels[row-1][column])
	} else {
		// center
		riskLevel.neighbors = append(riskLevel.neighbors, cavern.riskLevels[row][column+1])
		riskLevel.neighbors = append(riskLevel.neighbors, cavern.riskLevels[row][column-1])
		riskLevel.neighbors = append(riskLevel.neighbors, cavern.riskLevels[row+1][column])
		riskLevel.neighbors = append(riskLevel.neighbors, cavern.riskLevels[row-1][column])
	}
}

func openCsvFile(path string) [][]string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Unable to open csv file %s; err: %v\n", path, err)
	}
	reader := csv.NewReader(file)
	data, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("Unable to read csv file %s; err: %v\n", path, err)
	}
	return data
}

func parseToInt(r rune) int {
	s := string(r)
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		log.Fatalf("Unable to parse number %s to an Integer; err: %v\n", s)
	}
	return int(i)
}
