package main

import "github.com/mattrice12/go-practice/blackjack/lib"

func main() {
	deck := lib.PrepDeck()
	game := lib.PrepPlayers(&deck)

	lib.PlayRound(&deck, &game)
	// lib.ShowResults(&game)
}
