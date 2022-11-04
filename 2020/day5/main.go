package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type BoardingPass struct {
	seat   string
	seatID int
	row    int
	col    int
}

func main() {
	records := readCSV("input.csv")
	boardingPasses := parseToBoardingPass(records)
	high := 0
	for _, boardingPass := range boardingPasses {
		boardingPass.decode()
		if boardingPass.seatID > high {
			high = boardingPass.seatID
		}
	}
	fmt.Printf("Highest seatID: %d\n", high)
	seatIDMap := getSeatIDMap(boardingPasses)
	fmt.Printf("My seatID: %d\n", findMySeat(seatIDMap))
}

func findMySeat(seatIDMap map[int]struct{}) int {
	for i := 1; i < 1023; i++ {
		if _, ok := seatIDMap[i]; !ok {
			_, previous := seatIDMap[i-1]
			_, next := seatIDMap[i+1]
			if previous && next {
				return i
			}
		}
	}
	return -1
}

func getSeatIDMap(boardingPasses []*BoardingPass) map[int]struct{} {
	seatIDs := make(map[int]struct{}, len(boardingPasses))
	for _, bp := range boardingPasses {
		seatIDs[bp.seatID] = struct{}{}
	}
	return seatIDs
}

func (boardingPass *BoardingPass) decode() {
	binaryRow := strings.ReplaceAll(strings.ReplaceAll(boardingPass.seat[:7], "F", "0")[:7], "B", "1")
	binaryCol := strings.ReplaceAll(strings.ReplaceAll(boardingPass.seat[7:], "L", "0"), "R", "1")
	row, _ := strconv.ParseInt(binaryRow, 2, 64)
	boardingPass.row = int(row)
	col, _ := strconv.ParseInt(binaryCol, 2, 64)
	boardingPass.col = int(col)
	boardingPass.seatID = boardingPass.row*8 + boardingPass.col
}

func parseToBoardingPass(records [][]string) (boardingPasses []*BoardingPass) {
	for _, record := range records {
		boardingPass := new(BoardingPass)
		boardingPass.seat = record[0]
		boardingPasses = append(boardingPasses, boardingPass)
	}
	return
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
