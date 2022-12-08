package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	partOne, partTwo := findVisibleTrees(createGrid())
	fmt.Printf("PartOne: %d, PartTwo: %d\n", partOne, partTwo)
}

func createGrid() [][]int {
	data, _ := os.ReadFile("input.txt")
	rows := strings.Split(string(data), "\n")
	array := make([][]int, len(rows))

	for i, row := range rows {
		array[i] = make([]int, len(row))
		for j, r := range row {
			array[i][j], _ = strconv.Atoi(string(r))
		}
	}
	return array
}

func findVisibleTrees(grid [][]int) (c, highestScenicScore int) {
	c = len(grid)*4 - 4
	for x := 1; x < len(grid)-1; x++ {
		for y := 1; y < len(grid)-1; y++ {
			visible, scenicScore := isVisibleTree(grid, x, y)

			if visible {
				c++
			}
			if scenicScore > highestScenicScore {
				highestScenicScore = scenicScore
			}
		}
	}
	return
}

func isVisibleTree(grid [][]int, x, y int) (bool, int) {
	left, right := make([]int, 0), make([]int, 0)
	up, down := make([]int, 0), make([]int, 0)
	scenicScores := make([]int, 0)

	for i, row := range grid {
		if i < x {
			up = append(up, row[y])
		} else if i > x {
			down = append(down, row[y])
		} else {
			for j, n := range row {
				if j < y {
					left = append(left, n)
				}
				if j > y {
					right = append(right, n)
				}
			}
		}
	}
	h := grid[x][y]

	scenicScores = append(scenicScores, calcScenicScore(right, h, false))
	scenicScores = append(scenicScores, calcScenicScore(down, h, false))
	scenicScores = append(scenicScores, calcScenicScore(left, h, true))
	scenicScores = append(scenicScores, calcScenicScore(up, h, true))
	scenicScore := 1
	for _, n := range scenicScores {
		scenicScore *= n
	}
	sort.Sort(sort.Reverse(sort.IntSlice(left)))
	sort.Sort(sort.Reverse(sort.IntSlice(right)))
	sort.Sort(sort.Reverse(sort.IntSlice(up)))
	sort.Sort(sort.Reverse(sort.IntSlice(down)))

	return left[0] < h || right[0] < h || up[0] < h || down[0] < h, scenicScore
}

func calcScenicScore(line []int, h int, reverse bool) (scenicScore int) {
	if reverse {
		for i := len(line) - 1; i >= 0; i-- {
			if line[i] >= h {
				scenicScore++
				return
			}
			scenicScore++
		}
		return
	}
	for i := 0; i < len(line); i++ {
		if line[i] >= h {
			scenicScore++
			return
		}
		scenicScore++
	}
	return
}
