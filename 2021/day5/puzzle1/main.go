package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

type HydroVent int

type OceanFloor struct {
	rows           [][]HydroVent
	dangerousAreas int
}

const ix1, iy1, ix2, iy2 int = 0, 1, 2, 3

func main() {
	coordinates := parseVentCoordinates(openCsvFile("input.csv"))
	oceanFloor := newOceanFloor()
	oceanFloor.mapCoordinates(coordinates)
	fmt.Printf("dangerousAres: %d\n", oceanFloor.dangerousAreas)
}

func parseVentCoordinates(data [][]string) (coordinates [][]int) {
	for _, coordinate := range data {
		row := []int{
			parseToInt(coordinate[ix1]),
			parseToInt(coordinate[iy1]),
			parseToInt(coordinate[ix2]),
			parseToInt(coordinate[iy2]),
		}
		coordinates = append(coordinates, row)
	}
	return coordinates
}

func newOceanFloor() *OceanFloor {
	oceanFloor := new(OceanFloor)
	oceanFloor.rows = make([][]HydroVent, 1000)
	for i := 0; i < 1000; i++ {
		oceanFloor.rows[i] = make([]HydroVent, 1000)
	}
	return oceanFloor
}

func (oceanFloor *OceanFloor) mapCoordinates(coordinates [][]int) {
	for _, c := range coordinates {
		x1, y1, x2, y2 := c[ix1], c[iy1], c[ix2], c[iy2]
		if (x1 != x2) && (y1 != y2) {
			// only consider horizontal and vertical lines
			continue
		} else if (x1 == x2) && (y1 == y2) {
			// single point
			oceanFloor.rows[y1][x1]++
			oceanFloor.checkForOverlappingVents(x1, y1)
		} else {
			oceanFloor.drawVentLine(x1, y1, x2, y2)
		}
	}
}

func (oceanFloor *OceanFloor) drawVentLine(x1, y1, x2, y2 int) {
	if (x1 == x2) && (y1 != y2) {
		// vertical line
		if y1 < y2 {
			for i := y1; i <= y2; i++ {
				oceanFloor.rows[i][x1]++
				oceanFloor.checkForOverlappingVents(x1, i)
			}
		} else {
			for i := y1; i >= y2; i-- {
				oceanFloor.rows[i][x1]++
				oceanFloor.checkForOverlappingVents(x1, i)
			}
		}
	} else {
		// horizontal line
		if x1 < x2 {
			for i := x1; i <= x2; i++ {
				oceanFloor.rows[y1][i]++
				oceanFloor.checkForOverlappingVents(i, y1)
			}
		} else {
			for i := x1; i >= x2; i-- {
				oceanFloor.rows[y1][i]++
				oceanFloor.checkForOverlappingVents(i, y1)
			}
		}
	}
}

func (oceanFloor *OceanFloor) checkForOverlappingVents(x, y int) {
	if oceanFloor.rows[y][x] == 2 {
		oceanFloor.dangerousAreas++
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

func parseToInt(n string) int {
	i, err := strconv.ParseInt(n, 10, 64)
	if err != nil {
		log.Fatalf("Unable to parse number %s to an Integer; err: %v\n", n)
	}
	return int(i)
}
