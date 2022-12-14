package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	blockedCoordinates, maxY := scanRockStructure()
	fmt.Printf("PartOne: %d\n", simulateFallingSand(blockedCoordinates, maxY, false))
	blockedCoordinates, maxY = scanRockStructure()
	fmt.Printf("PartTwo: %d\n", simulateFallingSand(blockedCoordinates, maxY+2, true))
}

func simulateFallingSand(blockedCoordinates map[string]struct{}, maxY int, thereIsFloor bool) (n int) {
	for {
		sand := []int{500, 0}
		if thereIsFloor {
			sand[1] = -1
		}
		for {
			if sand[1] >= maxY && !thereIsFloor {
				return
			}
			if _, blocked := blockedCoordinates["500,0"]; blocked {
				return
			}
			if sand[1] == (maxY-1) && thereIsFloor {
				blockedCoordinates[fmt.Sprintf("%d,%d", sand[0], sand[1])] = struct{}{}
				n++
				break
			}
			if isFreeCoordinate(blockedCoordinates, sand[0], sand[1]+1) {
				// one step down
				sand[1]++
			} else if isFreeCoordinate(blockedCoordinates, sand[0]-1, sand[1]+1) {
				// one step down and to the left
				sand[0]--
				sand[1]++
			} else if isFreeCoordinate(blockedCoordinates, sand[0]+1, sand[1]+1) {
				// one step down and to the right
				sand[0]++
				sand[1]++
			} else {
				blockedCoordinates[fmt.Sprintf("%d,%d", sand[0], sand[1])] = struct{}{}
				n++
				break
			}
		}
	}
}

func isFreeCoordinate(blockedCoordinates map[string]struct{}, x, y int) bool {
	_, blocked := blockedCoordinates[fmt.Sprintf("%d,%d", x, y)]
	return !blocked
}

func scanRockStructure() (blockedCoordinates map[string]struct{}, maxY int) {
	file, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(file), "\n")
	blockedCoordinates = make(map[string]struct{})

	for _, line := range lines {
		coordinates := strings.Split(line, " -> ")
		setBlockedCoordinates(blockedCoordinates, coordinates)
	}
	for k := range blockedCoordinates {
		y, _ := strconv.Atoi(strings.Split(k, ",")[1])
		maxY = int(math.Max(float64(maxY), float64(y)))
	}
	return
}

func setBlockedCoordinates(blockedCoordinates map[string]struct{}, coordinates []string) {
	for i := 0; i < len(coordinates)-1; i++ {
		coordinateA := strings.Split(coordinates[i], ",")
		coordinateB := strings.Split(coordinates[i+1], ",")
		ax, _ := strconv.Atoi(coordinateA[0])
		ay, _ := strconv.Atoi(coordinateA[1])
		bx, _ := strconv.Atoi(coordinateB[0])
		by, _ := strconv.Atoi(coordinateB[1])
		blockedCoordinates[fmt.Sprintf("%d,%d", ax, ay)] = struct{}{}
		blockedCoordinates[fmt.Sprintf("%d,%d", bx, by)] = struct{}{}

		deltaX := int(math.Min(1, math.Abs(float64(bx)-float64(ax))))
		deltaY := int(math.Min(1, math.Abs(float64(by)-float64(ay))))

		minX := int(math.Min(float64(ax), float64(bx)))
		minY := int(math.Min(float64(ay), float64(by)))

		maxX := int(math.Max(float64(ax), float64(bx)))
		maxY := int(math.Max(float64(ay), float64(by)))

		for minX != maxX || minY != maxY {
			minX += deltaX
			minY += deltaY
			blockedCoordinates[fmt.Sprintf("%d,%d", minX, minY)] = struct{}{}
		}
	}
}
