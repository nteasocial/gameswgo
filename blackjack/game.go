package blackjack

import (
	"fmt"
	"github.com/nwunderly/terminal-games/gameutils"
	"math/rand"
	"time"
)

type Game struct {
	PlayerCards   []string
	DealerCards   []string
	PlayerPlaying bool
	DealerPlaying bool
	Deck          *Deck
}

func New() *Game {
	return &Game{
		PlayerCards:   []string{},
		DealerCards:   []string{},
		PlayerPlaying: true,
		DealerPlaying: true,
		Deck:          NewDeck(),
	}
}

func (game *Game) Run() {
	rand.Seed(time.Now().UnixNano())
	game.Deck.Shuffle()

	for game.PlayerPlaying {
		game.DoPlayerTurn()
	}
	if game.PlayerScore() > 21 {
		gameutils.Print("You lose!")
		return
	}
	gameutils.Print("You chose to end your turn. Beginning dealer turn.")

	for game.DealerPlaying {
		game.DoDealerTurn()
	}
	if game.DealerScore() > 21 {
		gameutils.Print("You win!")
		return
	}

	fmt.Printf("Game has ended. Final score:\n    Player: %d\n    Dealer: %d\n",
		game.PlayerScore(), game.DealerScore())

	switch game.Winner() {
	case "player":
		gameutils.Print("You win!")
	case "dealer":
		gameutils.Print("Dealer wins!")
	default:
		gameutils.Print("Draw.")
	}
}

func (game *Game) DoPlayerTurn() {
	score := game.PlayerScore()
	hit := gameutils.BoolInput(fmt.Sprintf("Your score is %d. Hit? (y/n): ", score))
	if hit {
		game.PlayerCards = game.Deck.Hit(game.PlayerCards)
		score = game.PlayerScore()
		if score > 21 {
			gameutils.Print("Bust! Your score was %d.", score)
			game.PlayerPlaying = false
			return
		}
	} else {
		game.PlayerPlaying = false
	}
}

func (game *Game) DoDealerTurn() {
	score := game.DealerScore()
	gameutils.Print("Dealer score is %d.", score)
	if score <= 5 || (score < 21 && score <= game.PlayerScore()) { // hit
		gameutils.Print("Dealer has decided to hit.")
		game.DealerCards = game.Deck.Hit(game.DealerCards)
		score = game.DealerScore()
		if score > 21 {
			gameutils.Print("Dealer busts! Their score was %d.", score)
			game.DealerPlaying = false
			return
		}
	} else {
		gameutils.Print("Dealer has decided to stay. Dealer turn is over.")
		game.DealerPlaying = false
	}
}

func (game *Game) PlayerScore() int {
	return Score(game.PlayerCards)
}

func (game *Game) DealerScore() int {
	return Score(game.DealerCards)
}

func (game *Game) Winner() string {
	playerScore, dealerScore := game.PlayerScore(), game.DealerScore()

	if playerScore > dealerScore {
		return "player"
	} else if dealerScore > playerScore {
		return "dealer"
	} else {
		return ""
	}
}
