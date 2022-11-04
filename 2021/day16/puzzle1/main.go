package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"math"
)

type TransmissionSystem struct {
	binaryStream []int
	bitsRead int
}

func main() {
	hexString := openCsvFile("input.csv")[0][0]
	transmissionSystem := new(TransmissionSystem)
	transmissionSystem.binaryStream = parseToBitSlice(hexString)

	versionCount := 0
	for len(transmissionSystem.binaryStream) != 0 {
		versionNumber, _ := transmissionSystem.read(3)
		versionCount += versionNumber
		typeID, _ := transmissionSystem.read(3)

		if isLiteralValue(typeID) {
			// literal value packet
			_, isLastGroup := transmissionSystem.read(5)
			for !isLastGroup {
				_, isLastGroup = transmissionSystem.read(5)
			}
		} else {
			// operator packet
			lengthTypeID, _ := transmissionSystem.read(1)
			if lengthTypeID == 0 {
				// total length in bits
				_, _ = transmissionSystem.read(15)

			} else {
				// number of sub-packetes
				_, _ = transmissionSystem.read(11)
			}
		}
	}
	fmt.Println("version count: ", versionCount)
}

func isLiteralValue(typeID int) bool {
	if typeID == 4 {
		return true
	}
	return false
}

func (transmissionSystem *TransmissionSystem) read(bits int) (decimal int, isLastGroup bool){
	if bits > len(transmissionSystem.binaryStream) {
		bits = 0
	}
	for i:=0; i<bits; i++ {
		y := bits-1-i
		bit := transmissionSystem.binaryStream[i]
		if i == 0 && bit == 0 {
			isLastGroup = true
		}
		decimal += bit * int(math.Pow(2.0, float64(y)))
	}
	transmissionSystem.binaryStream = transmissionSystem.binaryStream[bits:]
	transmissionSystem.bitsRead += bits
	return
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

func parseToBitSlice(hex string) []int {
	var bits []int
	for i:=0; i < len(hex); i++ {
		var nibble []int
		n, err := strconv.ParseUint(string(hex[i]), 16, 4)
		if err != nil {
			log.Fatalf("can not convert hex string to uint64; err: ", err)
		}
		for j:=0; j<4; j++ {
			nibble = append([]int{int(n) & 0x1}, nibble...)
			n = n >> 1
		}
		bits = append(bits, nibble...)
	}
	return bits
}
