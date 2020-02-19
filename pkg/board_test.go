package uttt

import "testing"

func TestBoardCheckNoWin(t *testing.T) {
	noWinBoard := Board{
		0, 0, 0,
		0, 0, 0,
		0, 0, 0,
	}
	if win := CheckWin(noWinBoard); win != 0 {
		t.Errorf("Expected 0 (no win), got %d", win)
	}
}

func TestBoardCheckWin(t *testing.T) {
	winBoard := Board{
		1, 0, 0,
		0, 1, 0,
		0, 0, 1,
	}
	if win := CheckWin(winBoard); win != Player1 {
		t.Errorf("Expected 1 (player 1 win), got %d",
			win)
	}
}

func TestFieldCheckNoWin(t *testing.T) {
	noWin := Board{0, 0, 0, 0, 0, 0, 0, 0, 0}
	noWinField := Field{
		noWin, noWin, noWin,
		noWin, noWin, noWin,
		noWin, noWin, noWin,
	}
	if win := CheckWin(noWinField); win != 0 {
		t.Errorf("Expected 0 (no win), got %d", win)
	}
}

func TestFieldCheckWin(t *testing.T) {
	noWin := Board{0, 0, 0, 0, 0, 0, 0, 0, 0}
	play1Win := Board{1, 0, 0, 0, 1, 0, 0, 0, 1}
	winField := Field{
		play1Win, play1Win, play1Win,
		noWin, noWin, noWin,
		noWin, noWin, noWin,
	}
	if win := CheckWin(winField); win != 1 {
		t.Errorf("Expected 1 (player 1 win), got %d",
			win)
	}
}
