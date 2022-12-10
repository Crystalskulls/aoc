package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.ReadFile("input.txt")
	instructions := strings.Split(string(file), "\n")
	x, v, cycle := 1, 0, 0
	cycleStrengths := make(map[int]int)
	screenRows := make([][]string, 6)
	for i := range screenRows {
		screenRows[i] = make([]string, 40)
	}
	for _, instruction := range instructions {
		cycle++
		draw(screenRows, cycle, x)
		cycleStrengths[cycle] = x
		if strings.HasPrefix(instruction, "addx") {
			v, _ = strconv.Atoi(strings.Split(instruction, " ")[1])
			cycle++
			draw(screenRows, cycle, x)
			cycleStrengths[cycle] = x
			x += v
			continue
		} else {
			v = 0
		}
	}
	relevantCyles := []int{20, 60, 100, 140, 180, 220}
	partOne := 0
	for _, k := range relevantCyles {
		partOne += (k * cycleStrengths[k])
	}
	fmt.Printf("PartOne: %d\n", partOne)
	fmt.Println("PartTwo: ")
	for _, row := range screenRows {
		fmt.Println(row)
	}
}

func draw(screenRows [][]string, cycle, x int) {
	row := (cycle - 1) / 40
	if row > 5 {
		return
	}
	if cycle-1-(row*40) >= x-1 && cycle-1-(row*40) <= x+1 {
		screenRows[row][cycle-1-(row*40)] = "#"
		return
	}
	screenRows[row][cycle-1-(row*40)] = "."
}
