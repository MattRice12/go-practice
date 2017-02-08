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

func playerHitPrompt(p *Player, d *Deck) string {
	total := HandValue(p)
	fmt.Println(p.Name, "total:", total)
	var move string
	fmt.Println("\nHit or Stay?")
	fmt.Scanln(&move)
	move = strings.ToLower(move)
	return move
}

// HitOrStay asks all players if they want to hti or stay
func HitOrStay(p *Player, d *Deck, g *Game) (Player, *Deck, *Game, string) {
	var move string
	if p.Name == g.Players[0].Name {
		move = playerHitPrompt(p, d)
	} else {
		move = compHitOrStay(p, d)
	}
	if move == "hit" {
		draw(p, d)
	} else if move == "stay" {
		return *p, d, g, move
	} else {
		fmt.Println("Sorry, I didn't catch that")
		HitOrStay(p, d, g)
	}
	return *p, d, g, move
}

// CompHitOrStay decides whether computer should hit or stay
func compHitOrStay(p *Player, d *Deck) string {
	total := HandValue(p)
	if total < 16 {
		return "hit"
	}
	return "stay"
}

func draw(p *Player, d *Deck) {
	var card Card
	d.Cards, card = d.Cards[:len(d.Cards)-1], d.Cards[len(d.Cards)-1]
	p.Hand = append(p.Hand, card)
}
