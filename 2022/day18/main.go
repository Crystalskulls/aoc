package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Cube struct {
	x int
	y int
	z int
}

func main() {
	cubes := createCubes()
	covered := 0
	for _, cubeA := range cubes {
		for _, cubeB := range cubes {
			if cubeA == cubeB {
				continue
			}
			if isAdjencentCube(cubeA, cubeB) {
				covered++
			}
		}
	}
	fmt.Println(len(cubes)*6 - covered)
}

func isAdjencentCube(cubeA, cubeB *Cube) (adjecent bool) {
	deltaX := int(math.Abs(float64(cubeA.x - cubeB.x)))
	deltaY := int(math.Abs(float64(cubeA.y - cubeB.y)))
	deltaZ := int(math.Abs(float64(cubeA.z - cubeB.z)))
	sum := deltaX + deltaY + deltaZ
	return sum == 1
}

func createCubes() []*Cube {
	file, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(file), "\n")
	cubes := make([]*Cube, len(lines))
	for i, line := range lines {
		coordinates := strings.Split(line, ",")
		x, _ := strconv.Atoi(coordinates[0])
		y, _ := strconv.Atoi(coordinates[1])
		z, _ := strconv.Atoi(coordinates[2])
		cubes[i] = newCube(x, y, z)
	}
	return cubes
}

func newCube(x, y, z int) *Cube {
	cube := new(Cube)
	cube.x = x
	cube.y = y
	cube.z = z
	return cube
}
