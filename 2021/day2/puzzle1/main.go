package main

import (
	"fmt"
	"log"
	"os"
	"encoding/csv"
	"strconv"
)

type position struct {
	value int
}

type depth struct {
	value int
}

func (p *position) forward(n int) {
	p.value += n
	fmt.Printf("forward %d, current horizontal position: %d\n", n, p.value)
}

func (d *depth) down(n int) {
	d.value += n
	fmt.Printf("down %d, current depth: %d\n", n, d.value)
}

func (d *depth) up(n int) {
	d.value -= n
	fmt.Printf("up %d, current depth: %d\n", n, d.value)
}

func readCSVFile(path string) [][]string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Can not open csv file %s; err: %v\n", path, err)
	}

	reader := csv.NewReader(file)
	data, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("Unable to read data of csv file %s; err: %v\n", path, err)
	}
	return data
}

func parseToInteger(data string) (i int) {
	i, err := strconv.Atoi(data)
	if err != nil {
		log.Fatal("Can't parse value %v to integer\n", data)
	}
	return i
}

func main() {
	commands := readCSVFile("input.csv")

	p := &position{
		value: 0,
	}

	d := &depth{
		value: 0,
	}

	for _, command := range commands {
		switch command[0] {
			case "forward":
				p.forward(parseToInteger(command[1]))
			case "down":
				d.down(parseToInteger(command[1]))
			case "up":
				d.up(parseToInteger(command[1]))
			default:
				fmt.Printf("Unknown Command: %s\n", command[0])
		}
	}

	fmt.Printf("position %d * depth %d = %d\n", p.value, d.value, p.value * d.value)
}
