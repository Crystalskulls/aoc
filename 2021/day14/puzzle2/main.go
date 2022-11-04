package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Polymer struct {
	rules map[string]string
	template string
	value string
}

func main() {
	polymer := NewPolymer(openCsvFile("pair_insertion_rules.csv"), openCsvFile("polymer_template.csv")[0][0])
	i := 12
	polymer.grow(i)
}

func (polymer *Polymer) grow(steps int) {
	hashmap := make(map[string]string)
	_ = calcResult(polymer.template)
	for i:=0; i<steps; i++ {
		polymer.value = polymer.template
		polymer.template = ""
		for j:=0; j<len(polymer.value)-1; j++ {
			pair := string(polymer.value[j]) + string(polymer.value[j+1])
			if v, ok := hashmap[pair]; ok {
				polymer.template
			}
			polymer.template += string(pair[0]) + polymer.rules[pair]
			if j == len(polymer.value)-2 {
				polymer.template += string(pair[1])
			}
		}
		_ = calcResult(polymer.template)
	}
	polymer.value = polymer.template
}

func NewPolymer(rules [][]string, template string) *Polymer {
	polymer := new(Polymer)
	polymer.rules = make(map[string]string)
	polymer.template = template
	for _, pair := range rules {
		polymer.rules[pair[0]] = pair[1]
	}
	return polymer
}

func calcResult(polymer string) int {
	groupedByRunes := make(map[string]int)
	for _, v := range polymer {
		groupedByRunes[string(v)]++
	}
	fmt.Println(groupedByRunes)
	min, max := 0, 0
	for _, v := range groupedByRunes {
		if min == 0 && max == 0 {
			min = v
			max = v
			continue
		}
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	return max-min
}

func openCsvFile(path string) [][]string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Unable to open csv file %s; err: %v\n", path, err)
	}
	reader := csv.NewReader(file)
	data, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("Unable to read csv file %s; err: %v\n", path, err)
	}
	return data
}

func parseToInt(n string) int {
	i, err := strconv.ParseInt(n, 10, 64)
	if err != nil {
		log.Fatalf("Unable to parse number %s to an Integer; err: %v\n", n)
	}
	return int(i)
}
