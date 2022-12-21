package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"gopkg.in/Knetic/govaluate.v3"
)

func main() {
	numberYellingMonkeys := make(map[string]int)
	operationMonkeys := make(map[string][]string)

	file, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(file), "\n")

	for _, line := range lines {
		snippets := strings.Split(line, " ")
		if len(snippets) == 2 {
			n, _ := strconv.Atoi(snippets[1])
			numberYellingMonkeys[strings.ReplaceAll(snippets[0], ":", "")] = n
		} else {
			operationMonkeys[strings.ReplaceAll(snippets[0], ":", "")] = snippets[1:]
		}
	}

	for len(operationMonkeys) > 0 {
		for name, operationSnippets := range operationMonkeys {
			if _, ok := numberYellingMonkeys[operationSnippets[0]]; !ok {
				continue
			}
			if _, ok := numberYellingMonkeys[operationSnippets[2]]; !ok {
				continue
			}
			operation, _ := govaluate.NewEvaluableExpression(fmt.Sprintf("%s %s %s", operationSnippets[0], operationSnippets[1], operationSnippets[2]))
			parameters := make(map[string]interface{}, 8)
			parameters[operationSnippets[0]] = numberYellingMonkeys[operationSnippets[0]]
			parameters[operationSnippets[2]] = numberYellingMonkeys[operationSnippets[2]]
			result, _ := operation.Evaluate(parameters)
			numberYellingMonkeys[name] = int(result.(float64))
			delete(operationMonkeys, name)
		}
	}
	fmt.Println("PartOne:", numberYellingMonkeys["root"])
}
