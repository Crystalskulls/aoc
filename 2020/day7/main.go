package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	records := readCSV("input.txt")
	rules := parseToRules(records)
	fmt.Println(rules)
	fmt.Println()

	bagsToInspect := findBags("shiny gold", rules)
	result := make(map[string]struct{}, 0)
	i := 0
	var bag string
	for {
		i++
		bag, bagsToInspect = bagsToInspect[0], bagsToInspect[1:]
		result[bag] = struct{}{}
		bagsToInspect = append(bagsToInspect, findBags(bag, rules)...)
		if len(bagsToInspect) == 0 {
			break
		}
	}
	fmt.Println(len(result))
}

func findBags(bag string, rules map[string][]string) []string {
	bagsToInspect := make([]string, 0)
	for k, bagList := range rules {
		if strings.Contains(strings.Join(bagList, " "), bag) {
			bagsToInspect = append(bagsToInspect, k)
		}
	}
	return bagsToInspect
}

func parseToRules(records []string) map[string][]string {
	bagMap := make(map[string][]string)
	for _, record := range records {
		words := strings.Split(record, " ")
		bags := make([]string, 0)
		for i := range words {
			if i < 3 {
				continue
			}
			if i%2 == 0 && i%4 != 0 {
				bags = append(bags, strings.ReplaceAll(strings.ReplaceAll(strings.Join(words[i-1:i+1], " "), ",", ""), ".", ""))
			}
		}
		bagMap[strings.Join(words[0:2], " ")] = bags
	}
	return bagMap
}

func readCSV(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Can't open file %s\n", path)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}
