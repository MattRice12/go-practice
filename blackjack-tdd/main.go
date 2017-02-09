package main

import "github.com/mattrice12/go-practice/blackjack-tdd/lib"

func main() {
	deck := lib.PrepDeck()
	players := lib.PrepPlayers(&deck)
	lib.PlayGame(&players, &deck)
}
