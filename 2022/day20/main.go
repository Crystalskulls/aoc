package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Number struct {
	value      int
	prevNumber *Number
	nextNumber *Number
}

func main() {
	numbers, zeroNumber := readInput()
	mixing(numbers)
	fmt.Printf("PartOne: %d\n", sumOfGroveCoordinates(numbers, zeroNumber))
}

func sumOfGroveCoordinates(numbers []*Number, zeroNumber *Number) (sum int) {
	for i := 1; i <= 3; i++ {
		number := zeroNumber
		j := (i * 1000) % len(numbers)
		for k := 1; k <= j; k++ {
			number = number.nextNumber
		}
		sum += number.value
	}
	return sum
}

func mixing(numbers []*Number) {
	for _, currentNumber := range numbers {
		if currentNumber.value > 0 {
			currentNumber.moveRight()
		} else {
			currentNumber.moveLeft()
		}
	}
}

func (number *Number) moveRight() {
	for i := 0; i < number.value; i++ {
		nextNumber := number.nextNumber
		currentNumberPrev := number.prevNumber
		number.prevNumber = nextNumber
		number.nextNumber = nextNumber.nextNumber
		nextNumber.nextNumber = number
		nextNumber.prevNumber = currentNumberPrev
		nextNumber.prevNumber.nextNumber = nextNumber
		number.nextNumber.prevNumber = number
	}
}

func (number *Number) moveLeft() {
	for i := 0; i > number.value; i-- {
		prevNumber := number.prevNumber
		prevNumberPrev := prevNumber.prevNumber
		prevNumber.nextNumber = number.nextNumber
		prevNumber.prevNumber = number
		number.nextNumber = prevNumber
		number.prevNumber = prevNumberPrev
		prevNumber.nextNumber.prevNumber = prevNumber
		number.prevNumber.nextNumber = number
	}
}

func readInput() (numbers []*Number, zeroNumber *Number) {
	input, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(input), "\n")
	numbers = make([]*Number, len(lines))
	for i, line := range lines {
		n, _ := strconv.Atoi(line)
		numbers[i] = &Number{
			value: n,
		}
		if n == 0 {
			zeroNumber = numbers[i]
		}
	}
	for j := range numbers {
		if j == 0 {
			numbers[j].prevNumber = numbers[len(numbers)-1]
			numbers[j].nextNumber = numbers[j+1]
		} else if j == len(numbers)-1 {
			numbers[j].nextNumber = numbers[0]
			numbers[j].prevNumber = numbers[j-1]
		} else {
			numbers[j].nextNumber = numbers[j+1]
			numbers[j].prevNumber = numbers[j-1]
		}
	}
	return numbers, zeroNumber
}

func (number *Number) String() string {
	return fmt.Sprintf("%p: %d, next: %p, prev: %p\n", number, number.value, number.nextNumber, number.prevNumber)
}
