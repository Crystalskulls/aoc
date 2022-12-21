package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"gopkg.in/Knetic/govaluate.v3"
)

func main() {
	i := 3423279930000
	for {
		i++
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
		operationMonkeys["root"][1] = "="
		numberYellingMonkeys["humn"] = i

		for len(operationMonkeys) > 0 {
			rootFound := false
			for name, operationSnippets := range operationMonkeys {
				if _, ok := numberYellingMonkeys[operationSnippets[0]]; !ok {
					continue
				}
				if _, ok := numberYellingMonkeys[operationSnippets[2]]; !ok {
					continue
				}
				if name == "root" {
					rootFound = true
					fmt.Printf("%d: %d = %d\n", i, numberYellingMonkeys[operationSnippets[0]], numberYellingMonkeys[operationSnippets[2]])
					if numberYellingMonkeys[operationSnippets[0]] == numberYellingMonkeys[operationSnippets[2]] {
						fmt.Println("PartTwo", i)
						os.Exit(1)
					}
				} else {
					operation, _ := govaluate.NewEvaluableExpression(fmt.Sprintf("%s %s %s", operationSnippets[0], operationSnippets[1], operationSnippets[2]))
					parameters := make(map[string]interface{}, 8)
					parameters[operationSnippets[0]] = numberYellingMonkeys[operationSnippets[0]]
					parameters[operationSnippets[2]] = numberYellingMonkeys[operationSnippets[2]]
					result, _ := operation.Evaluate(parameters)
					numberYellingMonkeys[name] = int(result.(float64))
					delete(operationMonkeys, name)
				}
			}
			if rootFound {
				break
			}
		}
	}
}
