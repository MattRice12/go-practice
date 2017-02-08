package lib

import (
	"fmt"
	"strings"
)

// Game contains all players
type Game struct {
	Players []Player
}

// PrepPlayers prepares the game
func PrepPlayers(d *Deck) Game {
	player1 := CreatePlayer("Matt", d)
	player2 := CreatePlayer("Hal 9000", d)
	players := []Player{player1, player2}
	game := Game{players}
	return game
}

// PlayRound plays the round
func PlayRound(d *Deck, g *Game) {
	showRound(g)
	if checkBust(g) {
		return
	}
	players := []Player{}
	moves := []string{}
	for _, p := range g.Players {
		var m string
		p, d, g, m = HitOrStay(&p, d, g)
		moves = append(moves, m)
		players = append(players, p)
	}
	game := Game{players}
	if moves[0] == moves[1] {
		findWinner(g)
		return
	}
	PlayRound(d, &game)
}

func findWinner(g *Game) {
	var finalCount []int
	for _, p := range g.Players {
		total := HandValue(&p)
		fmt.Println(p.Name+"'s'", "total:", total)
		finalCount = append(finalCount, total)
	}
	if finalCount[0] > finalCount[1] {
		fmt.Println(g.Players[0].Name, "Wins!")
	} else {
		fmt.Println(g.Players[1].Name, "Wins!")
	}
}

func showRound(g *Game) {
	print("\033[H\033[2J")
	for _, p := range g.Players {
		if p.Name == g.Players[0].Name {
			fmt.Println(p.Name+":", p.Hand)
		} else {
			fmt.Println(p.Name+":", p.Hand[0], strings.Repeat("{ Hidden }", len(p.Hand)-1))
		}
	}
}

func checkBust(g *Game) bool {
	for _, p := range g.Players {
		if Bust(&p) {
			ShowResults(g)
			fmt.Println("Player " + p.Name + " Busts!")
			return true
		}
	}
	return false
}

// ShowResults displays results
func ShowResults(g *Game) {
	print("\033[H\033[2J")
	for _, p := range g.Players {
		fmt.Println(p.Name+":", p.Hand)
	}
}

// HandValue checks total value of player's hand
func HandValue(p *Player) int {
	total := 0
	for _, card := range p.Hand {
		total += card.Value
	}
	return total
}

// Bust checks if the player's card values are over 21
func Bust(p *Player) bool {
	total := HandValue(p)
	if total > 21 {
		return true
	}
	return false
}
