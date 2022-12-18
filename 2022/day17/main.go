package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Rock struct {
	coordinates [][]int
}

func main() {
	queue := initQueue()
	movements := parseInput()
	fmt.Println("PartOne:", playTetris(queue, movements))
}

func playTetris(queue []*Rock, movements []rune) int {
	i := 0
	blockedCoordinates := make(map[string]struct{})
	for j := 0; j < 7; j++ {
		blockedCoordinates[fmt.Sprintf("%d,4", j)] = struct{}{}
	}
	for {
		if i == 2022 {
			break
		}
		rock := queue[0]
		queue = queue[1:]
		newRock := new(Rock)
		newRock.coordinates = make([][]int, len(rock.coordinates))
		for k := range newRock.coordinates {
			newRock.coordinates[k] = make([]int, 2)
			copy(newRock.coordinates[k], rock.coordinates[k])
		}
		queue = append(queue, newRock)
		for {
			move := movements[0]
			movements = movements[1:]
			movements = append(movements, move)
			if move == 62 {
				rock.moveRight(blockedCoordinates)
			} else {
				rock.moveLeft(blockedCoordinates)
			}
			if moveDownBlocked(rock, blockedCoordinates) {
				i++
				for _, coordinate := range rock.coordinates {
					blockedCoordinates[fmt.Sprintf("%d,%d", coordinate[0], coordinate[1])] = struct{}{}
				}
				sh := queue[0].getHigh()
				th, maxY := getTowerHeight(blockedCoordinates)
				delta := (sh + 3 + th) - (maxY)
				blockedCoordinates = updateBlockedCoordinates(blockedCoordinates, delta)
				break
			}
			rock.moveDown()
		}
	}
	th, _ := getTowerHeight(blockedCoordinates)
	return th
}

func (rock *Rock) getHigh() int {
	h := 0
	for _, cor := range rock.coordinates {
		if cor[1] > h {
			h = cor[1]
		}
	}
	return h + 1
}

func print(blockedCoordinates map[string]struct{}) {
	for i := 0; i <= 16; i++ {
		for j := 0; j <= 6; j++ {
			if _, blocked := blockedCoordinates[fmt.Sprintf("%d,%d", j, i)]; blocked {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func getTowerHeight(blockedCoordinates map[string]struct{}) (int, int) {
	var minY, maxY *int
	for k := range blockedCoordinates {
		y, _ := strconv.Atoi(strings.Split(k, ",")[1])
		if minY == nil && maxY == nil {
			minY = new(int)
			maxY = new(int)
			*minY = y
			*maxY = y
			continue
		}
		if y < *minY {
			*minY = y
		}
		if y > *maxY {
			*maxY = y
		}
	}
	return *maxY - *minY, *maxY
}

func updateBlockedCoordinates(blockedCoordinates map[string]struct{}, delta int) map[string]struct{} {
	newBlockedCoordinates := make(map[string]struct{})
	for k := range blockedCoordinates {
		coordinates := strings.Split(k, ",")
		x := coordinates[0]
		y, _ := strconv.Atoi(coordinates[1])
		newBlockedCoordinates[fmt.Sprintf("%s,%d", x, y+delta)] = struct{}{}
	}
	return newBlockedCoordinates
}

func (rock *Rock) moveRight(blockedCoordinates map[string]struct{}) {
	for _, coordinate := range rock.coordinates {
		if _, blocked := blockedCoordinates[fmt.Sprintf("%d,%d", coordinate[0]+1, coordinate[1])]; blocked {
			return
		}

	}
	var maxX *int
	for _, coordinate := range rock.coordinates {
		if maxX == nil {
			maxX = new(int)
			*maxX = coordinate[0]
		}
		if coordinate[0] > *maxX {
			*maxX = coordinate[0]
		}
	}
	if *maxX == 6 {
		return
	}
	for _, coordinate := range rock.coordinates {
		coordinate[0]++
	}
}

func (rock *Rock) moveLeft(blockedCoordinates map[string]struct{}) {
	for _, coordinate := range rock.coordinates {
		if _, blocked := blockedCoordinates[fmt.Sprintf("%d,%d", coordinate[0]-1, coordinate[1])]; blocked {
			return
		}

	}
	var minX *int
	for _, coordinate := range rock.coordinates {
		if minX == nil {
			minX = new(int)
			*minX = coordinate[0]
		}
		if coordinate[0] < *minX {
			*minX = coordinate[0]
		}
	}
	if *minX == 0 {
		return
	}
	for _, coordinate := range rock.coordinates {
		coordinate[0]--
	}
}

func (rock *Rock) moveDown() {
	for _, coordinate := range rock.coordinates {
		coordinate[1]++
	}
}

func moveDownBlocked(rock *Rock, blockedCoordinates map[string]struct{}) bool {
	for _, coordinate := range rock.coordinates {
		if _, blocked := blockedCoordinates[fmt.Sprintf("%d,%d", coordinate[0], coordinate[1]+1)]; blocked {
			return true
		}
	}
	return false
}

func parseInput() []rune {
	file, _ := os.ReadFile("input.txt")
	movements := make([]rune, 0)
	for _, r := range string(file) {
		movements = append(movements, r)
	}
	return movements
}

func initQueue() []*Rock {
	queue := make([]*Rock, 0)
	rock := new(Rock)
	rock.coordinates = [][]int{
		[]int{2, 0},
		[]int{3, 0},
		[]int{4, 0},
		[]int{5, 0},
	}
	queue = append(queue, rock)
	rock = new(Rock)
	rock.coordinates = [][]int{
		[]int{3, 0},
		[]int{3, 1},
		[]int{3, 2},
		[]int{2, 1},
		[]int{3, 1},
		[]int{4, 1},
	}
	queue = append(queue, rock)
	rock = new(Rock)
	rock.coordinates = [][]int{
		[]int{4, 0},
		[]int{4, 1},
		[]int{4, 2},
		[]int{2, 2},
		[]int{3, 2},
	}
	queue = append(queue, rock)
	rock = new(Rock)
	rock.coordinates = [][]int{
		[]int{2, 0},
		[]int{2, 1},
		[]int{2, 2},
		[]int{2, 3},
	}
	queue = append(queue, rock)
	rock = new(Rock)
	rock.coordinates = [][]int{
		[]int{2, 0},
		[]int{3, 0},
		[]int{2, 1},
		[]int{3, 1},
	}
	queue = append(queue, rock)
	return queue
}

func (rock *Rock) String() string {
	return fmt.Sprintln(rock.coordinates)
}
