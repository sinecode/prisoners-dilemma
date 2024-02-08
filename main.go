package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/sinecode/prisoners-dilemma/src"
	"github.com/sinecode/prisoners-dilemma/src/strategies"
)

var allStrategies = []src.Strategy{
	&strategies.AlwaysDefect{},
	&strategies.AlwaysCooperate{},
	&strategies.Random{},
	&strategies.TitForTat{},
	&strategies.GrimTrigger{},
}

func printUsage() {
	fmt.Println("Usage: prisoners game <turns> <strategy1> <strategy2>")
	fmt.Println("")
	fmt.Println("Choose from these strategies (case insensitive):")
	for _, strategy := range allStrategies {
		fmt.Printf("  * %-4s - %s\n", strategy.Name(), strategy.Description())
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

	strategy1Name := strings.ToLower(os.Args[argsIdx])
	var strategy1 src.Strategy
	for _, strategy := range allStrategies {
		if strategy1Name == strings.ToLower(strategy.Name()) {
			strategy1 = strategy
			break
		}
	}
	if strategy1 == nil {
		fmt.Println("Unknown <strategy1>:", strategy1Name)
		return
	}
	argsIdx++

	strategy2Name := strings.ToLower(os.Args[argsIdx])
	var strategy2 src.Strategy
	for _, strategy := range allStrategies {
		if strategy2Name == strings.ToLower(strategy.Name()) {
			strategy2 = strategy
			break
		}
	}
	if strategy2 == nil {
		fmt.Println("Unknown <strategy2>:", strategy2Name)
		return
	}

	game := src.NewGame(strategy1, strategy2, turns)
	result := game.Start()
	fmt.Printf("%-4s scores: %d\n", strategy1.Name(), result.Strategy1Scores)
	fmt.Printf("%-4s scores: %d\n", strategy2.Name(), result.Strategy2Scores)
	fmt.Println("")
	fmt.Printf("%-4s total score: %d\n", strategy1.Name(), result.Strategy1TotalScore)
	fmt.Printf("%-4s total score: %d\n", strategy2.Name(), result.Strategy2TotalScore)
	fmt.Println("")
	if result.Strategy1TotalScore > result.Strategy2TotalScore {
		fmt.Printf("%s WIN!!\n", strategy1.Name())
	} else if result.Strategy1TotalScore < result.Strategy2TotalScore {
		fmt.Printf("%s WIN!!\n", strategy2.Name())
	} else {
		fmt.Println("IT'S A DRAW")
	}
}
