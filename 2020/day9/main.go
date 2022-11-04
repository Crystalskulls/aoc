package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	numbers := readCSV("input.csv")
	fmt.Println(findFirstInvalidNumber(numbers))
}

func findFirstInvalidNumber(numbers []int) (invalidNumber int) {
	preamble := 25
	for i := preamble; i < len(numbers); i++ {
		if _, ok := calcValidNumbers(numbers[i-preamble : i])[numbers[i]]; !ok {
			invalidNumber = numbers[i]
			break
		}
	}
	return invalidNumber
}

func calcValidNumbers(numbers []int) map[int]struct{} {
	validNumbers := make(map[int]struct{}, 0)
	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers); j++ {
			validNumbers[numbers[i]+numbers[j]] = struct{}{}
		}
	}
	return validNumbers
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
