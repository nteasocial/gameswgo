package gameutils

type Game interface {
	Name() string
	Players() []*Player
	DoTurn()
	CheckWin() *Player
	Run()
}

type GameHelper struct {
	name    string
	players []*Player
}

func (g GameHelper) Name() string {
	return g.name
}

func (g GameHelper) Players() []*Player {
	return g.players
}

func (g GameHelper) DoTurn() {
	panic("IMPLEMENT ME!")
}

func (g GameHelper) CheckWin() {
	panic("IMPLEMENT ME!")
}

func (g GameHelper) Run() {
	panic("IMPLEMENT ME!")
}
