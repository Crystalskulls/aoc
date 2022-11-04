package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"time"
)

func main() {
	start := time.Now()
	data := openCsvFile("input.csv")
	var crabs []int
	for _, d := range data[0] {
		crabs = append(crabs, parseToInt(d))
	}
	cheapest := cheapestPosition(crabs)
	fmt.Println("fuel: ", cheapest)
	fmt.Println("time: ", time.Since(start))
}

func cheapestPosition(crabs []int) int {
	cheapest := -1
	for i := 1; i <= len(crabs); i++ {
		fuelByPosition := 0
		for j := 0; j < len(crabs); j++ {
			fuel := int(math.Abs(float64(crabs[j] - i)))
			fuelByPosition += fuel
		}
		if cheapest == -1 {
			cheapest = fuelByPosition
		} else if fuelByPosition < cheapest {
			cheapest = fuelByPosition
		}
	}
	return cheapest
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
