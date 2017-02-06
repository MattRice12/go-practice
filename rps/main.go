package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

// Outcome tracks the wins/losses/ties in the game
type Outcome struct {
	playerWins int
	compWins   int
	ties       int
}

func main() {
	const playTimes int = 1000
	outcome := Outcome{}
	runGame(playTimes, &outcome)
}

func runGame(times int, outcome *Outcome) {
	for i := 1; i <= times; i++ {
		playerMove := generateMove()
		compMove := generateMove()
		victoryConditions(playerMove, compMove, outcome)
		printOutcome(playerMove, compMove, outcome)
	}
}

func printOutcome(playerMove string, compMove string, o *Outcome) {
	playerWins := "Player-" + strconv.Itoa(o.playerWins)
	compWins := " | Comp-" + strconv.Itoa(o.compWins)
	ties := " | Ties-" + strconv.Itoa(o.ties)
	print("\033[H\033[2J")
	fmt.Println("Player: ", playerMove)
	fmt.Println("Comp: ", compMove)
	fmt.Println("Result: ", playerWins+compWins+ties)
	time.Sleep(time.Millisecond)
}

func victoryConditions(p string, c string, o *Outcome) {
	choiceSlice := []string{"Rock", "Paper", "Scissors", "Rock"}
	for i := range choiceSlice[:len(choiceSlice)-1] {
		if p == choiceSlice[i] && c == choiceSlice[i+1] {
			o.compWins++
			return
		}
	}
	if p == c {
		o.ties++
	} else {
		o.playerWins++
	}
}

func random(min, max int) int {
	rand.Seed(int64(time.Now().Nanosecond()))
	return rand.Intn(max-min) + min
}

func generateMove() string {
	var move string
	choiceSlice := []string{"Rock", "Paper", "Scissors"}
	num := random(0, 3)
	move = choiceSlice[num]
	return move
}

//
// func generatePlayerMove() string {
// 	move := ""
// 	fmt.Print("Rock, Paper, or Scissors?: ")
// 	fmt.Scanln(&move)
// 	move = strings.Title(strings.ToLower(move))
// 	move = validatePlayerMove(move)
// 	return move
// }
//
// func includes(slice []string, item string) bool {
// 	for _, a := range slice {
// 		if a == item {
// 			return true
// 		}
// 	}
// 	return false
// }
//
//
// func validatePlayerMove(move string) string {
// 	choiceSlice := []string{"Rock", "Paper", "Scissors"}
// 	if includes(choiceSlice, move) == false {
// 		fmt.Print("Invalid Input...")
// 		return generatePlayerMove()
// 	}
// 	return move
// }
