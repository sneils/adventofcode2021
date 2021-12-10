package bingo

import (
	"strings"

	"github.com/sneils/adventofcode2021/convert"
)

const MARKER = "x"

type Board [][]string

func New(input []string) Board {
	board := Board{}
	for _, line := range input {
		fields := strings.Fields(line)
		board = append(board, fields)
	}
	return board
}

func (board Board) Mark(n string) bool {
	found := false
	for y, row := range board {
		for x, col := range row {
			if col == n {
				board[y][x] = MARKER
				found = true
			}
		}
	}
	return found
}

func allEqual(input []string, marker string) bool {
	return strings.Join(input, "") == strings.Repeat(marker, len(input))
}

func (board Board) HasWon() bool {
	for _, row := range board {
		if allEqual(row, MARKER) {
			return true
		}
	}
	for x := 0; x < len(board); x++ {
		col := []string{}
		for _, row := range board {
			col = append(col, row[x])
		}
		if allEqual(col, MARKER) {
			return true
		}
	}
	return false
}

func (board Board) GetScore() int {
	score := 0
	for _, row := range board {
		for _, col := range row {
			if col != MARKER {
				score += convert.ToInt(col)
			}
		}
	}
	return score
}
