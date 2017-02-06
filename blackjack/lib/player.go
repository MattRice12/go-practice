package lib

import (
	"fmt"
	"strings"
)

// Player sets attributes for player
type Player struct {
	Name string
	Hand []Card
}

// CreatePlayer creates a player
func CreatePlayer(name string, d *Deck) Player {
	hand := createHand(d)
	player := Player{name, hand}
	return player
}

func createHand(d *Deck) []Card {
	hand := []Card{}
	var card Card
	for i := 0; i < 2; i++ {
		d.Cards, card = d.Cards[:len(d.Cards)-1], d.Cards[len(d.Cards)-1]
		hand = append(hand, card)
	}
	return hand
}

// HitOrStay Prompts player to hit
func HitOrStay(p *Player, d *Deck) {
	var move string
	fmt.Println("Hit or Stay?")
	fmt.Scanln(&move)
	move = strings.ToLower(move)
	if move == "hit" {
		draw(p, d)
		fmt.Println("Player 1: ", p.Hand)
		HitOrStay(p, d)
	} else if move == "stay" {
		return
	} else {
		fmt.Println("Sorry, I didn't catch that")
		HitOrStay(p, d)
	}
}

func draw(p *Player, d *Deck) {
	var card Card
	d.Cards, card = d.Cards[:len(d.Cards)-1], d.Cards[len(d.Cards)-1]
	p.Hand = append(p.Hand, card)
}
