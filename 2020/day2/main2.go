package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Policy struct {
	letter string
	pos1   int
	pos2   int
}

type Password struct {
	value  string
	policy *Policy
}

func main() {
	passwordDatabase := readCSV("./input.csv")
	passwords := parseToPasswordSlice(passwordDatabase)

	valid := 0
	for _, password := range passwords {
		if password.isValid() {
			valid++
		}
	}
	fmt.Println("Result: ", valid)
}

func readCSV(path string) [][]string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Can't open file %s\n", path)
	}
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("Can't read file %s\n", path)
	}
	return records
}

func parseToPasswordSlice(db [][]string) []*Password {
	ps := []*Password{}
	for _, entry := range db {
		p := new(Password)
		p.policy = new(Policy)
		snippets := strings.Split(entry[0], " ")
		p.value = snippets[2]
		pos := strings.Split(snippets[0], "-")
		pos1, _ := strconv.Atoi(pos[0])
		pos2, _ := strconv.Atoi(pos[1])
		p.policy.pos1 = pos1
		p.policy.pos2 = pos2
		p.policy.letter = snippets[1][0:1]
		ps = append(ps, p)
	}
	return ps
}

func (password *Password) isValid() bool {
	c := 0
	if string(password.value[password.policy.pos1-1]) == password.policy.letter {
		c++
	}
	if string(password.value[password.policy.pos2-1]) == password.policy.letter {
		c++
	}
	return c == 1
}
