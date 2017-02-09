package lib

import (
	"fmt"
	"strings"
)

// Game contains all players
type Game struct {
	Players []Player
}

// PrepPlayers sets value for `game` -- prepares the game
func PrepPlayers(d *Deck) Game {
	player1 := CreatePlayer("Matt", d)
	player2 := CreatePlayer("Hal 9000", d)
	players := []Player{player1, player2}
	game := Game{players}
	return game
}

// PlayRound (1) shows scores, (2) checks for a bust, (3)
func (g *Game) PlayRound(d *Deck) {
	g.showRound()
	if checkBust(g) || checkStay(g) {
		return
	}
	for i := range g.Players {
		g.Players[i].playerTurn(d, g)
	}
	g.PlayRound(d)
}

func checkBust(g *Game) bool {
	for _, p := range g.Players {
		if p.Bust() {
			g.showResults()
			fmt.Println("Player " + p.Name + " Busts!")
			return true
		}
	}
	return false
}

// Bust checks if the player's card values are over 21
func (p Player) Bust() bool {
	if p.Total > 21 {
		return true
	}
	return false
}

func checkStay(g *Game) bool {
	if (g.Players[0].Stay != true) || (g.Players[1].Stay != true) {
		return false
	}
	findWinner(g)
	return true
}

func findWinner(g *Game) {
	g.showResults()
	if g.Players[0].Total > g.Players[1].Total {
		fmt.Println(g.Players[0].Name, "Wins!")
	} else {
		fmt.Println(g.Players[1].Name, "Wins!")
	}
}

func (p *Player) playerTurn(d *Deck, g *Game) {
	g.showRound()
	if p.Stay != true {
		p.HitOrStay(d, g)
	}
}

// HandValue checks total value of player's hand
func (p *Player) HandValue() {
	total := 0
	for _, card := range p.Hand {
		total += card.Value
	}
	p.Total = total
}

// showResults displays results at beginning of each round
func (g *Game) showRound() {
	print("\033[H\033[2J")
	for _, p := range g.Players {
		if p.Name == g.Players[0].Name {
			p.revealedCards()
		} else {
			fmt.Println(p.Name+":", p.ShowHand()[0], strings.Repeat("{ Hidden }", len(p.Hand)-1))
			fmt.Println("Total: ?")
		}
		fmt.Println("")
	}
}

// showResults displays final results
func (g *Game) showResults() {
	print("\033[H\033[2J")
	for _, p := range g.Players {
		p.revealedCards()
		fmt.Println("")
	}
}

func (p *Player) revealedCards() {
	fmt.Print(p.Name + ": ")
	for i := range p.Hand {
		fmt.Print(p.ShowHand()[i] + " ")
	}
	fmt.Println("\nTotal:", p.Total)
}
