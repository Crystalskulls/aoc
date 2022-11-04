package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"sort"
	"time"
)

type Point struct {
	value    int
	neighbors []*Point
	isLowPoint bool
	basinSize int
	viewed bool
}

type Heightmap struct {
	rows [][]*Point
}

func main() {
	start := time.Now()
	heightmap := parseHeightmap(openCsvFile("input.csv"))
	lowestPoints := heightmap.findLowestPoints()

	var basins []int
	for _, point := range lowestPoints {
		point.calcBasinSize()
		basins = append(basins, point.basinSize)
	}
	sort.Ints(basins)
	largestBasins := basins[len(basins)-3:]
	result := 1
	for _, basin := range largestBasins {
		result *= basin
	}
	fmt.Println(result)
	fmt.Println("time: ", time.Since(start))
}

func (point *Point) calcBasinSize() {
	basinSize := 1
	point.viewed = true
	remainingNeighbors := point.neighbors
	stillNeighbors := true
	for stillNeighbors {
		stillNeighbors = false
		for _, neighbor := range remainingNeighbors {
			if neighbor.viewed {
				continue
			}
			neighbor.viewed = true
			if neighbor.value == 9 {
				continue
			}
			basinSize++
			stillNeighbors = true
			remainingNeighbors = append(remainingNeighbors, neighbor.neighbors...)
		}
	}
	point.basinSize = basinSize
}

func (heightmap *Heightmap) findLowestPoints() []*Point {
	var lowestPoints []*Point

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
				lowestPoints = append(lowestPoints, point)
			}
		}
	}
	return lowestPoints
}

func (point *Point) isLowestPoint() bool {
	for _, neighbor := range point.neighbors {
		if neighbor.value <= point.value {
			point.isLowPoint = false
			return false
		}
	}
	point.isLowPoint = true
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
