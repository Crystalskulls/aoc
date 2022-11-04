package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	records := readCSV("input.txt")
	s := 0
	questions := make(map[rune]struct{})
	for _, record := range records {
		if len(record) == 0 {
			s += len(questions)
			questions = make(map[rune]struct{})
		}
		for _, char := range record {
			questions[char] = struct{}{}
		}
	}
	s += len(questions)
	fmt.Printf("puzzle1: %d\n", s)

	questions2 := make(map[rune]int)
	s = 0
	c := 0
	for _, record := range records {
		if len(record) == 0 {
			for _, v := range questions2 {
				if v == c {
					s++
				}
			}
			c = 0
			questions2 = make(map[rune]int)
			continue
		}
		c++
		for _, char := range record {
			questions2[char]++
		}
	}
	for _, v := range questions2 {
		if v == c {
			s++
		}
	}
	fmt.Printf("puzzle1: %d\n", s)
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
