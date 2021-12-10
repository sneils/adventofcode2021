package bingo

import "testing"

func getSampleBoard() Board {
	sample := []string{
		"22 13 17 11  0",
		" 8  2 23  4 24",
		"21  9 14 16  7",
		" 6 10  3 18  5",
		" 1 12 20 15 19",
	}
	return New(sample)
}

func TestMark(t *testing.T) {
	board := getSampleBoard()
	pick := "22"
	if !board.Mark(pick) {
		t.Errorf("Expected to find %s, but didnt :(", pick)
	}
	expected := MARKER
	if board[0][0] != expected {
		t.Errorf("Expected Board[0][0] to be %s (marked), but was %s", expected, board[0][0])
	}
}

func TestHasWonHorizontally(t *testing.T) {
	horizontalBoard := getSampleBoard()
	horizontalWin := []string{"1", "12", "20", "15", "19"}
	for _, w := range horizontalWin {
		horizontalBoard.Mark(w)
	}
	if !horizontalBoard.HasWon() {
		t.Error("Expected board to be a winner horizontally, but was not :(")
	}
}

func TestHasWonVertically(t *testing.T) {
	verticalBoard := getSampleBoard()
	verticalWin := []string{"0", "24", "7", "5", "19"}
	for _, w := range verticalWin {
		verticalBoard.Mark(w)
	}
	if !verticalBoard.HasWon() {
		t.Error("Expected board to be a winner vertically, but was not :(")
	}
}

func TestGetScore(t *testing.T) {
	board := getSampleBoard()
	expected := 300
	score := board.GetScore()
	if score != expected {
		t.Errorf("Expected score to be %d, but was %d", expected, score)
	}
}
