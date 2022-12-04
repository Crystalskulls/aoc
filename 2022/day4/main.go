package main

import (
	"os"
	"strconv"
	"strings"
)

func main() {
	d, _ := os.ReadFile("input.txt")
	pairs := strings.Split(string(d), "\n")
	c1, c2 := 0, 0
	for _, p := range pairs {
		s1, s2 := strings.Split(p, ",")[0], strings.Split(p, ",")[1]
		start1, _ := strconv.Atoi(strings.Split(s1, "-")[0])
		start2, _ := strconv.Atoi(strings.Split(s2, "-")[0])
		end1, _ := strconv.Atoi(strings.Split(s1, "-")[1])
		end2, _ := strconv.Atoi(strings.Split(s2, "-")[1])
		if (start1 <= start2 && end1 >= end2) || (start2 <= start1 && end2 >= end1) {
			c1++
		}
		if (start1 <= end2) && (end1 >= start2) {
			c2++
		}
	}
	print("Part One: ", c1, " Part Two: ", c2)
}
