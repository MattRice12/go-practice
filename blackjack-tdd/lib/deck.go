package lib

import (
	"math/rand"
	"time"
)

// Deck is a set of cards
type Deck struct {
	Cards []Card
}

// PrepDeck turns a Card slice into a Deck
func PrepDeck() Deck {
	d := CreateDeck()
	deck := Deck{d}
	return deck
}

// CreateDeck creates deck
func CreateDeck() []Card {
	deck := []Card{}
	suits := [4]string{"Hearts", "Diamonds", "Clubs", "Spades"}
	for _, s := range suits {
		for i := 2; i <= 14; i++ {
			card := Card{s, i}
			deck = append(deck, card)
		}
	}
	return deck
}

func shuffle(d []Card) []Card {
	for i := 0; i < 52; i++ {
		randCard := random(&d)
		d[i], d[randCard] = d[randCard], d[i]
	}
	return d
}

func random(d *[]Card) int {
	rand.Seed(int64(time.Now().Nanosecond()))
	return rand.Intn(len(*d))
}
