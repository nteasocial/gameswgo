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
	fmt.Println("You chose to end your turn. Beginning dealer turn.")
	time.Sleep(time.Millisecond*500)

	for game.DealerPlaying {
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
	time.Sleep(time.Millisecond*500)
}

func (game *Game) DoDealerTurn() {
	score := game.DealerScore()
	fmt.Printf("Dealer score is %d.\n", score)
	time.Sleep(time.Millisecond*500)
	if score <= 17 { // hit
		fmt.Println("Dealer has decided to hit.")
		time.Sleep(time.Millisecond*500)
		game.DealerCards = game.Deck.Hit(game.DealerCards)
		score = game.DealerScore()
		if score > 21 {
			fmt.Println("Dealer busts! Their score was", score)
			time.Sleep(time.Millisecond*500)
			game.DealerPlaying = false
			return
		}
	} else {
		fmt.Println("Dealer has decided to stay. Dealer turn is over.")
		time.Sleep(time.Millisecond*500)
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
