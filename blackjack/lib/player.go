package lib

import (
	"fmt"
	"strings"
)

// Player sets attributes for player
type Player struct {
	Name  string
	Hand  []Card
	Total int
	Stay  bool
}

// CreatePlayer creates a player
func CreatePlayer(name string, d *Deck) Player {
	hand := createHand(d)
	player := Player{name, hand, 0, false}
	player.HandValue()
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

// HitOrStay asks all players if they want to hti or stay
func (p *Player) HitOrStay(d *Deck, g *Game) {
	move := p.getResponse(g)
	switch move {
	case "hit":
		p.draw(d)
	case "stay":
		p.Stay = true
	default:
		fmt.Println("Sorry, I didn't catch that")
		p.HitOrStay(d, g)
	}
	p.HandValue()
}

func (p *Player) getResponse(g *Game) string {
	if p.Name == (*g).Players[0].Name {
		return playerPrompt()
	}
	return compPrompt(p)
}

func playerPrompt() string {
	var move string
	fmt.Println("Hit or Stay?")
	fmt.Scanln(&move)
	move = strings.ToLower(move)
	return move
}

// CompHitOrStay decides whether computer should hit or stay
func compPrompt(p *Player) string {
	if p.Total < 16 {
		return "hit"
	}
	return "stay"
}

func (p *Player) draw(d *Deck) {
	var card Card
	d.Cards, card = d.Cards[:len(d.Cards)-1], d.Cards[len(d.Cards)-1]
	p.Hand = append(p.Hand, card)
}

// ShowHand displays the Face rather than Values
func (p *Player) ShowHand() []string {
	hand := []string{}
	for _, c := range p.Hand {
		card := "{ " + c.Suit + " " + c.Face + " }"
		hand = append(hand, card)
	}
	return hand
}
