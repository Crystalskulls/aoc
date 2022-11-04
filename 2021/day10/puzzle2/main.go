package main

import (
	"encoding/csv"
	"fmt"
	"sort"
	"log"
	"os"
)

type ScoreCard struct {
	score map[string]int
}

type NavigationSubsystem struct {
	lines []string
	legalPairs map[string]string
	characterStream []string
}

func main() {
	navigationSubsystem := initNavigationSubsystem(openCsvFile("input.csv"))
	completionStrings := navigationSubsystem.autocomplete()
	scores := autocompleteScore(completionStrings)
	fmt.Println("scores: ", scores)
	fmt.Println("middle score: ", scores[len(scores) / 2])
}

func (navigationSubsystem *NavigationSubsystem) autocomplete() (completionStrings []string) {
	for _, line := range navigationSubsystem.lines {
		navigationSubsystem.clearStream()
		syntaxError := false
		for _, b := range line {
			c := string(b)
			if _, ok := navigationSubsystem.legalPairs[c]; ok {
				// open char
				navigationSubsystem.characterStream = append(navigationSubsystem.characterStream, c)
			} else {
				// close char
				lastIndex := len(navigationSubsystem.characterStream)-1
				openChar := navigationSubsystem.characterStream[lastIndex]
				if closeChar := navigationSubsystem.legalPairs[openChar]; closeChar != c {
					// Syntax Error
					syntaxError = true
					break
				}
				navigationSubsystem.characterStream = navigationSubsystem.characterStream[:lastIndex]
			}
		}

		if !syntaxError && len(navigationSubsystem.characterStream) != 0 {
			completionString := navigationSubsystem.complete()
			fmt.Println("completionString: ", completionString)
			completionStrings = append(completionStrings, completionString)
		}
	}
	return completionStrings
}

func (navigationSubsystem *NavigationSubsystem) complete() (completionString string) {
	lastIndex := len(navigationSubsystem.characterStream)-1
	for i:=lastIndex; i>=0; i-- {
		c := navigationSubsystem.characterStream[i]
		completionString += navigationSubsystem.legalPairs[c]
	}
	return completionString
}

func (navigationSubsystem *NavigationSubsystem) clearStream() {
	navigationSubsystem.characterStream = []string{}
}

func autocompleteScore(completionStrings []string) []int {
	totalScores := make([]int, len(completionStrings))
	scoreCard := newScoreCard()
	for i, completionString := range completionStrings {
		score := 0
		for _, b := range completionString {
			c := string(b)
			score *= 5
			score += scoreCard.score[c]
		}
		totalScores[i] = score
	}
	sort.Ints(totalScores)
	return totalScores
}

func initNavigationSubsystem (data [][]string) *NavigationSubsystem{
	navigationSubsystem := new(NavigationSubsystem)
	for _, line := range data {
		navigationSubsystem.lines = append(navigationSubsystem.lines, line[0])
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
	scoreCard.score[")"] = 1
	scoreCard.score["]"] = 2
	scoreCard.score["}"] = 3
	scoreCard.score[">"] = 4
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
