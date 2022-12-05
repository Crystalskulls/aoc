package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"

	"github.com/Crystalskulls/aoc/2022/strconv"
	"github.com/golang-collections/collections/stack"
)

func main() {
	stackData, _ := os.ReadFile("stacks.txt")
	instructionData, _ := os.ReadFile("instructions.txt")
	stacksPartOne := parseStack(string(stackData))
	stacksPartTwo := parseStack(string(stackData))
	instructions := parseInstructions(string(instructionData))

	for _, instruction := range instructions {
		tmp := make([]string, instruction[0])
		for i := 1; i <= instruction[0]; i++ {
			tmp[instruction[0]-i] = stacksPartTwo[instruction[1]].Pop().(string)
			stacksPartOne[instruction[2]].Push(stacksPartOne[instruction[1]].Pop())
		}
		for _, t := range tmp {
			stacksPartTwo[instruction[2]].Push(t)
		}
	}

	var partOne, partTwo string
	for i := 1; i <= len(stacksPartTwo); i++ {
		partOne += stacksPartOne[i].Peek().(string)
		partTwo += stacksPartTwo[i].Peek().(string)
	}
	fmt.Printf("Part One: %s - Part Two: %s\n", partOne, partTwo)
}

func parseStack(data string) map[int]*stack.Stack {
	lines := strings.Split(data, "\n")
	stacks := make(map[int]*stack.Stack)

	for i := len(lines) - 1; i >= 0; i-- {
		for j, r := range lines[i] {
			if unicode.IsLetter(r) {
				k := (j / 4) + (j % 4)
				if stacks[k] == nil {
					stacks[k] = stack.New()
				}
				stacks[k].Push(string(r))
			}
		}
	}
	return stacks
}

func parseInstructions(data string) [][]int {
	lines := strings.Split(data, "\n")
	instructions := make([][]int, len(lines))
	for i, line := range lines {
		fields := strings.Fields(line)
		instructions[i] = []int{
			strconv.ParseToInt(fields[1]), strconv.ParseToInt(fields[3]), strconv.ParseToInt(fields[5]),
		}
	}
	return instructions
}
