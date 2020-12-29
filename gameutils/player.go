package gameutils

type Player interface {
	Name() string
	Human() bool
}

type HumanPlayer struct {
	name  string
	human bool
}

func (h HumanPlayer) Name() string {
	return h.name
}

func (h HumanPlayer) Human() bool {
	return h.human
}
