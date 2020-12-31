package tic_tac_toe

type AI interface {
	PickMove(game *Game) (int, int)
}

type ShittyAI struct {
	Me rune
}

func (ai ShittyAI) PickMove(game *Game) (int, int) {
	panic("IMPLEMENT ME PLASE")
}