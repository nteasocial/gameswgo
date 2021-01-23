package main

import (
	"fmt"
	"github.com/nwunderly/terminal-games/blackjack"
	"github.com/nwunderly/terminal-games/gameutils"
	"github.com/nwunderly/terminal-games/tic-tac-toe"
	"os"
)

func main() {
	var game string

	if len(os.Args) <= 1 {
		game = gameutils.Input("Pick a game: ")
	} else {
		game = os.Args[1]
	}

	switch game {
	case "blackjack":
		blackjack.New().Run()
	case "tic-tac-toe":
	case "ttt":
		tic_tac_toe.New().Run()
	default:
		panic(fmt.Errorf("INVALID GAME NAME: %s", game))
	}
}
