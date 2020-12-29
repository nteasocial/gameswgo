package blackjack

import (
	"fmt"
	"github.com/nwunderly/terminal-games/gameutils"
	"math/rand"
	"time"
)

type Game struct {
	PlayerCards []string
	DealerCards []string
	PlayerPlaying bool
	DealerPlaying bool
	Deck *Deck
}

func Blackjack() *Game {
	return &Game{
		PlayerCards: []string{},
		DealerCards: []string{},
		PlayerPlaying: true,
		DealerPlaying: true,
		Deck: NewDeck(),
	}
}

func (game *Game) Run() {
	rand.Seed(time.Now().UnixNano())
	game.Deck.Shuffle()

	for game.PlayerPlaying {
		game.DoPlayerTurn()
	}
	if game.PlayerScore() > 21 {
		fmt.Println("You lose!")
		return
	}

	for game.DealerPlaying {
		time.Sleep(time.Millisecond*500)
		game.DoDealerTurn()
	}
	if game.DealerScore() > 21 {
		fmt.Println("You win!")
		return
	}

	switch game.Winner() {
	case "player":
		fmt.Println("You win!")
	case "dealer":
		fmt.Println("Dealer wins!")
	default:
		fmt.Println("Draw.")
	}
}

func (game *Game) DoPlayerTurn() {
	score := game.PlayerScore()
	hit := gameutils.BoolInput(fmt.Sprintf("Your score is %d. Hit? (y/n): ", score))
	if hit {
		game.PlayerCards = game.Deck.Hit(game.PlayerCards)
		score = game.PlayerScore()
		if score > 21 {
			fmt.Println("Bust! Your score was", score)
			game.PlayerPlaying = false
			return
		}
	} else {
		game.PlayerPlaying = false
	}
}

func (game *Game) DoDealerTurn() {
	score := game.DealerScore()
	fmt.Printf("Dealer score is %d.\n", score)
	if score <= 17 { // hit
		fmt.Println("Dealer has decided to hit.")
		game.DealerCards = game.Deck.Hit(game.DealerCards)
		score = game.DealerScore()
		if score > 21 {
			fmt.Println("Dealer busts! Their score was", score)
			game.DealerPlaying = false
			return
		}
	} else {
		fmt.Println("Dealer has decided to stay. Dealer turn is over.")
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
