package main

import (
	"fmt"
	"strings"
	uttt "uttt/pkg"
)

func main() {
	fmt.Println("Let's play a game of ultimate tic tac toe!")
	board := uttt.Board{
		0, 0, 1,
		1, 2, 1,
		0, 1, 2,
	}
	fmt.Println(strings.Join(board.RenderASCII(), "\n"))
}
