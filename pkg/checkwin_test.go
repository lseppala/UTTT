package uttt

import "testing"

type FakeNeverWin struct{}

func (fake FakeNeverWin) HasMarkAt(i, j int) Mark {
	return Unoccupied
}

func TestNoCheckWin(t *testing.T) {
	fake := FakeNeverWin{}
	if win := CheckWin(fake); win != Unoccupied {
		t.Errorf("Expected 0 (no win), got %d", win)
	}
}

type FakeDiagWin struct{}

func (fake FakeDiagWin) HasMarkAt(i, j int) Mark {
	if i == j {
		return Player1
	} else {
		return Unoccupied
	}
}

func TestDiagCheckWin(t *testing.T) {
	fake := FakeDiagWin{}
	if win := CheckWin(fake); win != Player1 {
		t.Errorf("Expected Player 1, got %d", win)
	}
}
