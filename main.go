package main

import (
	"fmt"

	"github.com/sinecode/prisoners-dilemma/src"
	"github.com/sinecode/prisoners-dilemma/src/strategies"
)

func main() {
	player1 := &strategies.Random{}
	player2 := &strategies.Random{}
	game := src.NewGame(player1, player2, 10)
	result := game.Start()
	fmt.Println("")
	fmt.Println("Player 1 scores:", result.Player1Scores)
	fmt.Println("Player 2 scores:", result.Player2Scores)
	fmt.Println("")
	fmt.Println("")
	fmt.Println("Player 1 total score:", result.Player1TotalScore)
	fmt.Println("Player 2 total score:", result.Player2TotalScore)
	fmt.Println("")
	if result.Player1TotalScore > result.Player2TotalScore {
		fmt.Println("PLAYER 1 WIN!!")
	} else if result.Player1TotalScore < result.Player2TotalScore {
		fmt.Println("PLAYER 2 WIN!!")
	} else {
		fmt.Println("IT'S A DRAW")
	}
}
