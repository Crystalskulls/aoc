package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"github.com/fatih/color"
	"os"
	"strconv"
)

type Octopus struct {
	energyLevel int
	neighbors []*Octopus
	flashed bool
	readyToFlash bool
}

type Cavern struct {
	octopuses [][]*Octopus
}

func main() {
	cavern := mapCavern(openCsvFile("input.csv"))
	totalFlashes := 0
	fmt.Println(cavern)
	allOctopusesFlashed := false
	for i:= 0; !allOctopusesFlashed; i++ {
		flashes, allOctopusesFlashed := gainEnergy(cavern)
		totalFlashes += flashes
		fmt.Printf("After step %d:\n", i + 1)
		fmt.Println(cavern)
		if allOctopusesFlashed {
			fmt.Println("octopuses flash simultaneously after step: ", i + 1)
			break
		}
	}
	fmt.Println("total flashes: ", totalFlashes)
}

func gainEnergy(cavern *Cavern) (flashes int, allOctopusesFlashed bool) {
	// First, the energy level of each octopus increases by 1.
	for _, row := range cavern.octopuses {
		for _, octopus := range row {
			octopus.increaseEnergyLevel()
		}
	}
	// Then, any octopus with an energy level greater than 9 flashes.
	for _, row := range cavern.octopuses {
		for _, octopus := range row {
			if octopus.readyToFlash && !octopus.flashed{
				octopus.flash()
			}
		}
	}
	// Finally, any octopus that flashed during this step has its energy level set to 0, as it used all of its energy to flash.
	for _, row := range cavern.octopuses {
		for _, octopus := range row {
			if octopus.flashed {
				flashes++
				octopus.energyLevel = 0
				octopus.readyToFlash = false
				octopus.flashed = false
			}
		}
	}
	if flashes == 100 {
		return flashes, true
	}
	return flashes, false
}

func (octopus *Octopus) increaseEnergyLevel() {
	octopus.energyLevel++
	if octopus.energyLevel > 9 {
		octopus.readyToFlash = true
	}
}

func mapCavern(data [][]string) *Cavern {
	cavern := new(Cavern)
	for _, line := range data {
		var octopusRow []*Octopus
		for _, energyLevel := range line[0] {
			octopus := new(Octopus)
			octopus.energyLevel = parseToInt(energyLevel)
			octopus.readyToFlash = false
			octopus.flashed = false
			octopusRow = append(octopusRow, octopus)
		}
		cavern.octopuses = append(cavern.octopuses, octopusRow)
	}

	for i, line := range cavern.octopuses {
		for j, octopus := range line {
			octopus.findNeighbors(i, j, cavern)
		}
	}
	return cavern
}

