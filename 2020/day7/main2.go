package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	records := readCSV("input.txt")
	rules := parseToRules(records)
	bagsToInspect := rules["shiny gold"]
	var bag string
	result := make([]string, 0)
	result = append(result, bagsToInspect...)
	for {
		bag, bagsToInspect = bagsToInspect[0], bagsToInspect[1:]
		if rules[bag][0] != "other bags" {
			result = append(result, rules[bag]...)
			bagsToInspect = append(bagsToInspect, rules[bag]...)
		}
		if len(bagsToInspect) == 0 {
			break
		}
	}
	fmt.Println(len(result))
}

func parseToRules(records []string) map[string][]string {
	bagMap := make(map[string][]string)
	var bag string
	for _, record := range records {
		words := strings.Split(record, " ")
		bags := make([]string, 0)
		for i := range words {
			if i < 3 {
				continue
			}
			if i%2 == 0 && i%4 != 0 {
				bag = strings.ReplaceAll(strings.ReplaceAll(strings.Join(words[i-1:i+1], " "), ",", ""), ".", "")
				i, err := strconv.Atoi(words[i-2])
				if err != nil {
					// no other bags
					bags = append(bags, bag)
				} else {
					for j := 0; j < i; j++ {
						bags = append(bags, bag)
					}
				}
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
