package blackjack

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/nteasocial/gameswgo/gameutils"
)

type Deck struct {
	cards []string
}

func NewDeck() *Deck {
	deck := &Deck{
		cards: []string{
			"ace", "ace",
			"2", "2",
			"3", "3",
			"4", "4",
			"5", "5",
			"6", "6",
			"7", "7",
			"8", "8",
			"9", "9",
			"10", "10",
		},
	}
	deck.Shuffle()
	return deck
}

func (deck *Deck) Shuffle() {
	gameutils.Shuffle(deck.cards)
}

func (deck *Deck) Hit(hand []string) []string {
	rand.Seed(time.Now().UnixNano())
	card, cards := gameutils.Pop(deck.cards)
	deck.cards = cards
	hand = append(hand, card)
	fmt.Println("    Card drawn: ", card)
	fmt.Println("    New hand: ", hand)
	return hand
}

func Score(hand []string) int {
	score := 0
	aces := 0
	// handle int value cards
	for i := 0; i < len(hand); i++ {
		if hand[i] == "ace" {
			aces += 1
		} else {
			n, err := strconv.Atoi(hand[i])
			// this shouldn't ever happen
			if err != nil {
				panic(err)
			}
			score += n
		}
	}
	// handle Aces
	for j := 0; j < aces; j++ {
		toCompare := 21 - (aces - 1)
		if score+11 < toCompare {
			score += 11
		} else {
			score += 1
		}
	}
	return score
}
