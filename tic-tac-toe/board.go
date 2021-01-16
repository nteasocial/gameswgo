package tic_tac_toe

import (
	"fmt"
)

type Board struct{
	Rows [][]int
	State *State
}

type State struct {
	Rows []int
	Cols []int
	Diags[]int
}

func NewBoard() *Board {
	state := State{
		Rows:  []int{0, 0, 0},
		Cols:  []int{0, 0, 0},
		Diags: []int{0, 0, 0},
	}
	board := Board{
		Rows:  [][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}},
		State: &state,
	}
	return &board
}

func (board *Board) DoMove(player, row, col int) (int, error) {
	if board.Rows[row][col] != 0 {
		return 0, fmt.Errorf("there is already a piece there")
	} else {
		board.Rows[row][col] = player

		// update win state
		board.State.Rows[row] += player
		board.State.Cols[col] += player

		// update diagonals
		// if sum of coords is even it's on a diagonal
		if (row + col) % 2 == 0 {
			// not middle box
			if row != 1 {
				// sum will be 0 or 2, making diag 0 or 1
				diag := (row + col) / 2
				board.State.Diags[diag] += player
			}
		}
	}

	// todo: check win conditions

	return 0, nil
}
