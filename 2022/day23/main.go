package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

type Elf struct {
	x             int
	y             int
	proposeX      int
	proposeY      int
	adjacentElves map[string]*Elf
	lazy          bool
}

func main() {
	grove := scanGrove()
	plantSeedlings(grove)
}

func plantSeedlings(grove [][]*Elf) {
	directions := [][]string{
		[]string{"N", "NE", "NW"},
		[]string{"S", "SE", "SW"},
		[]string{"W", "NW", "SW"},
		[]string{"E", "NE", "SE"},
	}
	i := 1
	for {
		considerAdjecentElves(grove)
		proposeMoving(grove, directions)
		mc, expandGrove := moveElves(grove)
		grove = updateGrove(grove, expandGrove)
		directions = append(directions, directions[0])
		directions = directions[1:]
		if i == 10 {
			fmt.Println("PartOne:", countEmptyGroundTiles(grove))
		}
		//fmt.Println("mc:", mc)
		if mc == 0 {
			fmt.Println("PartTwo:", i)
			break
		}
		i++
	}
}

func updateGrove(grove [][]*Elf, expandGrove bool) [][]*Elf {
	delta := 0
	if expandGrove {
		delta = 2
	}
	newGrove := make([][]*Elf, len(grove)+delta)
	for row := range newGrove {
		newGrove[row] = make([]*Elf, len(grove[0])+delta)
	}

	for _, row := range grove {
		for _, elf := range row {
			if elf != nil {
				if expandGrove {
					elf.x++
					elf.y++
				}
				newGrove[elf.x][elf.y] = elf
			}
		}
	}
	return newGrove
}

func moveElves(grove [][]*Elf) (mc int, expandGrove bool) {
	moves := make(map[string][]*Elf)
	for _, row := range grove {
		for _, elf := range row {
			if elf == nil || elf.lazy {
				continue
			}
			move := fmt.Sprintf("%d,%d", elf.proposeX, elf.proposeY)
			if _, moveFound := moves[move]; !moveFound {
				moves[move] = make([]*Elf, 0)
			}
			moves[move] = append(moves[move], elf)
		}
	}

	for _, elves := range moves {
		if len(elves) > 1 {
			for _, elf := range elves {
				elf.lazy = true
			}
		}
	}

	maxX, maxY := 0, 0
	minX, minY := len(grove), len(grove[0])
	for _, row := range grove {
		for _, elf := range row {
			if elf == nil {
				continue
			}
			if elf.proposeX < minX {
				minX = elf.proposeX
			}
			if elf.proposeX > maxX {
				maxX = elf.proposeX
			}
			if elf.proposeY < minY {
				minY = elf.proposeY
			}
			if elf.proposeY > maxY {
				maxY = elf.proposeY
			}
			if elf.lazy {
				continue
			}
			mc++
			elf.x = elf.proposeX
			elf.y = elf.proposeY
			elf.lazy = true
		}
	}
	if minX == 0 || minY == 0 || maxX == len(grove[0])-1 || maxY == len(grove)-1 {
		expandGrove = true
	}
	return mc, expandGrove
}

func countEmptyGroundTiles(grove [][]*Elf) (emptyGroundTiles int) {
	maxX, maxY := 0, 0
	minX, minY := len(grove), len(grove[0])

	for x, row := range grove {
		for y, elf := range row {
			if elf == nil {
				continue
			}
			if x < minX {
				minX = x
			}
			if x > maxX {
				maxX = x
			}
			if y < minY {
				minY = y
			}
			if y > maxY {
				maxY = y
			}
		}
	}
	for x := minX; x <= maxX; x++ {
		for y := minY; y <= maxY; y++ {
			if grove[x][y] == nil {
				emptyGroundTiles++
			}
		}
	}
	return emptyGroundTiles
}

