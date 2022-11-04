package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	data := openCsvFile("input.csv")
	total := 0
	var outputSignalGroups [][]string
	var signalGroupList [][]string

	for _, line := range data {
		signalGroupList = append(signalGroupList, line[0:len(line)-4])
		outputSignalGroups = append(outputSignalGroups, line[len(line)-4:len(line)])
	}

	for x, line := range signalGroupList {
		numberLetterMap := newNumberLetterMap()
		for _, signalGroup := range line {
			length := len(signalGroup)
			if length == 2 {
				numberLetterMap[1] = signalGroup
			} else if length == 3 {
				numberLetterMap[7] = signalGroup
			} else if length == 4 {
				numberLetterMap[4] = signalGroup
			} else if length == 7 {
				numberLetterMap[8] = signalGroup
			}
		}

		for _, signalGroup := range line {
			length := len(signalGroup)
			if length == 6 {
				containsOne := false
				containsFour := false
				containsSeven := false
				counter := 0
				for _, letter := range numberLetterMap[1] {
					if strings.Contains(signalGroup, string(letter)) {
						counter++
					}
				}
				if counter == 2 {
					containsOne = true
				}
				counter = 0
				for _, letter := range numberLetterMap[4] {
					if strings.Contains(signalGroup, string(letter)) {
						counter++
					}
				}
				if counter == 4 {
					containsFour = true
				}
				counter = 0
				for _, letter := range numberLetterMap[7] {
					if strings.Contains(signalGroup, string(letter)) {
						counter++
					}
				}
				if counter == 3 {
					containsSeven = true
				}

				if containsOne && containsFour && containsSeven {
					numberLetterMap[9] = signalGroup
				} else if containsOne && containsSeven {
					numberLetterMap[0] = signalGroup
				} else {
					numberLetterMap[6] = signalGroup
				}
			}
		}

		for _, signalGroup := range line {
			length := len(signalGroup)
			if length == 5 {
				counter := 0
				containsOne := false
				for _, letter := range numberLetterMap[1] {
					if strings.Contains(signalGroup, string(letter)) {
						counter++
					}
				}
				if counter == 2 {
					containsOne = true
				}
				counter = 0
				for _, letter := range numberLetterMap[6] {
					if strings.Contains(signalGroup, string(letter)) {
						counter++
					}
				}
				if counter == 5 {
					numberLetterMap[5] = signalGroup
				} else if containsOne {
					numberLetterMap[3] = signalGroup
				} else {
					numberLetterMap[2] = signalGroup
				}
			}
		}

		fmt.Println("numberLetterMap: ", numberLetterMap)
		value := 0
		for i, signalGroup := range outputSignalGroups[x] {
			for k, v := range numberLetterMap {
				result := true
				for _, letter := range signalGroup {
					if !strings.Contains(v, string(letter)) {
						result = false
					}
				}
				if result == true && (len(signalGroup) == len(v)) {
					value += k * int(math.Pow(10.0, float64(3-i)))
				}
			}
		}
		fmt.Println("value: ", value)
		total += value
	}

	fmt.Println("total: ", total)

	fmt.Println("time: ", time.Since(start))
}

func newNumberLetterMap() map[int]string {
	return map[int]string{
		0: "",
		1: "",
		2: "",
		3: "",
		4: "",
		5: "",
		6: "",
		7: "",
		8: "",
		9: "",
	}
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
