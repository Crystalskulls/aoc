package main

import (
	"encoding/csv"
	"fmt"
	"github.com/fatih/color"
	"log"
	"os"
	"strconv"
)

type Number struct {
	value  int
	marked bool
}

type Row struct {
	numbers       []*Number
	markedNumbers int
}

type Column struct {
	numbers       []*Number
	markedNumbers int
}

type BingoBoard struct {
	rows    []*Row
	columns []*Column
	bingo   bool
}

func main() {
	bingoNumbers := parseBingoNumbers(openCsvFile("numbers.csv"))
	bingoBoards := parseBingoBoards(openCsvFile("boards.csv"))
	winnerBoard, lastNumber := play(bingoNumbers, bingoBoards)
	if winnerBoard == nil {
		log.Fatal("winnerBoard is nil")
	}
	sum := sumUnmarkedNumbers(winnerBoard)
	fmt.Println(winnerBoard)
	fmt.Printf("%d * %d = %d\n", sum, lastNumber, sum*lastNumber)
}

func play(bingoNumbers []int, bingoBoards []*BingoBoard) (winnerBoard *BingoBoard, lastNumber int) {
	for _, bingoNumber := range bingoNumbers {
		for _, bingoBoard := range bingoBoards {
			bingoBoard.check(bingoNumber)
			if bingoBoard.bingo {
				return bingoBoard, bingoNumber
			}
		}
	}
	return nil, -1
}

func (bingoBoard *BingoBoard) check(bingoNumber int) {
	for _, row := range bingoBoard.rows {
		for j, number := range row.numbers {
			if !(number.marked) && number.value == bingoNumber {
				number.marked = true
				row.markedNumbers++
				bingoBoard.columns[j].markedNumbers++
			}
			if row.markedNumbers == 5 || bingoBoard.columns[j].markedNumbers == 5 {
				bingoBoard.bingo = true
			}
		}
	}

}

func parseBingoBoards(data [][]string) (bingoBoards []*BingoBoard) {
	currentBoard := newBingoBoard()
	i := 0
	for j, line := range data {
		for k, element := range line {
			n := parseToInt(element)
			number := new(Number)
			number.value = n
			currentBoard.rows[i].numbers[k] = number
			currentBoard.columns[k].numbers[i] = number
		}
		i++
		if (j+1)%5 == 0 {
			bingoBoards = append(bingoBoards, currentBoard)
			currentBoard = newBingoBoard()
			i = 0
		}
	}
	return bingoBoards
}

func newBingoBoard() *BingoBoard {
	bingoBoard := new(BingoBoard)
	for i := 0; i < 5; i++ {
		bingoBoard.rows = append(bingoBoard.rows, newRow())
		bingoBoard.columns = append(bingoBoard.columns, newColumn())
	}
	return bingoBoard
}

func newRow() *Row {
	row := new(Row)
	for i := 0; i < 5; i++ {
		row.numbers = append(row.numbers, new(Number))
	}
	return row
}

func newColumn() *Column {
	column := new(Column)
	for i := 0; i < 5; i++ {
		column.numbers = append(column.numbers, new(Number))
	}
	return column
}

func parseBingoNumbers(data [][]string) (bingoNumbers []int) {
	for _, number := range data[0] {
		bingoNumbers = append(bingoNumbers, parseToInt(number))
	}
	return bingoNumbers
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

func sumUnmarkedNumbers(bingoBoard *BingoBoard) int {
	sum := 0
	for _, row := range bingoBoard.rows {
		for _, number := range row.numbers {
			if number.marked {
				continue
			}
			sum += number.value
		}
	}
	return sum
}

func (bingoBoard *BingoBoard) String() string {
	red := color.New(color.FgRed).SprintFunc()
	s := ""
	for _, row := range bingoBoard.rows {
		for _, number := range row.numbers {
			if number.marked {
				s += fmt.Sprintf("%v ", red(number.value))
			} else {
				s += fmt.Sprintf("%v ", number.value)
			}
		}
		s += fmt.Sprintf("\n")
	}
	return s
}
