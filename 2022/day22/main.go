package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type Player struct {
	x         int
	y         int
	direction string
}

func main() {
	board := parseBoard()
	player := newPlayer(board)
	player.Walk(getPathDescription(), board)
	facingPoints := map[string]int{
		"^": 3,
		"<": 2,
		">": 0,
		"v": 1,
	}
	fmt.Println((player.x+1)*1000 + (4 * (player.y + 1)) + facingPoints[player.direction])
}

func (player *Player) Walk(path []string, board [][]string) {
	for _, e := range path {
		steps, err := strconv.Atoi(e)
		if err != nil {
			player.Turn(e)
			continue
		}
		player.move(steps, board)
	}
}

func (player *Player) Turn(direction string) {
	directionMap := map[string]map[string]string{
		"L": map[string]string{
			"^": "<",
			"v": ">",
			"<": "v",
			">": "^",
		},
		"R": map[string]string{
			"^": ">",
			"v": "<",
			"<": "^",
			">": "v",
		},
	}
	player.direction = directionMap[direction][player.direction]
}

func (player *Player) move(steps int, board [][]string) {
	for i := 0; i < steps; i++ {
		if player.direction == ">" {
			player.moveRight(board)
		} else if player.direction == "<" {
			player.moveLeft(board)
		} else if player.direction == "^" {
			player.moveUp(board)
		} else {
			player.moveDown(board)
		}
	}
}

func (player *Player) moveRight(board [][]string) {
	nextCol := player.y + 1
	if nextCol == len(board[player.x]) || strings.TrimSpace(board[player.x][nextCol]) == "" {
		for col, s := range board[player.x] {
			if s == "#" {
				return
			}
			if s == "." {
				player.y = col
				return
			}
		}
	}
	if board[player.x][nextCol] == "#" {
		return
	}
	player.y = nextCol
}

func (player *Player) moveLeft(board [][]string) {
	nextCol := player.y - 1
	if nextCol < 0 || strings.TrimSpace(board[player.x][nextCol]) == "" {
		for col := len(board[player.x]) - 1; col >= 0; col-- {
			if board[player.x][col] == "#" {
				return
			}
			if board[player.x][col] == "." {
				player.y = col
				return
			}
		}
	}
	if board[player.x][nextCol] == "#" {
		return
	}
	player.y = nextCol
}

func (player *Player) moveDown(board [][]string) {
	nextRow := player.x + 1
	if nextRow == len(board) || strings.TrimSpace(board[nextRow][player.y]) == "" {
		for x, row := range board {
			if row[player.y] == "#" {
				return
			}
			if row[player.y] == "." {
				player.x = x
				return
			}
		}
	}
	if board[nextRow][player.y] == "#" {
		return
	}
	player.x = nextRow
}

func (player *Player) moveUp(board [][]string) {
	nextRow := player.x - 1
	if nextRow < 0 || strings.TrimSpace(board[nextRow][player.y]) == "" {
		for row := len(board) - 1; row >= 0; row-- {
			if board[row][player.y] == "#" {
				return
			}
			if board[row][player.y] == "." {
				player.x = row
				return
			}
		}
	}
	if board[nextRow][player.y] == "#" {
		return
	}
	player.x = nextRow
}

func newPlayer(board [][]string) *Player {
	player := new(Player)
	player.x = 0
	for col, char := range board[0] {
		if char == "." {
			player.y = col
			break
		}
	}
	player.direction = ">"
	return player
}

func getPathDescription() []string {
	file, _ := os.ReadFile("inputPath.txt")
	line := string(file)
	path := make([]string, 0)
	li := 0
	for i, r := range line {
		if unicode.IsLetter(r) {
			path = append(path, line[li:i])
			path = append(path, string(r))
			li = i + 1
		}
	}
	path = append(path, line[li:])
	return path
}

func parseBoard() [][]string {
	file, _ := os.ReadFile("inputBoard.txt")
	lines := strings.Split(string(file), "\n")
	rows := len(lines)
	cols := 0
	for _, line := range lines {
		if len(line) > cols {
			cols = len(line)
		}
	}
	board := make([][]string, rows)
	for i := range board {
		board[i] = make([]string, cols)
	}
	for i, line := range lines {
		for j, r := range line {
			board[i][j] = string(r)
		}
	}
	return board
}
