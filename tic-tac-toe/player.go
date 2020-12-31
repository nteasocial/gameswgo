package tic_tac_toe

import (
	"github.com/nwunderly/terminal-games/gameutils"
	"strconv"
	"strings"
)

type Player interface {
	Human() bool // false if bot
	GetTurn() (int, int)
}

type HumanPlayer struct {
	Player
	GamePiece rune
	Game      *Game
}

func (p *HumanPlayer) Human() bool { return true }

func (p *HumanPlayer) GetTurn() (int, int) {
	input := gameutils.Input("Input move as two integers: ")
	move := strings.Split(input, " ")
	x, y := func(s []string) (int, int) {
		a, _ := strconv.Atoi(s[0])
		b, _ := strconv.Atoi(s[0])
		return a, b
	}(move)
	return x, y
}

type BotPlayer struct {
	Player
	GamePiece string
	Game      *Game
	AI        AI
}

func (p *BotPlayer) Human() bool { return false }

func (p *BotPlayer) GetTurn() (int, int) {
	return p.AI.PickMove(p.Game)
}
