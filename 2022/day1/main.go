package main

import (
	"fmt"
	"sort"

	"github.com/Crystalskulls/aoc/2022/file"
	"github.com/Crystalskulls/aoc/2022/strconv"
)

func main() {
	calories := file.ReadLines("./input.csv")
	totalCalories := make([]int, 0)
	sum := 0
	for _, calorie := range calories {
		if len(calorie) == 0 {
			totalCalories = append(totalCalories, sum)
			sum = 0
			continue
		}
		sum += strconv.ParseToInt(calorie)
	}
	totalCalories = append(totalCalories, sum)
	sort.Ints(totalCalories)
	lastIndex := len(totalCalories) - 1
	fmt.Println("Part One: ", totalCalories[lastIndex])
	sum = 0
	for _, totalCalorie := range totalCalories[len(totalCalories)-3:] {
		sum += totalCalorie
	}
	fmt.Println("Part Two: ", sum)
}
