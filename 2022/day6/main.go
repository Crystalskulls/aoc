package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("Part One: %d - Part Two: %d\n", getMarker(4), getMarker(14))
}

func getMarker(dc int) (marker int) {
	file, _ := os.ReadFile("input.txt")
	bs := string(file)
	cm := make(map[rune]struct{})
	for i := dc - 1; i <= len(bs); i++ {
		cs := bs[i-(dc-1) : i+1]
		for _, c := range cs {
			cm[c] = struct{}{}
		}
		if len(cm) == dc {
			return i + 1
		}
		cm = make(map[rune]struct{})
	}
	return 0
}
