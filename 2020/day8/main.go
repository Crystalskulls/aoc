package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Instruction struct {
	operation string
	argument  int
	executed  bool
}

func main() {
	instructions := readCSV("input.csv")
	fmt.Println(boot(instructions))
}

func boot(instructions []*Instruction) (accumulator int) {
	instructionPointer := 0
	for {
		if instructions[instructionPointer].executed {
			return
		}
		switch instructions[instructionPointer].operation {
		case "acc":
			accumulator += instructions[instructionPointer].argument
			instructions[instructionPointer].executed = true
			instructionPointer++
		case "jmp":
			instructions[instructionPointer].executed = true
			instructionPointer += instructions[instructionPointer].argument
		case "nop":
			instructions[instructionPointer].executed = true
			instructionPointer++
		}
	}
}

func readCSV(path string) []*Instruction {
	instructions := make([]*Instruction, 0)
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Can't open file %s\n", path)
	}
	reader := csv.NewReader(file)
	reader.Comma = ' '
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("Can't read file %s\n", path)
	}
	for _, record := range records {
		i := new(Instruction)
		i.operation = record[0]
		i.argument, err = strconv.Atoi(record[1])
		if err != nil {
			log.Fatal("Can't Atoi ", record[1])
		}
		instructions = append(instructions, i)
	}
	return instructions
}
