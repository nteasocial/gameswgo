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
		Game: game,
	}
	x := &HumanPlayer{
		GamePiece: 'x',
		Game: game,
	}

	players := []Player{x, o}
	game.Players = players

	game.Board = NewBoard()

	return game
}

func (game *Game) Run() {
	fmt.Println("RUN SUCCESSFULLY CALLED.")

	var turns int

	for {

		for i, player := range game.Players {
			fmt.Printf("----- BEGIN PLAYER %d TURN %d -----\n", i, turns)

			game.Board.Show()

			row, col := player.GetTurn()

			var playerNum int

			switch i {
			case 0:
				playerNum = -1
			case 1:
				playerNum = 1
			default:
				fmt.Println("ERROR: player id ", i)
				return
			}

			winner, err := game.Board.DoMove(playerNum, row, col)

			if err != nil {
				fmt.Println(err)
				return
			}

			if winner != 0 {
				fmt.Printf("Player %d WINS!\n", winner)
				game.Board.Show()
				return
			}

			fmt.Printf("----- END PLAYER %d TURN %d -----\n", i, turns)
		}

		turns++

	}

}
