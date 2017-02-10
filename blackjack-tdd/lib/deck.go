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
func PrepDeck(nd int) Deck {
	d := CreateDeck()
	deck := Deck{d}
	decks := []Deck{}
	for i := 0; i < nd; i++ {
		decks = append(decks, deck)
	}
	deck = flatten(&decks)
	return deck
}

func flatten(d *[]Deck) Deck {
	var flatDeck []Card
	for _, deck := range *d {
		for _, card := range deck.Cards {
			flatDeck = append(flatDeck, card)
		}
	}
	return Deck{flatDeck}
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
