package main

import (
	"fmt"
	"github.com/akamensky/argparse"
	"github.com/nwunderly/terminal-games/blackjack"
	"github.com/nwunderly/terminal-games/tic-tac-toe"
	"os"
)

func main() {
	parser := argparse.NewParser("terminal-games", "")
	game := parser.String("g", "game", &argparse.Options{
		Required: false,
		Validate: nil,
		Help:     "",
		Default:  nil,
	})

	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Println(err)
		return
	}

	switch *game {
	case "blackjack":
		blackjack.New().Run()
	case "tic-tac-toe":
	case "ttt":
		tic_tac_toe.New().Run()
	default:
	}

}
