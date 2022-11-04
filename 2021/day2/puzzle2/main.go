package main

import (
	"fmt"
	"log"
	"os"
	"encoding/csv"
	"strconv"
)

type submarine struct {
	horizontalPosition int
	depth int
	aim int
}

func (sub *submarine) forward(n int) {
	sub.horizontalPosition += n
	sub.depth += sub.aim * n
	fmt.Printf("forward %d, current horizontal position: %d; current aim: %d\n", n, sub.horizontalPosition, sub.aim * n)
}

func (sub *submarine) down(n int) {
	sub.aim += n
	fmt.Printf("down %d, current aim: %d\n", n, sub.aim)
}

func (sub *submarine) up(n int) {
	sub.aim -= n
	fmt.Printf("up %d, current aim: %d\n", n, sub.aim)
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

	sub := &submarine{
		horizontalPosition: 0,
		depth: 0,
		aim: 0,
	}

	for _, command := range commands {
		switch command[0] {
			case "forward":
				sub.forward(parseToInteger(command[1]))
			case "down":
				sub.down(parseToInteger(command[1]))
			case "up":
				sub.up(parseToInteger(command[1]))
			default:
				fmt.Printf("Unknown Command: %s\n", command[0])
		}
	}

	fmt.Printf("position %d * depth %d = %d\n", sub.horizontalPosition, sub.depth, sub.horizontalPosition * sub.depth)
}
