package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

type Point struct {
	value    int
	neighbors []*Point
}

type Heightmap struct {
	rows [][]*Point
}

func main() {
	heightmap := parseHeightmap(openCsvFile("input.csv"))
	lowestPoints := heightmap.findLowestPoints()
	riskLevel := calcRiskLevel(lowestPoints)
	fmt.Println(lowestPoints)
	fmt.Println(riskLevel)
}

func calcRiskLevel(lowestPoints []int) int {
	riskLevel := 0
	for _, v := range lowestPoints {
		riskLevel += v + 1
	}
	return riskLevel
}

func (heightmap *Heightmap) findLowestPoints() []int {
	var lowestPoints []int

	for i, row := range heightmap.rows {
		for j, point := range row {
			if i == 0 && j == 0 {
				point.neighbors = append(point.neighbors, heightmap.rows[i][j+1])
				point.neighbors = append(point.neighbors, heightmap.rows[i+1][j])
			} else if i == 0 && j == len(row)-1 {
				point.neighbors = append(point.neighbors, heightmap.rows[i][j-1])
				point.neighbors = append(point.neighbors, heightmap.rows[i+1][j])
			} else if i == len(heightmap.rows)-1 && j == 0 {
				point.neighbors = append(point.neighbors, heightmap.rows[i][j+1])
				point.neighbors = append(point.neighbors, heightmap.rows[i-1][j])
			} else if i == len(heightmap.rows)-1 && j == len(row)-1 {
				point.neighbors = append(point.neighbors, heightmap.rows[i][j-1])
				point.neighbors = append(point.neighbors, heightmap.rows[i-1][j])
			} else if i == 0 {
				point.neighbors = append(point.neighbors, heightmap.rows[i][j-1])
				point.neighbors = append(point.neighbors, heightmap.rows[i][j+1])
				point.neighbors = append(point.neighbors, heightmap.rows[i+1][j])
			} else if i == len(heightmap.rows)-1 {
				point.neighbors = append(point.neighbors, heightmap.rows[i][j-1])
				point.neighbors = append(point.neighbors, heightmap.rows[i][j+1])
				point.neighbors = append(point.neighbors, heightmap.rows[i-1][j])
			} else if j == 0 {
				point.neighbors = append(point.neighbors, heightmap.rows[i-1][j])
				point.neighbors = append(point.neighbors, heightmap.rows[i+1][j])
				point.neighbors = append(point.neighbors, heightmap.rows[i][j+1])
			} else if j == len(row)-1 {
				point.neighbors = append(point.neighbors, heightmap.rows[i-1][j])
				point.neighbors = append(point.neighbors, heightmap.rows[i+1][j])
				point.neighbors = append(point.neighbors, heightmap.rows[i][j-1])
			} else {
				point.neighbors = append(point.neighbors, heightmap.rows[i][j-1])
				point.neighbors = append(point.neighbors, heightmap.rows[i][j+1])
				point.neighbors = append(point.neighbors, heightmap.rows[i-1][j])
				point.neighbors = append(point.neighbors, heightmap.rows[i+1][j])
			}

			if point.isLowestPoint() {
				lowestPoints = append(lowestPoints, point.value)
			}
		}
	}
	return lowestPoints
}

func (point *Point) isLowestPoint() bool {
	for _, neighbor := range point.neighbors {
		if neighbor.value <= point.value {
			return false
		}
	}
	return true
}

func parseHeightmap(data [][]string) *Heightmap {
	heightmap := new(Heightmap)
	for _, row := range data {
		var points []*Point
		for _, s := range row {
			for _, r := range s {
				point := new(Point)
				point.value = int(r - '0')
				points = append(points, point)
			}
		}
		heightmap.rows = append(heightmap.rows, points)
	}
	return heightmap
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
