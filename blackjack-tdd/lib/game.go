package lib

import "fmt"

// PlayGame just prints stuff; it doesn't play the game
func PlayGame(p []Player, d *Deck) {
	fmt.Println(p[0])
	fmt.Println("")
	fmt.Println(d)
}
