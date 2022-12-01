package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

type seat struct {
	status     string
	nextStatus string
	neighbors  []*seat
}

func main() {
	data := readCSV("./input.csv")
	waitingArea := parseSeats(data)
	fmt.Println(processRules(waitingArea))
}

func processRules(waitingArea [][]*seat) int {
	changed := true
	for changed {
		changed = false
		for _, row := range waitingArea {
			for _, seat := range row {
				if seat.status == "." {
					continue
				}
				if seat.status == "L" && !seat.hasOccupiedNeighbors() {
					seat.nextStatus = "#"
					changed = true
					continue
				}
				if seat.status == "#" && seat.occupiedNeighbors() >= 4 {
					seat.nextStatus = "L"
					changed = true
				}
			}
		}
		setNextStatus(waitingArea)
	}
	return countOccupiedSeats(waitingArea)
}

func countOccupiedSeats(waitingArea [][]*seat) int {
	i := 0
	for _, row := range waitingArea {
		for _, seat := range row {
			if seat.status == "#" {
				i++
			}
		}
	}
	return i
}

func setNextStatus(waitingArea [][]*seat) {
	for _, row := range waitingArea {
		for _, seat := range row {
			seat.status = seat.nextStatus
		}
	}
}

func (s *seat) occupiedNeighbors() int {
	i := 0
	for _, n := range s.neighbors {
		if n.status == "#" {
			i++
		}
	}
	return i
}

func (s *seat) hasOccupiedNeighbors() bool {
	for _, n := range s.neighbors {
		if n.status == "#" {
			return true
		}
	}
	return false
}

func parseSeats(records []string) [][]*seat {
	waitingArea := make([][]*seat, len(records))
	for row := range waitingArea {
		waitingArea[row] = make([]*seat, len(records[0]))
	}
	for i, r := range records {
		for j, v := range r {
			s := newSeat(string(v))
			waitingArea[i][j] = s
		}
	}
	for i, row := range waitingArea {
		for j, seat := range row {
			neighborIndexes := getNeighborIndexes(i, j)
			for _, ni := range neighborIndexes {
				if ni[0] < 0 || ni[1] < 0 || ni[0] > len(records)-1 || ni[1] > len(records[0])-1 {
					continue
				}
				seat.neighbors = append(seat.neighbors, waitingArea[ni[0]][ni[1]])
			}
		}
	}
	return waitingArea
}

func newSeat(status string) *seat {
	s := new(seat)
	s.status = status
	s.neighbors = make([]*seat, 0)
	return s
}

func getNeighborIndexes(i, j int) [][]int {
	m := [][]int{
		{-1, -1},
		{-1, 0},
		{-1, 1},
		{0, -1},
		{0, 1},
		{1, -1},
		{1, 0},
		{1, 1},
	}

	k := make([][]int, 8)
	for l, _ := range k {
		k[l] = append(k[l], i+m[l][0])
		k[l] = append(k[l], j+m[l][1])
	}
	return k
}

func readCSV(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Can't open file %s\n", path)
	}
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("Can't read file %s\n", path)
	}
	lines := make([]string, 0)
	for _, record := range records {
		lines = append(lines, record[0])
	}
	return lines
}
