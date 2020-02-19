package uttt

import (
	"fmt"
)

type Board [9]Mark

type Field [9]Board

type Move struct {
	Board int
	Spot  int
}

func (board Board) HasMarkAt(i, j int) Mark {
	return board[(i*3)+j]
}

func (field Field) HasMarkAt(i, j int) Mark {
	return CheckWin(field[(i*3)+j])
}

var boardMark = map[Mark]string{
	Unoccupied: " ",
	Player1:    "X",
	Player2:    "O",
}

func (board *Board) RenderASCII() []string {
	var render []string
	for i := 0; i < 3; i++ {
		render = append(render, "   |   |   ")
		render = append(render,
			fmt.Sprintf(" %s | %s | %s ",
				boardMark[board.HasMarkAt(i, 0)],
				boardMark[board.HasMarkAt(i, 1)],
				boardMark[board.HasMarkAt(i, 2)],
			))
		if i < 2 {
			render = append(render, "___|___|___")
		} else {
			render = append(render, "   |   |   ")
		}
	}
	return render
}

func (board *Board) GetMoves(boardIndex int) []*Move {
	moves := make([]*Move, 0, 9)
	if CheckWin(board) == Unoccupied {
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
	if CheckWin(field) == Unoccupied {
		nextBoardIndex := lastMove.Spot
		nextBoard := field[nextBoardIndex]

		if CheckWin(nextBoard) == Unoccupied {
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
