package lib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCard(t *testing.T) {
	var card Card
	card.Suit = "Hearts"
	card.Value = 2
	assert.Equal(t, Card{Suit: "Hearts", Value: 2}, card)
}
