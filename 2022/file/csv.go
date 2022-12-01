package file

import (
	"bufio"
	"encoding/csv"
	"log"
	"os"
)

func ReadCsvFile(path string) [][]string {
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

func ReadLines(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Can't open file %s\n", path)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split((bufio.ScanLines))
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}
