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
	direction rune
	x         int
	y         int
}

func main() {
	data := readCSV("./input.csv")
	instructions := createInstructions(data)
	s := new(ship)
	s.direction = 'E'
	fmt.Println(s.navigate(instructions))
}

func (s *ship) navigate(instructions []*instruction) (md int) {
	for _, inst := range instructions {
		switch inst.action {
		case 'N':
			s.y += inst.value
		case 'S':
			s.y -= inst.value
		case 'E':
			s.x += inst.value
		case 'W':
			s.x -= inst.value
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
	p := []rune{
		'N', 'E', 'S', 'W',
	}
	var si int
	for i, r := range p {
		if r == s.direction {
			si = i
		}
	}
	r := d / 90
	for z := r; z > 0; z = d / 90 {
		d -= 90
		if c == 'L' {
			si -= 1
			if si < 0 {
				si = len(p) - 1
			}
		}
		if c == 'R' {
			si += 1
			if si > len(p)-1 {
				si = 0
			}
		}
	}
	s.direction = p[si]
}

func (s *ship) forward(v int) {
	switch s.direction {
	case 'N':
		s.y += v
	case 'S':
		s.y -= v
	case 'E':
		s.x += v
	case 'W':
		s.x -= v
	}
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
