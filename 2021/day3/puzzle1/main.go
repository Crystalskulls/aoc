package main

import (
	"os"
	"encoding/csv"
	"fmt"
	"math"
)

type powerConsumption struct {
	gamma int
	epsilon int
}

func main() {
	f, _ := os.Open("input.csv")
	r := csv.NewReader(f)
	diagnosticReport, _ := r.ReadAll()
	p := new(powerConsumption)
	byteLength := len(diagnosticReport[0][0])
	p.calcGamma(diagnosticReport, byteLength)
	p.calcEpsilon(diagnosticReport, byteLength)
	fmt.Println("gamma: ", p.gamma)
	fmt.Println("epsilon: ", p.epsilon)
	fmt.Println("power consumption: ", int(p.gamma * p.epsilon))
}

func (p *powerConsumption) calcEpsilon(diagnosticReport [][]string, byteLength int) {
	l := byteLength
	for i:=0; i<byteLength; i++ {
		l--
		counterOne, counterZero := 0, 0
		for _, number := range diagnosticReport {
			// ASCII 49 -> 1; 48 -> 0
			if number[0][i] == 49 {
				counterOne++
			} else {
				counterZero++
			}
		}

		if counterOne < counterZero {
			p.epsilon += int(math.Pow(2, float64(l)))
		}
	}
}

func (p *powerConsumption) calcGamma(diagnosticReport [][]string, byteLength int) {
	l := byteLength
	for i:=0; i<byteLength; i++ {
		l--
		counterOne, counterZero := 0, 0
		for _, number := range diagnosticReport {
			// ASCII 49 -> 1; 48 -> 0
			if number[0][i] == 49 {
				counterOne++
			} else {
				counterZero++
			}
		}

		if counterOne > counterZero {
			p.gamma += int(math.Pow(2, float64(l)))
		}
	}
}
