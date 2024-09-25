package tic_tac_toe

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/nteasocial/gameswgo/gameutils"
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
	move := strings.Fields(input)
	x, _ := strconv.Atoi(move[0])
	y, _ := strconv.Atoi(move[1])
	fmt.Printf("REQUESTED POSITION: %d %d\n", x, y)
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
