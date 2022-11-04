package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
)

type Cave struct {
	connectedCaves []*Cave
	name           string
	isSmall        bool
	alreadyVisited bool
}

type CaveSystem struct {
	caves         []*Cave
	start         *Cave
	end           *Cave
	exploredPaths [][]*Cave
}

func main() {
	caveSystem := parseCaveMap(openCsvFile("input.csv"))
	var path []*Cave
	var paths = caveSystem.explore(caveSystem.start, path)
	fmt.Println("caveSystem.paths :", len(paths))
}

func (caveSystem *CaveSystem) explore(cave *Cave, path []*Cave) [][]*Cave  {
	path = append(path, cave)
	if cave == caveSystem.end {
		return [][]*Cave{path}
	}
	var paths [][]*Cave
	for _, nextCave := range cave.connectedCaves {
		if !alreadyVisited(path, nextCave) {
			newpaths := caveSystem.explore(nextCave, path)
			for _, newpath := range newpaths {
				paths = append(paths, newpath)
			}
		}
	}
	return paths
}

func parseCaveMap(data [][]string) *CaveSystem {
	caveSystem := new(CaveSystem)
	for _, line := range data {
		for _, cave := range strings.Split(line[0], "-") {
			if cave == "start" && caveSystem.start == nil {
				caveSystem.start = NewCave(cave, strings.ToLower(cave) == cave)
				caveSystem.caves = append(caveSystem.caves, caveSystem.start)
			}
			if cave == "end" && caveSystem.end == nil {
				caveSystem.end = NewCave(cave, strings.ToLower(cave) == cave)
				caveSystem.caves = append(caveSystem.caves, caveSystem.end)
			}
			if caveSystem.contains(cave) == -1 {
				caveSystem.caves = append(caveSystem.caves, NewCave(cave, strings.ToLower(cave) == cave))
			}

		}
	}

	for _, line := range data {
		caves := strings.Split(line[0], "-")
		cave1 := caveSystem.getCave(caves[0])
		cave2 := caveSystem.getCave(caves[1])
		cave1.connectedCaves = append(cave1.connectedCaves, cave2)
		cave2.connectedCaves = append(cave2.connectedCaves, cave1)
	}
	return caveSystem
}

func (caveSystem *CaveSystem) contains(caveName string) int {
	for i, cave := range caveSystem.caves {
		if cave.name == caveName {
			return i
		}
	}
	return -1
}

func alreadyVisited(path []*Cave, cave *Cave) bool {
	if !cave.isSmall {
		return false
	}
	i := 0
	for _, c := range path {
		if c == cave {
			i++
		}
	}
	if i >= 1 {
		return true
	}
	return false
}

func (caveSystem *CaveSystem) getCave(caveName string) *Cave {
	for _, cave := range caveSystem.caves {
		if cave.name == caveName {
			return cave
		}
	}
	return nil
}

func NewCave(name string, isSmall bool) *Cave {
	cave := new(Cave)
	cave.name = name
	cave.isSmall = isSmall
	return cave
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

func (caveSystem *CaveSystem) printPaths() {
	for _, path := range caveSystem.exploredPaths {
		for _, cave := range path {
			fmt.Printf("%s ", cave.name)
		}
		fmt.Println()
	}
}
