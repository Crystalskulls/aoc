package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Sensor struct {
	x                 int
	y                 int
	closestBeacon     *Beacon
	manhattanDistance int
	borderCoordinates [][]int
}

type Beacon struct {
	x int
	y int
}

func main() {
	sensors := parseSensorData()
	outerCoordinates := make(map[string]struct{})
	for _, sensor := range sensors {
		sensor.setBorderCoordinates()
		sensor.setOuterCoordinates(outerCoordinates)
	}
	/*sensors[3].setBorderCoordinates()
	fmt.Println(sensors)
	fmt.Println(sensors[3].borderCoordinates)*/
	fmt.Println(findBeacon(sensors, outerCoordinates))

}

func findBeacon(sensors []*Sensor, outerCoordinates map[string]struct{}) int {
	for _, sensor := range sensors {
		for _, borderCoordinate := range sensor.borderCoordinates {
			found := true
			for _, otherSensor := range sensors {
				minY := otherSensor.y - otherSensor.manhattanDistance
				maxY := otherSensor.y + otherSensor.manhattanDistance
				if minY > borderCoordinate[1] || maxY < borderCoordinate[1] {
					continue
				}
				deltaX := int(math.Abs(math.Abs((float64(otherSensor.y - borderCoordinate[1]))) - float64(otherSensor.manhattanDistance)))
				minX := otherSensor.x - deltaX
				maxX := otherSensor.x + deltaX
				if borderCoordinate[0] >= minX && borderCoordinate[0] <= maxX {
					found = false
					break
				}
			}
			if found {
				x := borderCoordinate[0]
				y := borderCoordinate[1]
				if surounded(x, y, outerCoordinates) {
					return x*4000000 + y
				}
			}
		}
	}
	return -1
}

func surounded(x, y int, outerCoordinates map[string]struct{}) bool {
	if _, ok := outerCoordinates[fmt.Sprintf("%d,%d", x, y-1)]; !ok {
		return false
	}
	if _, ok := outerCoordinates[fmt.Sprintf("%d,%d", x, y+1)]; !ok {
		return false
	}
	if _, ok := outerCoordinates[fmt.Sprintf("%d,%d", x-1, y)]; !ok {
		return false
	}
	if _, ok := outerCoordinates[fmt.Sprintf("%d,%d", x+1, y)]; !ok {
		return false
	}
	return true
}

func parseSensorData() []*Sensor {
	file, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(file), "\n")
	sensors := make([]*Sensor, len(lines))

	for i, line := range lines {
		coordinates := strings.Split(strings.ReplaceAll(strings.ReplaceAll(line, "Sensor at ", ""), ": closest beacon is at ", ", "), ", ")
		beacon := newBeacon(coordinates)
		sensors[i] = newSensor(coordinates, beacon)
	}
	return sensors
}

func newBeacon(coordinates []string) *Beacon {
	beacon := new(Beacon)
	beacon.x, _ = strconv.Atoi(strings.Split(coordinates[2], "=")[1])
	beacon.y, _ = strconv.Atoi(strings.Split(coordinates[3], "=")[1])
	return beacon
}

func newSensor(coordinates []string, beacon *Beacon) *Sensor {
	sensor := new(Sensor)
	sensor.closestBeacon = beacon
	sensor.x, _ = strconv.Atoi(strings.Split(coordinates[0], "=")[1])
	sensor.y, _ = strconv.Atoi(strings.Split(coordinates[1], "=")[1])
	sensor.manhattanDistance = int(math.Abs(float64(sensor.x-sensor.closestBeacon.x))) + int(math.Abs(float64(sensor.y-sensor.closestBeacon.y)))
	return sensor
}

func (sensor *Sensor) String() string {
	return fmt.Sprintf("(%d,%d) -> (%d,%d); distance: %d\n", sensor.x, sensor.y, sensor.closestBeacon.x, sensor.closestBeacon.y, sensor.manhattanDistance)
}

func (sensor *Sensor) setBorderCoordinates() {
	sensor.borderCoordinates = make([][]int, 0)
	for y := sensor.y - (sensor.manhattanDistance + 1); y <= sensor.y+(sensor.manhattanDistance+1); y++ {
		deltaX := int(math.Abs(math.Abs((float64(sensor.y - y))) - float64(sensor.manhattanDistance+1)))
		minX := sensor.x - deltaX
		maxX := sensor.x + deltaX
		if minX == maxX {
			sensor.borderCoordinates = append(sensor.borderCoordinates, []int{minX, y})
		} else {
			sensor.borderCoordinates = append(sensor.borderCoordinates, []int{minX, y})
			sensor.borderCoordinates = append(sensor.borderCoordinates, []int{maxX, y})
		}

	}
}

func (sensor *Sensor) setOuterCoordinates(outerCoordinates map[string]struct{}) {
	for y := sensor.y - sensor.manhattanDistance; y <= sensor.y+sensor.manhattanDistance; y++ {
		deltaX := int(math.Abs(math.Abs((float64(sensor.y - y))) - float64(sensor.manhattanDistance)))
		minX := sensor.x - deltaX
		maxX := sensor.x + deltaX
		outerCoordinates[fmt.Sprintf("%d,%d", minX, y)] = struct{}{}
		outerCoordinates[fmt.Sprintf("%d,%d", maxX, y)] = struct{}{}
	}
}
