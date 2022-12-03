package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	data, _ := os.ReadFile("input.txt")
	backpacks := strings.Split(string(data), "\n")
	sum, sum2 := 0, 0
	g := make([]string, 3)
	for i, b := range backpacks {
		compartments := []string{b[0 : len(b)/2], b[len(b)/2:]}
		sum += compare(compartments)
		g[i%3] = b
		if (i+1)%3 == 0 {
			sum2 += compare(g)
			g = make([]string, 3)
		}
	}
	fmt.Printf("Part 1: %d - Part2: %d\n", sum, sum2)
}

func compare(items []string) int {
	for _, r := range items[0] {
		if len(items) == 2 && strings.ContainsRune(items[1], r) {
			return calcPrio(r)
		}
		if strings.ContainsRune(items[1], r) && strings.ContainsRune(items[2], r) {
			return calcPrio(r)
		}
	}
	return 0
}

func calcPrio(r rune) int {
	if r > 96 {
		return int(r) - 96
	}
	return int(r) - 38
}
