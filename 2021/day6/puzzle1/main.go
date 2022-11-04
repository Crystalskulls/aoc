package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"time"
	"os"
	"strconv"
)

type Lanternfish struct {
	age int
	baby bool
}

func main() {
	start := time.Now()
	data := openCsvFile("input.csv")
	var lanternfishes []*Lanternfish
	for _, d := range data[0] {
		lanternfishes = append(lanternfishes, newLanternfish(parseToInt(d)))
	}
	fmt.Printf("Initial state:   %v\n", lanternfishes)
	days := 80
	lanternfishes = growthRate(lanternfishes, days)
	fmt.Printf("Total Lanternfishes after %d days: %d\n", days, len(lanternfishes))
	fmt.Printf("time: %v\n", time.Since(start))
}

func growthRate(lanternfishes []*Lanternfish, days int) []*Lanternfish{
	for i:=1; i<=days; i++ {
		for _, lanternfish := range lanternfishes {
			if lanternfish.age == 0 {
				lanternfishes = append(lanternfishes, newLanternfish(8))
			}
			if lanternfish.baby {
				lanternfish.baby = false
			} else {
				lanternfish.decreaseAge()
			}
		}
		fmt.Printf("after day %d: %v\n", i, len(lanternfishes))
	}
	return lanternfishes
}

func (lanternfish *Lanternfish) decreaseAge() {
	if lanternfish.age == 0 {
		lanternfish.age = 6
		lanternfish.baby = true
	}
	lanternfish.age--
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

func newLanternfish(age int) (lanternfish *Lanternfish) {
	lanternfish = new(Lanternfish)
	lanternfish.age = age
	return lanternfish
}

func parseToInt(n string) int {
	i, err := strconv.ParseInt(n, 10, 64)
	if err != nil {
		log.Fatalf("Unable to parse number %s to an Integer; err: %v\n", n)
	}
	return int(i)
}

func (lanternfish *Lanternfish) String() string {
	return fmt.Sprintf("%d", lanternfish.age)
}
