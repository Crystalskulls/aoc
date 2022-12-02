package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	data, _ := os.ReadFile("input.csv")
	calSets := strings.Split(string(data), "\n\n")
	totalCals := make([]int, len(calSets))
	for i, calSet := range calSets {
		for _, cal := range strings.Split(calSet, "\n") {
			c, _ := strconv.Atoi(cal)
			totalCals[i] += c
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(totalCals)))
	fmt.Printf("Part One: %d - Part Two: %d\n", totalCals[0], totalCals[0]+totalCals[1]+totalCals[2])
}