func proposeMoving(grove [][]*Elf, directions [][]string) {
	for _, row := range grove {
		for _, elf := range row {
			if elf == nil || elf.lazy {
				continue
			}
			hasValidDirection := false
			for _, directionSet := range directions {
				if elf.isValidDirection(directionSet) {
					hasValidDirection = true
					elf.setMovingProposal(directionSet[0])
					break
				}
			}
			if !hasValidDirection {
				elf.lazy = true
			}
		}
	}
}

func (elf *Elf) isValidDirection(directions []string) bool {
	for _, direction := range directions {
		if elf.adjacentElves[direction] != nil {
			return false
		}
	}
	return true
}

func (elf *Elf) setMovingProposal(direction string) {
	if direction == "N" {
		elf.proposeX = elf.x - 1
		elf.proposeY = elf.y
	} else if direction == "S" {
		elf.proposeX = elf.x + 1
		elf.proposeY = elf.y
	} else if direction == "E" {
		elf.proposeX = elf.x
		elf.proposeY = elf.y + 1
	} else {
		elf.proposeX = elf.x
		elf.proposeY = elf.y - 1
	}
}

func considerAdjecentElves(grove [][]*Elf) {
	for _, row := range grove {
		for _, elf := range row {
			if elf != nil {
				elf.setAdjecentElves(grove)
				for _, adjecentElf := range elf.adjacentElves {
					if adjecentElf != nil {
						elf.lazy = false
						break
					}
				}
				if elf.lazy {
					elf.proposeX = elf.x
					elf.proposeY = elf.y
				}
			}
		}
	}
}

func scanGrove() [][]*Elf {
	file, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(file), "\n")
	grove := make([][]*Elf, len(lines)+2)

	for row := range grove {
		grove[row] = make([]*Elf, len(lines[0])+2)
	}

	for row, line := range lines {
		for col, r := range line {
			if string(r) == "#" {
				grove[row+1][col+1] = newElf(row+1, col+1)
				continue
			}
		}
	}
	return grove
}

func newElf(x, y int) *Elf {
	elf := new(Elf)
	elf.x = x
	elf.y = y
	elf.lazy = true
	return elf
}

func (elf *Elf) setAdjecentElves(grove [][]*Elf) {
	adjecentElves := make(map[string]*Elf)
	rowLimit, columnLimit := float64(len(grove)-1), float64(len(grove)-1)
	for x := int(math.Max(0, float64(elf.x-1))); x <= int(math.Min(float64(elf.x+1), rowLimit)); x++ {
		for y := int(math.Max(0, float64(elf.y-1))); y <= int(math.Min(float64(elf.y+1), columnLimit)); y++ {
			if x != elf.x || y != elf.y {
				if x < elf.x && y < elf.y {
					adjecentElves["NW"] = grove[x][y]
				} else if x < elf.x && y == elf.y {
					adjecentElves["N"] = grove[x][y]
				} else if x < elf.x && y > elf.y {
					adjecentElves["NE"] = grove[x][y]
				} else if x == elf.x && y < elf.y {
					adjecentElves["W"] = grove[x][y]
				} else if x == elf.x && y > elf.y {
					adjecentElves["E"] = grove[x][y]
				} else if x > elf.x && y < elf.y {
					adjecentElves["SW"] = grove[x][y]
				} else if x > elf.x && y == elf.y {
					adjecentElves["S"] = grove[x][y]
				} else {
					adjecentElves["SE"] = grove[x][y]
				}
			}
		}
	}
	elf.adjacentElves = adjecentElves
}

func printGrove(grove [][]*Elf) {
	for _, row := range grove {
		for _, elf := range row {
			if elf == nil {
				fmt.Print(".")
				continue
			}
			fmt.Print("#")
		}
		fmt.Println()
	}
	fmt.Println()
}

func (elf *Elf) String() string {
	return fmt.Sprintf("x: %d y: %d px: %d py: %d\n", elf.x, elf.y, elf.proposeX, elf.proposeY)
}
