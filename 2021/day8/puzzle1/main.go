package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {
	data := openCsvFile("input.csv")
	i := 0

	for _, row := range data {
		for _, signalGroup := range row[len(row)-4:len(row)] {
			length := len(signalGroup)
			if length == 2 || length == 3 || length == 4 || length == 7 {
				i++
			}
		}
	}
	fmt.Println("i: ", i)
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
