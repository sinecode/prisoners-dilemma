package main

import (
	"fmt"

	"github.com/sinecode/prisoners-dilemma/src"
	"github.com/sinecode/prisoners-dilemma/src/strategies"
)

func main() {
	player1 := &strategies.Random{}
	player2 := &strategies.AlwaysCooperate{}
	game := src.NewGame(player1, player2, 10)
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
