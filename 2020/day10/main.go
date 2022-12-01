package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	adapters := readCSV("./input.csv")
	adapters = append(adapters, 0)
	sort.Ints(adapters)
	adapters = append(adapters, adapters[len(adapters)-1]+3)
	fmt.Println(adapters)
	joltDiff := make(map[int]int)
	for i := len(adapters) - 1; i > 0; i-- {
		joltDiff[adapters[i]-adapters[i-1]]++
	}
	fmt.Println(joltDiff)
	fmt.Println(joltDiff[1] * joltDiff[3])

	for i := 2; i < len(adapters)-2; i++ {
		s := (adapters[i] - adapters[i-1]) + (adapters[i+1] - adapters[i])
		if s == 2 {
			joltDiff[2]++
		}
	}
	fmt.Println(joltDiff[2])
}

func readCSV(path string) []int {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Can't open file %s\n", path)
	}
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("Can't read file %s\n", path)
	}
	data := make([]int, 0)
	for _, r := range records {
		i, _ := strconv.Atoi(r[0])
		data = append(data, i)
	}
	return data
}
