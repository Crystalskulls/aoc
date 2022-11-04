package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"time"
)

type ScoreCard struct {
	score map[string]int
}

type Instruction struct {
	chunks          string
	characterStream []string
}

type NavigationSubsystem struct {
	lines      []*Instruction
	legalPairs map[string]string
}

func main() {
	start := time.Now()
	navigationSubsystem := initNavigationSubsystem(openCsvFile("input.csv"))
	syntaxErrors := make(chan string)
	totalScore := 0
	scoreCard := newScoreCard()
	go navigationSubsystem.compile(syntaxErrors)
	for syntaxError := range syntaxErrors {
		totalScore += scoreCard.score[syntaxError]
	}
	fmt.Println("total score: ", totalScore)
	fmt.Println("time: ", time.Since(start))
}

func (navigationSubsystem *NavigationSubsystem) compile(syntaxErrors chan<- string) {
	for _, instruction := range navigationSubsystem.lines {
		instruction.checkSyntax(navigationSubsystem, syntaxErrors)
	}
	close(syntaxErrors)
}

func (instruction *Instruction) checkSyntax(navigationSubsystem *NavigationSubsystem, syntaxErrors chan<- string) {
	for _, b := range instruction.chunks {
		c := string(b)
		if _, ok := navigationSubsystem.legalPairs[c]; ok {
			// open char
			instruction.characterStream = append(instruction.characterStream, c)
		} else {
			// close char
			lastIndex := len(instruction.characterStream) - 1
			openChar := instruction.characterStream[lastIndex]
			if closeChar := navigationSubsystem.legalPairs[openChar]; closeChar != c {
				// Syntax Error
				syntaxErrors <- c
				break
			}
			instruction.characterStream = instruction.characterStream[:lastIndex]
		}
	}
}

func initNavigationSubsystem(data [][]string) *NavigationSubsystem {
	navigationSubsystem := new(NavigationSubsystem)
	for _, line := range data {
		instruction := new(Instruction)
		instruction.chunks = line[0]
		navigationSubsystem.lines = append(navigationSubsystem.lines, instruction)
	}
	navigationSubsystem.legalPairs = map[string]string{
		"(": ")",
		"[": "]",
		"{": "}",
		"<": ">",
	}
	return navigationSubsystem
}

func newScoreCard() *ScoreCard {
	scoreCard := new(ScoreCard)
	scoreCard.score = make(map[string]int)
	scoreCard.score[")"] = 3
	scoreCard.score["]"] = 57
	scoreCard.score["}"] = 1197
	scoreCard.score[">"] = 25137
	return scoreCard
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
