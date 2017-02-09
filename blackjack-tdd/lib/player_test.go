package lib

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateHand(t *testing.T) {
	deck := PrepDeck()
	name := "Player" + strconv.Itoa(1)
	player := Player{name, []Card{}, 0, false}
	player.CreateHand(&deck)
	if len(player.Hand) != 2 {
		t.Fatalf("Expected 2 cards; got %v cards", len(player.Hand))
	}
	assert.Equal(t, 2, len(player.Hand))
}

func TestCreatePlayer(t *testing.T) {
	assert := assert.New(t)

	deck := PrepDeck()
	player := CreatePlayer(&deck, 1)
	if player.Name != "Player1" {
		t.Fatalf("Expected Player Name to be `Player1`; got `%v`.", player.Name)
	}
	assert.Equal("Player1", player.Name)

	if len(player.Hand) != 2 {
		t.Fatalf("Expected Player1 to have 2 cards; got `%v` card.", len(player.Hand))
	}
	assert.Equal(2, len(player.Hand))

	if player.Stay != false {
		t.Fatalf("Expected player.Stay to be false; it was %v", player.Stay)
	}
	assert.False(player.Stay, "player.Stay is true; should be false.")
}

func TestPrepPlayers(t *testing.T) {
	deck := PrepDeck()
	players := PrepPlayers(&deck)
	if len(players) != 4 {
		t.Errorf("Expected 4 players; got %v palyers", len(players))
	}
	assert.Equal(t, 4, len(players))
	assert.Equal(t, 44, len(deck.Cards), "4 players added; 8 cards drawn; 44 should be remaining.")
}

func TestDraw(t *testing.T) {
	deck := PrepDeck()
	p := CreatePlayer(&deck, 1)
	p.Draw(&deck)
	assert.Equal(t, 3, len(p.Hand), "After first draw, player should have 3 cards")
	assert.Equal(t, 49, len(deck.Cards))

	p.Draw(&deck)
	assert.Equal(t, 4, len(p.Hand), "After second draw, player should have 4 cards")
	assert.Equal(t, 48, len(deck.Cards))
}
