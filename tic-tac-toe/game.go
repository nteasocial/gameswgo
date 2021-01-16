package tic_tac_toe

import "fmt"

type Game struct {
	Players []Player
	Board   *Board
}

func New(/*twoPlayer bool*/) *Game {
	game := &Game{
		Players: nil,
		Board: nil,
	}

	o := &HumanPlayer{
		GamePiece: 'o',
		Game: nil,
	}
	x := &HumanPlayer{
		GamePiece: 'x',
		Game: nil,
	}

	players := []Player{x, o}
	game.Players = players

	game.Board = NewBoard()

	return game
}

func (game *Game) Run() {
	fmt.Println("RUN SUCCESSFULLY CALLED.")
}
