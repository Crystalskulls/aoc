package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	var puzzle int
	flag.IntVar(&puzzle, "puzzle", 1, "select puzzle '1' or '2'")
	flag.Parse()

	numbers := readCSV("./input.csv")
	switch puzzle {
	case 1:
		puzzle1(numbers)
	case 2:
		puzzle2(numbers)
	}
}

func puzzle1(numbers []int) {
	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(numbers); j++ {
			if i == j {
				continue
			}
			if (numbers[i] + numbers[j]) == 2020 {
				fmt.Printf("Result puzzle1: %d\n", numbers[i]*numbers[j])
				os.Exit(0)
			}
		}
	}
}

func puzzle2(numbers []int) {
	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(numbers); j++ {
			if i == j {
				continue
			}
			for k := 0; k < len(numbers); k++ {
				if k == j {
					continue
				}
				if (numbers[i] + numbers[j] + numbers[k]) == 2020 {
					fmt.Printf("Result puzzle2: %d\n", numbers[i]*numbers[j]*numbers[k])
					os.Exit(0)
				}
			}
		}
	}
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
	k := make([]int, len(s))
	for i, v := range s {
		j, err := strconv.Atoi(v[0])
		if err != nil {
			log.Fatalf("Can't convert string %s to integer\n", v)
		}
		k[i] = j
	}
	return k
}
