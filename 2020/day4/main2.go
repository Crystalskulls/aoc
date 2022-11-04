package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
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
				continue
			}
			if field == "byr" {
				v, _ := strconv.Atoi(passport[field])
				if v < 1920 || v > 2002 {
					valid = false
					continue
				}
			}
			if field == "iyr" {
				v, _ := strconv.Atoi(passport[field])
				if v < 2010 || v > 2020 {
					valid = false
					continue
				}
			}
			if field == "eyr" {
				v, _ := strconv.Atoi(passport[field])
				if v < 2020 || v > 2030 {
					valid = false
					continue
				}
			}
			if field == "hgt" {
				if !strings.HasSuffix(passport[field], "cm") && !strings.HasSuffix(passport[field], "in") {
					valid = false
					continue
				}
				if strings.HasSuffix(passport[field], "cm") {
					v, _ := strconv.Atoi(passport[field][:len(passport[field])-2])
					if v < 150 || v > 193 {
						valid = false
						continue
					}
				}
				if strings.HasSuffix(passport[field], "in") {
					v, _ := strconv.Atoi(passport[field][:len(passport[field])-2])
					if v < 59 || v > 76 {
						valid = false
						continue
					}
				}
			}
			if field == "hcl" {
				r := regexp.MustCompile("^#[0-9a-f]{6}$")
				if match := r.MatchString(passport[field]); !match {
					valid = false
					continue
				}
			}
			if field == "ecl" {
				r := regexp.MustCompile("amb|blu|brn|gry|grn|hzl|oth")
				if match := r.MatchString(passport[field]); !match {
					valid = false
					continue
				}
			}
			if field == "pid" {
				r := regexp.MustCompile("^\\d{9}$")
				if match := r.MatchString(passport[field]); !match {
					valid = false
					continue
				}
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
