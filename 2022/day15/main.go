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
}

type Beacon struct {
	x int
	y int
}

func main() {
	sensors, beaconCoordinates, sensorCoordinates := parseSensorData()
	blockedCoordinates := make(map[int]map[int]struct{})
	y := 2000000
	for _, sensor := range sensors {
		sensor.SetBlockedCoordinates(blockedCoordinates, y)
	}
	sum := len(blockedCoordinates[y])
	for x := range beaconCoordinates[y] {
		if _, ok := blockedCoordinates[y][x]; ok {
			sum--
		}
	}
	for x := range sensorCoordinates[y] {
		if _, ok := blockedCoordinates[y][x]; ok {
			sum--
		}
	}
	fmt.Println(sum)
}

func parseSensorData() ([]*Sensor, map[int]map[int]struct{}, map[int]map[int]struct{}) {
	file, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(file), "\n")
	sensors := make([]*Sensor, len(lines))
	beaconCoordinates := make(map[int]map[int]struct{})
	sensorCoordinates := make(map[int]map[int]struct{})

	for i, line := range lines {
		coordinates := strings.Split(strings.ReplaceAll(strings.ReplaceAll(line, "Sensor at ", ""), ": closest beacon is at ", ", "), ", ")
		beacon := newBeacon(coordinates)
		if _, ok := beaconCoordinates[beacon.y]; !ok {
			beaconCoordinates[beacon.y] = make(map[int]struct{})
		}
		beaconCoordinates[beacon.y][beacon.x] = struct{}{}
		sensors[i] = newSensor(coordinates, beacon)
		if _, ok := sensorCoordinates[sensors[i].y]; !ok {
			sensorCoordinates[sensors[i].y] = make(map[int]struct{})
		}
		sensorCoordinates[sensors[i].y][sensors[i].x] = struct{}{}
	}
	return sensors, beaconCoordinates, sensorCoordinates
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

func (sensor *Sensor) SetBlockedCoordinates(blockedCoordinates map[int]map[int]struct{}, y int) {
	if sensor.y-sensor.manhattanDistance <= y && sensor.y+sensor.manhattanDistance >= y {
		if blockedCoordinates[y] == nil {
			blockedCoordinates[y] = make(map[int]struct{})
		}
		deltaX := int(math.Abs(math.Abs((float64(sensor.y - y))) - float64(sensor.manhattanDistance)))
		minX := sensor.x - deltaX
		maxX := sensor.x + deltaX
		for x := minX; x <= maxX; x++ {
			blockedCoordinates[y][x] = struct{}{}
		}
	}
}
