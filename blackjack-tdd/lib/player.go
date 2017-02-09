package lib

import "strconv"

// Player is a player struct
type Player struct {
	Name  string
	Hand  []Card
	Total int
	Stay  bool
}

// PrepPlayers prepares a player
func PrepPlayers(d *Deck) []Player {
	players := []Player{}
	for i := 0; i < 4; i++ {
		player := CreatePlayer(d, i)
		players = append(players, player)
	}
	return players
}

// Draw comment
func (p *Player) Draw(d *Deck) {
	var card Card
	card, d.Cards = d.Cards[len(d.Cards)-1], d.Cards[:len(d.Cards)-1]
	p.Hand = append(p.Hand, card)
}

// CreatePlayer creates players
func CreatePlayer(d *Deck, i int) Player {
	name := "Player" + strconv.Itoa(i)
	player := Player{name, []Card{}, 0, false}
	player.CreateHand(d)
	return player
}

// CreateHand comment
func (p *Player) CreateHand(d *Deck) {
	for i := 0; i < 2; i++ {
		p.Draw(d)
	}
}
