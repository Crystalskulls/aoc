package main

import (
	"encoding/csv"
	"fmt"
	"github.com/fatih/color"
	"log"
	"math"
	"os"
	"strconv"
)

type Instruction struct {
	foldAlong string
	index     int
}

type TransparentPaper struct {
	dotLines         [][]bool
	foldInstructions []*Instruction
}

func main() {
	transparentPaper := mapDots(openCsvFile("input.csv"))
	transparentPaper.foldInstructions = createInstructions(openCsvFile("fold_instructions.csv"))
	transparentPaper.fold()
	fmt.Println(transparentPaper)
}

func (transparentPaper *TransparentPaper) fold() {
	for _, instruction := range transparentPaper.foldInstructions {
		if instruction.foldAlong == "y" {
			for i := instruction.index + 1; i < len(transparentPaper.dotLines); i++ {
				for j := 0; j < len(transparentPaper.dotLines[i]); j++ {
					if transparentPaper.dotLines[i][j] {
						newIndex := instruction.index - (i - instruction.index)
						transparentPaper.dotLines[newIndex][j] = true
					}
				}
			}
			transparentPaper.dotLines = transparentPaper.dotLines[:instruction.index]
		} else {
			for i := 0; i < len(transparentPaper.dotLines); i++ {
				for j := instruction.index - 1; j >= 0; j-- {
					if transparentPaper.dotLines[i][j] {
						newIndex := instruction.index + int((math.Abs(float64(j - instruction.index))))
						transparentPaper.dotLines[i][newIndex] = true
					}
				}
			}

			for i, _ := range transparentPaper.dotLines {
				transparentPaper.dotLines[i] = transparentPaper.dotLines[i][instruction.index+1:]
			}
		}
	}
}

func (transparentPaper *TransparentPaper) countDots() int {
	i := 0
	for _, line := range transparentPaper.dotLines {
		for _, dot := range line {
			if dot {
				i++
			}
		}
	}
	return i
}

func createInstructions(data [][]string) (instructions []*Instruction) {
	for _, line := range data {
		instructions = append(instructions,
			&Instruction{
				foldAlong: line[0],
				index:     parseToInt(line[1]),
			},
		)
	}
	return instructions
}

func mapDots(data [][]string) *TransparentPaper {
	transparentPaper := newTransparentPaper(1311)
	for _, line := range data {
		x, y := parseToInt(line[0]), parseToInt(line[1])
		transparentPaper.dotLines[y][x] = true
	}
	return transparentPaper
}

func newTransparentPaper(i int) *TransparentPaper {
	transparentPaper := new(TransparentPaper)
	transparentPaper.dotLines = make([][]bool, i)
	for j := 0; j < i; j++ {
		transparentPaper.dotLines[j] = make([]bool, i)
	}
	return transparentPaper
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

func (transparentPaper *TransparentPaper) String() string {
	s := ""
	red := color.New(color.FgRed).SprintFunc()
	for _, line := range transparentPaper.dotLines {
		for _, dot := range line {
			if dot {
				s += fmt.Sprintf("%v", red("#"))
			} else {
				s += "."
			}
		}
		s += fmt.Sprintf("\n")
	}
	return s
}
