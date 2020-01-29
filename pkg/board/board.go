package board

import "fmt"

type Board [9]int

type Field [9]Board

type Move struct {
	Board int
	Spot  int
}

func (field *Field) Copy() *Field {
	orgBoard := *field
	copyBoard := orgBoard
	return &copyBoard
}

func (field *Field) Display() {
}

func (board *Board) Display() {
	temp := make(map[int]string)
	temp[-1] = "O"
	temp[1] = "X"
	temp[0] = " "

	for _, s := range *board {
		fmt.Println(s)
	}
}

func (move *Move) Copy() *Move {
	return &Move{move.Board, move.Spot}
}

func (board *Board) CheckWin() int {
	columnSum, rowSum := 0, 0
	for ii := 0; ii < 3; ii++ {
		for jj := 0; jj < 3; jj++ {
			columnSum += board[3*ii+jj]
			rowSum += board[ii+3*jj]
		}
		switch columnSum {
		case 3:
			return 1
		case -3:
			return -1
		}
		switch rowSum {
		case 3:
			return 1
		case -3:
			return -1
		}
		columnSum, rowSum := 0, 0

	}
	switch board[0] + board[4] + board[8] {
	case 3:
		return 1
	case -3:
		return -1
	}
	switch board[2] + board[4] + board[6] {
	case 3:
		return 1
	case -3:
		return -1
	}

	return 0
}

func (field *Field) CheckWin() int {
	var colSum, rowSum int
	for ii := 0; ii < 3; ii++ {
		for jj := 0; jj < 3; jj++ {
			colSum += field[3*ii+jj].CheckWin()
			rowSum += field[ii+3*jj].CheckWin()
		}
		switch colSum {
		case 3:
			return 1
		case -3:
			return -1
		}
		switch rowSum {
		case 3:
			return 1
		case -3:
			return -1
		}
		rowSum, colSum = 0, 0
	}

	switch field[0].CheckWin() + field[4].CheckWin() + field[8].CheckWin() {
	case 3:
		return 1
	case -3:
		return -1
	}
	switch field[2].CheckWin() + field[4].CheckWin() + field[6].CheckWin() {
	case 3:
		return 1
	case -3:
		return -1
	}

	return 0
}

func (board *Board) GetMoves(boardIndex int) []*Move {
	moves := make([]*Move, 0, 9)
	if board.CheckWin() == 0 {
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
	if field.CheckWin() == 0 {
		nextBoardIndex := lastMove.Spot
		nextBoard := field[nextBoardIndex]

		if nextBoard.CheckWin() == 0 {
			moves = append(moves, nextBoard.GetMoves(nextBoardIndex)...)
		} else {
			for index, board := range field {
				if board.CheckWin() == 0 {
					moves = append(moves, board.GetMoves(index)...)
				}
			}
		}
	}
	return moves
}

func (field *Field) MakeMove(move *Move, player int) {
	field[move.Board][move.Spot] = player
}
