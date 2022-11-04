package main

import (
	"os"
	"encoding/csv"
	"fmt"
	"math"
)

type lifeSupportRating struct {
	oxy float64
	co2 float64
}

func main() {
	f, _ := os.Open("input.csv")
	r := csv.NewReader(f)
	diagnosticReport, _ := r.ReadAll()
	lsr := new(lifeSupportRating)
	byteLength := len(diagnosticReport[0][0])
	lsr.calcOxy(diagnosticReport, byteLength)
	lsr.calcCo2(diagnosticReport, byteLength)
	fmt.Println("oxy: ", lsr.oxy)
	fmt.Println("co2: ", lsr.co2)
	fmt.Println("life support rating: ", int(lsr.oxy * lsr.co2))
}

func (lsr *lifeSupportRating) calcCo2(diagnosticReport [][]string, byteLength int) {
	var remainingNumbers [][]string
	remainingNumbers = diagnosticReport

	for i:=0; i<byteLength; i++ {
		if len(remainingNumbers) == 1 {
			break
		}
		var sOne, sZero [][]string
		for _, number := range remainingNumbers {
			// ASCII 49 -> 1; 48 -> 0
			if number[0][i] == 49 {
				sOne = append(sOne, number)
			} else {
				sZero = append(sZero, number)
			}
		}

		if len(sOne) < len(sZero) {
			remainingNumbers = sOne
		} else if len(sOne) == len(sZero) {
			remainingNumbers = sZero
		} else {
			remainingNumbers = sZero
		}
	}

	lsr.co2 = decimalValue(remainingNumbers[0][0])
}

func (lsr *lifeSupportRating) calcOxy(diagnosticReport [][]string, byteLength int) {
	var remainingNumbers [][]string
	remainingNumbers = diagnosticReport

	for i:=0; i<byteLength; i++ {
		if len(remainingNumbers) == 1 {
			break
		}
		var sOne, sZero [][]string
		for _, number := range remainingNumbers {
			// ASCII 49 -> 1; 48 -> 0
			if number[0][i] == 49 {
				sOne = append(sOne, number)
			} else {
				sZero = append(sZero, number)
			}
		}

		if len(sOne) > len(sZero) {
			remainingNumbers = sOne
		} else if len(sOne) == len(sZero) {
			remainingNumbers = sOne
		} else {
			remainingNumbers = sZero
		}
	}

	lsr.oxy = decimalValue(remainingNumbers[0][0])
}

func decimalValue(bits string) float64 {
	var d float64
	l := len(bits)
	for _, bit := range bits {
		l--
		if bit == 49 {
			d += math.Pow(2, float64(l))
		}
	}
	return d
}
