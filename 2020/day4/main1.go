package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Passport map[string]string

func main() {
	lines := readCSV("./input.txt")
	passports := parseToPassports(lines)
	validPassports := make([]Passport, 0)

	mandatoryFields := []string{
		"byr",
		"iyr",
		"eyr",
		"hgt",
		"hcl",
		"ecl",
		"pid",
	}

	for _, passport := range passports {
		valid := true
		for _, field := range mandatoryFields {
			if _, ok := passport[field]; !ok {
				valid = false
			}
		}
		if valid {
			validPassports = append(validPassports, passport)
		}
	}

	fmt.Printf("Valid Passports: %v\n", len(validPassports))
}

func readCSV(path string) []string {
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		log.Fatalf("Can't open file %s\n", path)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split((bufio.ScanLines))
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func parseToPassports(lines []string) []Passport {
	passports := make([]Passport, 0)
	passport := make(Passport)
	for _, line := range lines {
		if len(line) == 0 {
			passports = append(passports, passport)
			passport = make(Passport)
			continue
		}
		pairs := strings.Split(line, " ")
		for _, pair := range pairs {
			tmp := strings.Split(pair, ":")
			passport[tmp[0]] = tmp[1]
		}
	}
	passports = append(passports, passport)
	return passports
}