func (octopus *Octopus) findNeighbors(row, column int, cavern *Cavern) {
	lastIndex := len(cavern.octopuses)-1
	if column == 0 && row == 0 {
		// top-left
		octopus.neighbors = append(octopus.neighbors, cavern.octopuses[row][column+1])
		octopus.neighbors = append(octopus.neighbors, cavern.octopuses[row+1][column])
		octopus.neighbors = append(octopus.neighbors, cavern.octopuses[row+1][column+1])
	} else if column == 0 && row == lastIndex {
		// bottom-left
		octopus.neighbors = append(octopus.neighbors, cavern.octopuses[row-1][column])
		octopus.neighbors = append(octopus.neighbors, cavern.octopuses[row][column+1])
		octopus.neighbors = append(octopus.neighbors, cavern.octopuses[row-1][column+1])
	} else if column == lastIndex && row == 0 {
		// top-right
		octopus.neighbors = append(octopus.neighbors, cavern.octopuses[row][column-1])
		octopus.neighbors = append(octopus.neighbors, cavern.octopuses[row+1][column])
		octopus.neighbors = append(octopus.neighbors, cavern.octopuses[row+1][column-1])
	} else if column == lastIndex && row == lastIndex {
		// bottom-right
		octopus.neighbors = append(octopus.neighbors, cavern.octopuses[row-1][column])
		octopus.neighbors = append(octopus.neighbors, cavern.octopuses[row][column-1])
		octopus.neighbors = append(octopus.neighbors, cavern.octopuses[row-1][column-1])
	} else if column == 0 {
		// edge-left
		octopus.neighbors = append(octopus.neighbors, cavern.octopuses[row-1][column])
		octopus.neighbors = append(octopus.neighbors, cavern.octopuses[row-1][column+1])
		octopus.neighbors = append(octopus.neighbors, cavern.octopuses[row][column+1])
		octopus.neighbors = append(octopus.neighbors, cavern.octopuses[row+1][column+1])
		octopus.neighbors = append(octopus.neighbors, cavern.octopuses[row+1][column])
	} else if column == lastIndex {
		// edge-right
		octopus.neighbors = append(octopus.neighbors, cavern.octopuses[row-1][column])
		octopus.neighbors = append(octopus.neighbors, cavern.octopuses[row-1][column-1])
		octopus.neighbors = append(octopus.neighbors, cavern.octopuses[row][column-1])
		octopus.neighbors = append(octopus.neighbors, cavern.octopuses[row+1][column-1])
		octopus.neighbors = append(octopus.neighbors, cavern.octopuses[row+1][column])
	} else if row == 0 {
		// edge-top
		octopus.neighbors = append(octopus.neighbors, cavern.octopuses[row][column-1])
		octopus.neighbors = append(octopus.neighbors, cavern.octopuses[row][column+1])
		octopus.neighbors = append(octopus.neighbors, cavern.octopuses[row+1][column-1])
		octopus.neighbors = append(octopus.neighbors, cavern.octopuses[row+1][column])
		octopus.neighbors = append(octopus.neighbors, cavern.octopuses[row+1][column+1])
	} else if row == lastIndex {
		// edge-bottom
		octopus.neighbors = append(octopus.neighbors, cavern.octopuses[row][column-1])
		octopus.neighbors = append(octopus.neighbors, cavern.octopuses[row][column+1])
		octopus.neighbors = append(octopus.neighbors, cavern.octopuses[row-1][column-1])
		octopus.neighbors = append(octopus.neighbors, cavern.octopuses[row-1][column])
		octopus.neighbors = append(octopus.neighbors, cavern.octopuses[row-1][column+1])
	} else {
		// center
		octopus.neighbors = append(octopus.neighbors, cavern.octopuses[row][column-1])
		octopus.neighbors = append(octopus.neighbors, cavern.octopuses[row-1][column-1])
		octopus.neighbors = append(octopus.neighbors, cavern.octopuses[row-1][column])
		octopus.neighbors = append(octopus.neighbors, cavern.octopuses[row-1][column+1])
		octopus.neighbors = append(octopus.neighbors, cavern.octopuses[row][column+1])
		octopus.neighbors = append(octopus.neighbors, cavern.octopuses[row+1][column+1])
		octopus.neighbors = append(octopus.neighbors, cavern.octopuses[row+1][column])
		octopus.neighbors = append(octopus.neighbors, cavern.octopuses[row+1][column-1])
	}
}

func (octopus *Octopus) flash() {
	octopus.flashed = true
	for _, neighbor := range octopus.neighbors {
		neighbor.increaseEnergyLevel()
		if neighbor.readyToFlash && !neighbor.flashed {
			neighbor.flash()
		}
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

func (cavern *Cavern) String() string {
	red := color.New(color.FgRed).SprintFunc()
	s := ""
	for _, row := range cavern.octopuses {
		for _, octopus := range row {
			if octopus.energyLevel == 0 {
				s += fmt.Sprint(red(octopus.energyLevel))
			} else {
				s += fmt.Sprint(octopus.energyLevel)
			}
		}
		s += fmt.Sprintf("\n")
	}
	return s
}
