package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

type bus struct {
	id     int
	offset int
}

func main() {
	busMap, firstBus := readCSV("./input.csv")
	fmt.Println(firstBus)
	fmt.Println(busMap)
	fmt.Println(findTimestamp(busMap, firstBus))
}

func findTimestamp(busMap map[int]*bus, firstBus *bus) int {
	t := firstBus.id * 4347826086956
	for {
		s := false
		t += firstBus.id
		fmt.Println(t)
		for _, b := range busMap {
			if (t+b.offset)%b.id != 0 {
				s = true
			}
			if s {
				break
			}
		}
		if !s {
			return t
		}
	}
}

func (b *bus) String() string {
	return fmt.Sprintf("Bus-ID: %d - offset: %d", b.id, b.offset)
}

func readCSV(path string) (busMap map[int]*bus, firstBus *bus) {
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

func parseToInt(s [][]string) (busMap map[int]*bus, firstBus *bus) {
	k := make(map[int]*bus)
	for i, v := range s[0] {

		j, err := strconv.Atoi(v)
		if err != nil {
			continue
		}
		if i == 0 {
			firstBus = new(bus)
			firstBus.id = j
			firstBus.offset = 0
			continue
		}
		b := new(bus)
		b.id = j
		b.offset = i
		k[j] = b
	}
	return k, firstBus
}
