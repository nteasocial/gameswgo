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

/*************************************************\
		BOARD STATE - FINDING A DIAGONAL

			----------------
			|  SUM   DIFF  |
			----------------

			0		1		2

0		0	0	|		|	2	2
		---------------------------
1				| 2	  0	|
		---------------------------
2		2	2	|		|	4	0

DIAGONAL 0:		x - y = 0
DIAGONAL 1:		x + y = 2	= len(rows|cols)+1
\*************************************************/

func NewBoard() *Board {
	state := State{
		Rows:  []int{0, 0, 0},
		Cols:  []int{0, 0, 0},
		Diags: []int{0, 0},
	}
	board := Board{
		Rows:  [][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}},
		State: &state,
	}
	return &board
}

func (board *Board) DoMove(player, row, col int) (int, error) {
	if board.Rows[row][col] != 0 {
		return 0, fmt.Errorf("there is already a piece there: %d", board.Rows[row][col])
	} else {
		// update board
		board.Rows[row][col] = player

		// update win state
		board.State.Rows[row] += player
		board.State.Cols[col] += player

		// update diagonals
		// if sum of coords is even it's on a diagonal
		// note: middle spot is on both diagonals
		if (row + col) % 2 == 0 {
			// diagonal 0
			if (row - col) == 0 {
				board.State.Diags[0] += player
			}
			// diagonal 1
			if (row + col) == 2 {
				board.State.Diags[1] += player
			}
		}

		// check for win
		for _, sum := range board.State.Rows {
			if sum == 3 || sum == -3 {
				return sum/3, nil
			}
		}
		for _, sum := range board.State.Cols {
			if sum == 3 || sum == -3 {
				return sum/3, nil
			}
		}
		for _, sum := range board.State.Diags {
			if sum == 3 || sum == -3 {
				return sum/3, nil
			}
		}
	}

	// no winner, no error
	return 0, nil
}

func (board *Board) Show() {
	fmt.Println("BOARD")
	for _, row := range board.Rows {
		fmt.Println(row)
	}
	fmt.Println("STATE")
	fmt.Println("R", board.State.Rows)
	fmt.Println("C", board.State.Cols)
	fmt.Println("D", board.State.Diags)
}
