package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	sum := 0
	for _, snafuNumber := range parseSNAFUNumbers() {
		sum += convertToDecimal(snafuNumber)
	}
	fmt.Println(convertToSNAFU(sum))
}

func convertToDecimal(snafuNumber string) (decimal int) {
	for i, r := range snafuNumber {
		n, err := strconv.Atoi(string(r))
		if err != nil {
			if string(r) == "-" {
				n = -1
			} else {
				n = -2
			}
		}
		decimal += n * int(math.Pow(5, float64((len(snafuNumber)-(i+1)))))
	}
	return
}

func convertToSNAFU(decimal int) (snafuNumber string) {
	snafu := make([]string, 0)

	snafuMap := map[int]string{
		0: "0",
		1: "-",
		2: "=",
	}
	quotient := decimal
	mind := "0"
	for {
		snafu = append(snafu, mind)
		if quotient == 0 {
			break
		}
		remainder := quotient % 5
		quotient = quotient / 5
		currentValue, _ := strconv.Atoi(snafu[len(snafu)-1])
		currentValue += remainder
		if remainder > 2 {
			snafu[len(snafu)-1] = snafuMap[5-(currentValue)]
			mind = "1"
		} else {
´			if currentValue > 2 {
				snafu[len(snafu)-1] = snafuMap[5-(currentValue)]
				mind = "1"
			} else {
				mind = "0"
				snafu[len(snafu)-1] = strconv.Itoa(currentValue)
			}
		}
	}
	if snafu[len(snafu)-1] == "0" {
		snafu = snafu[:len(snafu)-1]
	}
	for i := len(snafu) - 1; i >= 0; i-- {
		snafuNumber += snafu[i]
	}
	return
}

func parseSNAFUNumbers() []string {
	file, _ := os.ReadFile("input.txt")
	return strings.Split(string(file), "\n")
}
