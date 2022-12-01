package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	busIDs := readCSV("./input.csv")
	earliestBus, wait := findEarliestBus(busIDs, 1013728)
	fmt.Println(earliestBus * wait)
}

func findEarliestBus(busIDs []int, timestamp int) (earliestBus, wait int) {
	wait = -1
	for _, busID := range busIDs {
		k := timestamp / busID
		d := k*busID + busID
		w := d - timestamp
		if wait == -1 || w < wait {
			wait = w
			earliestBus = busID
		}
	}
	return
}

func readCSV(path string) []int {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("PathError occured. Can't open file %s - error: %v\n", path, err)
	}
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("Can't read file %s - error: %v\n", path, err)
	}
	return parseToInt(records)
}

func parseToInt(s [][]string) []int {
	k := make([]int, 0)
	for _, v := range s[0] {
		j, err := strconv.Atoi(v)
		if err != nil {
			continue
		}
		k = append(k, j)
	}
	return k
}
