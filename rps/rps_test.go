package main

import (
	"strconv"
	"testing"
)

func TestGenerateMove(t *testing.T) {
	possibleOutcomes := []string{"Rock", "Paper", "Scissors"}

	if move := generateMove(); !includes(possibleOutcomes, move) {
		t.Errorf("Expected 'Paper', 'Rock', or 'Scissors' but got %v instead", move)
	}
}

func TestVictoryConditions(t *testing.T) {
	playerMove := generateMove()
	compMove := generateMove()
	start := Outcome{}
	end := start
	victoryConditions(playerMove, compMove, &start)
	if start == end {
		t.Errorf("Expected Outcome{} to change. Start: %v, End: %v", start, end)
	}
}

func TestPlayerWins(t *testing.T) {
	playerMove := "Paper"
	compMove := "Rock"
	outcome := Outcome{}
	victoryConditions(playerMove, compMove, &outcome)
	if outcome.playerWins != 1 {
		t.Errorf("Expected player to have '1' win. Instead, player has '%v' wins.", outcome.playerWins)
	}
	if outcome.compWins != 0 {
		t.Errorf("Expected comp to have '0' wins. Instead, comp has '%v' win.", outcome.compWins)
	}
	if outcome.ties != 0 {
		t.Errorf("Expected ties to equal '0'. Instead, ties equal '%v'.", outcome.ties)
	}
}

func TestCompWins(t *testing.T) {
	playerMove := "Paper"
	compMove := "Scissors"
	outcome := Outcome{}
	victoryConditions(playerMove, compMove, &outcome)
	if outcome.playerWins != 0 {
		t.Errorf("Expected player to have '0' wins. Instead, player has '%v' win.", outcome.playerWins)
	}
	if outcome.compWins != 1 {
		t.Errorf("Expected comp to have '1' win. Instead, comp has '%v' wins.", outcome.compWins)
	}
	if outcome.ties != 0 {
		t.Errorf("Expected ties to equal '0'. Instead, ties equal '%v'.", outcome.ties)
	}
}

func TestTies(t *testing.T) {
	playerMove := "Paper"
	compMove := "Paper"
	outcome := Outcome{}
	victoryConditions(playerMove, compMove, &outcome)
	if outcome.playerWins != 0 {
		t.Errorf("Expected player to have '0' wins. Instead, player has '%v' win.", outcome.playerWins)
	}
	if outcome.compWins != 0 {
		t.Errorf("Expected comp to have 0 wins. Instead, comp has '%v' wins.", outcome.compWins)
	}
	if outcome.ties != 1 {
		t.Errorf("Expected ties to equal '1'. Instead, ties equal '%v'.", outcome.ties)
	}
}

func TestPrintOutcome(t *testing.T) {
	playerMove := "Paper"
	compMove := "Paper"
	outcome := Outcome{}

	victoryConditions(playerMove, compMove, &outcome)

	playerWins := "Player-" + strconv.Itoa(outcome.playerWins)
	compWins := " | Comp-" + strconv.Itoa(outcome.compWins)
	ties := " | Ties-" + strconv.Itoa(outcome.ties)

	output := playerWins + compWins + ties
	if output != "Player-0 | Comp-0 | Ties-1" {
		t.Errorf("Result: %v\n Expected: Player-0 | Comp-0 | Ties-1", output)
	}
}

func includes(slice []string, item string) bool {
	for _, a := range slice {
		if a == item {
			return true
		}
	}
	return false
}
