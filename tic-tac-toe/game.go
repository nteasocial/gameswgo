package tic_tac_toe

import "fmt"

type Game struct {
	Players map[string]Player
	Board   *Board
}

func New(/*twoPlayer bool*/) *Game {
	o := new(HumanPlayer)
	x := new(HumanPlayer)

	players := map[string]Player{
		"X": x,
		"O": o,
	}

	game := &Game{
		Players: players,
	}

	return game
}

func (game *Game) Run() {
	fmt.Println("RUN SUCCESSFULLY CALLED.")
}
