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
	numbers := readCSV("input.csv")
	invalidNumber := findFirstInvalidNumber(numbers)
	fmt.Println(invalidNumber)
	fmt.Println(findContiguousSet(numbers, invalidNumber))
}

func findContiguousSet(numbers []int, invalidNumber int) (encryptionWeakness int) {
	set := make([]int, 0)
	startIndex := 0
	for {
		for i := startIndex; i < len(numbers); i++ {
			set = append(set, numbers[i])
			s := sum(set)
			if s == invalidNumber && len(set) > 1 {
				return calcWeakness(set)
			}
			if s > invalidNumber {
				startIndex++
				set = make([]int, 0)
				break
			}
		}
	}
}

func sum(numbers []int) (sum int) {
	for _, number := range numbers {
		sum += number
	}
	return
}

func calcWeakness(numbers []int) (encryptionWeakness int) {
	sort.Ints(numbers)
	return numbers[0] + numbers[len(numbers)-1]
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
