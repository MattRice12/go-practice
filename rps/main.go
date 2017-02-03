package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	playerMove := generatePlayerMove()
	compMove := generateCompMove()
	outcome := victoryConditions(playerMove, compMove)
	printOutcome(playerMove, compMove, outcome)
}

func printOutcome(playerMove, compMove, outcome string) {
	fmt.Println("Player: ", playerMove)
	fmt.Println("Comp: ", compMove)
	fmt.Println("Result: ", outcome)
}

func victoryConditions(p, c string) string {
	pWin := "Player Wins!"
	cWin := "Computer Wins!"
	if p == c {
		return "Tie"
	} else if p == "Paper" && c == "Rock" {
		return pWin
	} else if p == "Rock" && c == "Scissors" {
		return pWin
	} else if p == "Scissors" && c == "Paper" {
		return pWin
	}
	return cWin
}

func generatePlayerMove() string {
	move := ""
	fmt.Print("Rock, Paper, or Scissors?: ")
	fmt.Scanln(&move)
	move = strings.Title(strings.ToLower(move))
	move = validatePlayerMove(move)
	return move
}

func includes(slice []string, item string) bool {
	for _, a := range slice {
		if a == item {
			return true
		}
	}
	return false
}

func validatePlayerMove(move string) string {
	choiceSlice := []string{"Rock", "Paper", "Scissors"}
	if includes(choiceSlice, move) == false {
		fmt.Print("Invalid Input...")
		return generatePlayerMove()
	}
	return move
}

func random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}

func generateCompMove() string {
	var move string
	choiceSlice := []string{"Rock", "Paper", "Scissors"}
	num := random(0, 3)
	move = choiceSlice[num]
	return move
}
