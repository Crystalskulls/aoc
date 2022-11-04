package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"time"
	"os"
	"strconv"
)

func main() {
	start := time.Now()
	lanternfishes := openCsvFile("input.csv")
	swarm := make(map[int]int)
	for i:=0; i<9; i++ {
		swarm[i] = 0
	}
	for _, lanternfish := range lanternfishes[0] {
		swarm[parseToInt(lanternfish)]++
	}
	fmt.Println(swarm)
	days := 256
	count := growthRate(swarm, days)
	fmt.Println("count: ", count)
	fmt.Printf("time: %v\n", time.Since(start))
}

func growthRate(swarm map[int]int, days int) int{
	for i:=0; i<days; i++ {
		temp := swarm[0]
		swarm[0] = swarm[1]
		swarm[1] = swarm[2]
		swarm[2] = swarm[3]
		swarm[3] = swarm[4]
		swarm[4] = swarm[5]
		swarm[5] = swarm[6]
		swarm[6] = swarm[7] + temp
		swarm[7] = swarm[8]
		swarm[8] = temp
	}

	total := 0
	for _, v := range swarm {
		total += v
	}
	return total
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

func parseToInt(n string) int {
	i, err := strconv.ParseInt(n, 10, 64)
	if err != nil {
		log.Fatalf("Unable to parse number %s to an Integer; err: %v\n", n)
	}
	return int(i)
}
