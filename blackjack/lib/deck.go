package lib

import (
	"math/rand"
	"time"
)

// Deck holds all the Cards
type Deck struct {
	Cards []Card
}

// PrepDeck sets value for `deck` -- creates and shuffles a deck
func PrepDeck() Deck {
	d := createDeck()
	d = shuffleDeck(d)
	deck := Deck{d}
	return deck
}

func createDeck() []Card {
	// Create empty Slice for type Card
	deck := []Card{}
	suits := [4]string{"Heart", "Diamond", "Club", "Spade"}
	for _, suit := range suits {
		for value := 2; value <= 14; value++ {
			face := StringifyValue(value)
			card := Card{suit, value, face}
			deck = append(deck, card)
		}
	}
	return deck
}

// shuffleDeck replaces card at position[j] with a random card; j++
func shuffleDeck(deck []Card) []Card {
	for j := 0; j < 51; j++ {
		randCard := random(deck)
		deck[j], deck[randCard] = deck[randCard], deck[j]
	}

	return deck
}

// random selects a random card in the Card Slice
func random(d []Card) int {
	rand.Seed(int64(time.Now().Nanosecond()))
	return rand.Intn(len(d))
}
