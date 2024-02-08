package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/sinecode/prisoners-dilemma/src"
	"github.com/sinecode/prisoners-dilemma/src/strategies"
)

var allStrategies = []src.Player{
	&strategies.AlwaysDefect{},
	&strategies.AlwaysCooperate{},
	&strategies.Random{},
	&strategies.TitForTat{},
	&strategies.GrimTrigger{},
}

func printUsage() {
	fmt.Println("Usage: prisoners game <turns> <player1> <player2>")
	fmt.Println("")
	fmt.Println("Choose from these players (case insensitive):")
	for _, strategy := range allStrategies {
		fmt.Println("  *", strategy.Name(), "-", strategy.Strategy())
	}
}

func main() {
	argsIdx := 1
	if len(os.Args) != 5 || os.Args[argsIdx] == "-h" || os.Args[argsIdx] == "--help" {
		printUsage()
		return
	}
	argsIdx++

	turnsArg := os.Args[argsIdx]
	turns, err := strconv.Atoi(turnsArg)
	if err != nil || turns <= 0 {
		fmt.Println("<turns> must be a positive integer")
		return
	}
	argsIdx++

	player1Name := strings.ToLower(os.Args[argsIdx])
	var player1 src.Player
	for _, strategy := range allStrategies {
		if player1Name == strings.ToLower(strategy.Name()) {
			player1 = strategy
			break
		}
	}
	if player1 == nil {
		fmt.Println("Unknown <player1>:", player1Name)
		return
	}
	argsIdx++

	player2Name := strings.ToLower(os.Args[argsIdx])
	var player2 src.Player
	for _, strategy := range allStrategies {
		if player2Name == strings.ToLower(strategy.Name()) {
			player2 = strategy
			break
		}
	}
	if player2 == nil {
		fmt.Println("Unknown <player2>:", player2Name)
		return
	}

	game := src.NewGame(player1, player2, turns)
	result := game.Start()
	fmt.Printf("%s scores: %d\n", player1.Name(), result.Player1Scores)
	fmt.Printf("%s total score: %d\n", player1.Name(), result.Player1TotalScore)
	fmt.Println("")
	fmt.Printf("%s scores: %d\n", player2.Name(), result.Player2Scores)
	fmt.Printf("%s total score: %d\n", player2.Name(), result.Player2TotalScore)
	fmt.Println("")
	if result.Player1TotalScore > result.Player2TotalScore {
		fmt.Printf("%s WIN!!\n", player1.Name())
	} else if result.Player1TotalScore < result.Player2TotalScore {
		fmt.Printf("%s WIN!!\n", player2.Name())
	} else {
		fmt.Println("IT'S A DRAW")
	}
}
