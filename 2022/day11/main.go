package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"gopkg.in/Knetic/govaluate.v3"
)

type Monkey struct {
	Items         []int
	Operation     *govaluate.EvaluableExpression
	Test          int
	ThrowToIndex  map[bool]int
	TotalInspects int
}

func main() {
	monkeys, lcm := createMonkeysFromInput()
	fmt.Printf("PartOne: %d\n", stuffSlingingSimianShenanigans(monkeys, 20, lcm))
	monkeys, lcm = createMonkeysFromInput()
	fmt.Printf("PartTwo: %d\n", stuffSlingingSimianShenanigans(monkeys, 10000, lcm))
}

func stuffSlingingSimianShenanigans(monkeys []*Monkey, rounds, lcm int) (monkeyBusiness int) {
	for r := 1; r <= rounds; r++ {
		for i := range monkeys {
			for _, item := range monkeys[i].Items {
				monkeys[i].TotalInspects++
				worryLevel := item
				parameters := make(map[string]interface{}, 8)
				parameters["old"] = worryLevel
				result, _ := monkeys[i].Operation.Evaluate(parameters)
				d := 1
				if rounds == 20 {
					d = 3
				}
				newWorryLevel := int(result.(float64)) / d

				decision := false
				if newWorryLevel%monkeys[i].Test == 0 {
					decision = true
				}
				monkeys[monkeys[i].ThrowToIndex[decision]].Items = append(monkeys[monkeys[i].ThrowToIndex[decision]].Items, newWorryLevel%lcm)
			}
			monkeys[i].Items = make([]int, 0)
		}
	}
	totalInspects := make([]int, len(monkeys))
	for i, monkey := range monkeys {
		totalInspects[i] = monkey.TotalInspects
	}

	sort.Sort(sort.Reverse(sort.IntSlice(totalInspects)))
	return totalInspects[0] * totalInspects[1]
}

func createMonkeysFromInput() ([]*Monkey, int) {
	file, _ := os.ReadFile("input.txt")
	monkeys := strings.Split(string(file), "\n\n")
	ms := make([]*Monkey, len(monkeys))
	lcm := 1

	for i, monkeyAttributes := range monkeys {
		lines := strings.Split(monkeyAttributes, "\n")
		monkey := new(Monkey)
		startingItems := make([]int, 0)
		for _, s := range strings.Split(strings.ReplaceAll(strings.ReplaceAll(strings.TrimSpace(lines[1]), "Starting items: ", ""), ",", ""), " ") {
			i, _ := strconv.Atoi(s)
			startingItems = append(startingItems, i)
		}
		operation, _ := govaluate.NewEvaluableExpression(strings.TrimSpace(strings.Split(lines[2], "=")[1]))
		test, _ := strconv.Atoi(strings.Split(lines[3], " ")[5])
		monkey.Items = startingItems
		monkey.Operation = operation
		monkey.Test = test
		lcm *= test
		monkey.ThrowToIndex = make(map[bool]int)
		monkey.ThrowToIndex[true], _ = strconv.Atoi(strings.Split(lines[4], " ")[9])
		monkey.ThrowToIndex[false], _ = strconv.Atoi(strings.Split(lines[5], " ")[9])

		ms[i] = monkey
	}
	return ms, lcm
}
