package main

import "github.com/nwunderly/terminal-games/blackjack"

func main() {
	game := blackjack.Blackjack()
	game.Run()
}
