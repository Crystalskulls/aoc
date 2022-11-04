package main

import (
	"fmt"
	"log"
	"os"
	"encoding/csv"
	"strconv"
	"time"
)

func main() {
	start := time.Now()
	data := readCSVFile("input.csv")
	measurements := parseToInteger(data)
	counter := 0

	for i:=0; i < len(measurements)-3; i++ {
		currentMeasurement := measurements[i] + measurements[i+1] + measurements[i+2]
		nextMeasurement := measurements[i+1] + measurements[i+2] + measurements[i+3]

		if nextMeasurement > currentMeasurement {
			counter++
			fmt.Printf("%X: %d (increased)\n", i, nextMeasurement)
		} else if nextMeasurement < currentMeasurement {
			fmt.Printf("%X: %d (decreased)\n", i, nextMeasurement)
		} else {
			fmt.Printf("%X: %d (no change)\n", i, nextMeasurement)
		}
	}
	fmt.Printf("increased: %d\n", counter)
	fmt.Printf("time: %s\n", time.Since(start))
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

func parseToInteger(data [][]string) []int {
	var parsedData []int
	for _, value := range data {
		measurement, err := strconv.Atoi(value[0])
		if err != nil {
			log.Fatal("Can't parse value %v to integer\n", value[0])
		}
		parsedData = append(parsedData, measurement)
	}
	return parsedData
}
