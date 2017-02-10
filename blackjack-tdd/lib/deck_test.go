package lib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeck(t *testing.T) {
	card1 := Card{Suit: "Hearts", Value: 2}
	card2 := Card{Suit: "Diamonds", Value: 10}
	cards := []Card{card1, card2}
	deck := Deck{cards}

	if assert.NotNil(t, deck) {
		assert.Equal(t, []Card{Card{"Hearts", 2}, Card{"Diamonds", 10}}, deck.Cards)
	}
}

func TestCreateDeck(t *testing.T) {
	d := CreateDeck()
	if assert.NotNil(t, d) {
		if len(d) != 52 {
			t.Errorf("Expected 52 cards; got %d cards.", len(d))
		}
		if d[0].Suit != "Hearts" || d[0].Value == 0 {
			t.Errorf("Expected 2 of Hearts; got %d of %v", d[0].Value, d[0].Suit)
		}
	}
}

func TestPrepMultipleDecks(t *testing.T) {
	numDecks := 3
	d := PrepDeck(numDecks)
	if assert.NotNil(t, d) {
		assert.Equal(t, 156, len(d.Cards))
	}
}
