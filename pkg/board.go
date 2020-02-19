package uttt

import "fmt"

type Board [9]Mark

type Field [9]Board

type Move struct {
	Board int
	Spot  int
}

type Mark int

const (
	Unoccupied Mark = iota
	Player1
	Player2
)

type HasMark interface {
	HasMarkAt(i, j int) Mark
}

func (board Board) HasMarkAt(i, j int) Mark {
	return board[(i*3)+j]
}

func (field Field) HasMarkAt(i, j int) Mark {
	return CheckWin(field[(i*3)+j])
}

func CheckWin(marked HasMark) Mark {
	for i := 0; i < 3; i++ {
		if mark := marked.HasMarkAt(i, 0); mark == marked.HasMarkAt(i, 1) &&
			mark == marked.HasMarkAt(i, 0) {
			return mark
		}
		if mark := marked.HasMarkAt(0, i); mark == marked.HasMarkAt(1, i) &&
			mark == marked.HasMarkAt(2, i) {
			return mark
		}
		if mark := marked.HasMarkAt(i, i); mark == marked.HasMarkAt(i, i) &&
			mark == marked.HasMarkAt(i, i) {
			return mark
		}
		if mark := marked.HasMarkAt(3-i, 3-i); mark == marked.HasMarkAt(3-i, 3-i) &&
			mark == marked.HasMarkAt(3-i, 3-i) {
			return mark
		}
	}
	return Unoccupied
}

func (field *Field) Copy() *Field {
	orgBoard := *field
	copyBoard := orgBoard
	return &copyBoard
}

func (field *Field) Display() {
}

func (board *Board) Display() {
	temp := make(map[Mark]string)
	temp[Unoccupied] = " "
	temp[Player1] = "X"
	temp[Player2] = "O"

	for _, s := range *board {
		fmt.Println(s)
	}
}

func (move *Move) Copy() *Move {
	return &Move{move.Board, move.Spot}
}

func (board *Board) GetMoves(boardIndex int) []*Move {
	moves := make([]*Move, 0, 9)
	if CheckWin(board) == 0 {
		for index, spot := range board {
			if spot == 0 {
				newMove := &Move{boardIndex, index}
				moves = append(moves, newMove)
			}
		}
	}
	return moves
}

func (field *Field) GetMoves(lastMove *Move) []*Move {
	moves := make([]*Move, 0, 81)
	if CheckWin(field) == 0 {
		nextBoardIndex := lastMove.Spot
		nextBoard := field[nextBoardIndex]

		if CheckWin(nextBoard) == 0 {
			moves = append(moves, nextBoard.GetMoves(nextBoardIndex)...)
		} else {
			for index, board := range field {
				if CheckWin(board) == 0 {
					moves = append(moves, board.GetMoves(index)...)
				}
			}
		}
	}
	return moves
}

func (field *Field) MakeMove(move *Move, player Mark) {
	field[move.Board][move.Spot] = player
}
