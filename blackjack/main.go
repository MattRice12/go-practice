package main

import (
	"fmt"

	"github.com/mattrice12/go-practice/blackjack/lib"
)

func main() {
	deck := lib.PrepDeck()
	player1 := lib.CreatePlayer("Matt", &deck)
	player2 := lib.CreatePlayer("Hal 9000", &deck)
	fmt.Println(player1.Name+": ", player1.Hand[0], player1.Hand[1])
	fmt.Println(player2.Name+": ", player2.Name, player2.Hand[0], "{ Hidden }")

	lib.HitOrStay(&player1, &deck)
}
