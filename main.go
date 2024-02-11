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
	&strategies.Pavlov{},
	&strategies.FirmButFair{},
}

func main() {
	if len(os.Args) < 2 || os.Args[1] == "-h" || os.Args[1] == "--help" {
		printUsage()
		return
	}

	if os.Args[1] == "game" {
		playGame()
	} else if os.Args[1] == "tournament" {
		playtournament()
	} else {
		printUsage()
	}
}

func printUsage() {
	fmt.Println("Usage:")
	fmt.Println("  * prisoners game <turns> <strategy1> <strategy2>")
	fmt.Println("  * prisoners tournament <turns>")
	fmt.Println("")
	fmt.Println("For a game, choose from these strategies (case insensitive):")
	for _, strategy := range allStrategies {
		fmt.Printf("  * %-6s - %s\n", strategy.Name(), strategy.Description())
	}
}

func playGame() {
	if len(os.Args) != 5 {
		printUsage()
		return
	}

	argsIdx := 2

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
	fmt.Printf("%-6s moves:  %s\n", strategy1.Name(), result.ReadableStrategy1Moves())
	fmt.Printf("%-6s moves:  %s\n", strategy2.Name(), result.ReadableStrategy2Moves())
	fmt.Println("")
	fmt.Printf("%-6s scores: %d\n", strategy1.Name(), result.Strategy1Scores)
	fmt.Printf("%-6s scores: %d\n", strategy2.Name(), result.Strategy2Scores)
	fmt.Println("")
	fmt.Printf("%-6s total score: %d\n", strategy1.Name(), result.Strategy1TotalScore)
	fmt.Printf("%-6s total score: %d\n", strategy2.Name(), result.Strategy2TotalScore)
	fmt.Println("")
	if result.Strategy1TotalScore > result.Strategy2TotalScore {
		fmt.Printf("%s WIN!!\n", strategy1.Name())
	} else if result.Strategy1TotalScore < result.Strategy2TotalScore {
		fmt.Printf("%s WIN!!\n", strategy2.Name())
	} else {
		fmt.Println("IT'S A DRAW")
	}
}

func playtournament() {
	if len(os.Args) != 3 {
		printUsage()
		return
	}

	turnsArg := os.Args[2]
	turns, err := strconv.Atoi(turnsArg)
	if err != nil || turns <= 0 {
		fmt.Println("<turns> must be a positive integer")
		return
	}

	tournament := *src.NewTournament(allStrategies, turns)
	result := tournament.Start()
	sortedStrategies, sortedScores := result.ReadableRanking()
	fmt.Println("STRATEGY  SCORE")
	fmt.Println("")
	for i := 0; i < len(sortedStrategies); i++ {
		fmt.Printf("%-9s %d\n", sortedStrategies[i], sortedScores[i])
	}
}
