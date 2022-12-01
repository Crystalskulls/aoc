package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

type instruction struct {
	action rune
	value  int
}

type ship struct {
	x int
	y int
	w *waypoint
}

type waypoint struct {
	x int
	y int
}

func main() {
	data := readCSV("./input.csv")
	instructions := createInstructions(data)
	s := new(ship)
	w := new(waypoint)
	w.x = 10
	w.y = 1
	s.w = w
	fmt.Println(s.navigate(instructions))
}

func (s *ship) navigate(instructions []*instruction) (md int) {
	for _, inst := range instructions {
		switch inst.action {
		case 'N':
			s.w.y += inst.value
		case 'S':
			s.w.y -= inst.value
		case 'E':
			s.w.x += inst.value
		case 'W':
			s.w.x -= inst.value
		case 'L':
			s.turn('L', inst.value)
		case 'R':
			s.turn('R', inst.value)
		case 'F':
			s.forward(inst.value)
		}
	}
	return int(math.Abs(float64(s.x)) + math.Abs(float64(s.y)))
}

func (s *ship) turn(c rune, d int) {
	r := d / 90
	for z := r; z > 0; z = d / 90 {
		d -= 90
		var xt, yt int
		xt = s.w.x
		yt = s.w.y

		if c == 'R' {
			s.w.y = -xt
			s.w.x = yt
		} else {
			s.w.y = xt
			s.w.x = -yt
		}
	}
}

func (s *ship) forward(v int) {
	x := s.w.x * v
	y := s.w.y * v
	s.x += x
	s.y += y
}

func (inst *instruction) String() string {
	return fmt.Sprintf("%v, %d", inst.action, inst.value)
}

func createInstructions(data [][]string) (instructions []*instruction) {
	for _, d := range data {
		inst := new(instruction)
		inst.action = rune(d[0][0])
		inst.value, _ = strconv.Atoi(d[0][1:])
		instructions = append(instructions, inst)
	}
	return
}

func readCSV(path string) [][]string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Can't open file %s\n", path)
	}
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("Can't read file %s\n", path)
	}
	return records
}
