package src

import (
	"fmt"
)

func StartGame(player1 Player, player2 Player, turns int) {
	fmt.Println("Game is started!")

	player1Scores := make([]int, turns)
	player1In := make(chan Signal)
	player1Out := make(chan Signal)
	go player1.Play(player1In, player1Out, turns)

	player2Scores := make([]int, turns)
	player2In := make(chan Signal)
	player2Out := make(chan Signal)
	go player2.Play(player2In, player2Out, turns)

	for turn := 0; turn < turns; turn++ {
		fmt.Println("Starting turn", turn)
		player1Move := <-player1Out
		player2Move := <-player2Out
		player1In <- player2Move
		player2In <- player1Move
		if player1Move == Cooperate && player2Move == Cooperate {
			player1Scores[turn] = 3
			player2Scores[turn] = 3
		} else if player1Move == Cooperate && player2Move == Defect {
			player1Scores[turn] = 0
			player2Scores[turn] = 5
		} else if player1Move == Defect && player2Move == Cooperate {
			player1Scores[turn] = 5
			player2Scores[turn] = 0
		} else if player1Move == Defect && player2Move == Defect {
			player1Scores[turn] = 1
			player2Scores[turn] = 1
		} else {
			panic("Invalid move")
		}
		fmt.Println("")
	}
	fmt.Println("Player 1 scores:", player1Scores)
	fmt.Println("Player 2 scores:", player2Scores)
	fmt.Println("")
	player1Total := 0
	for i := 0; i < len(player1Scores); i++ {
		player1Total += player1Scores[i]
	}
	player2Total := 0
	for i := 0; i < len(player2Scores); i++ {
		player2Total += player2Scores[i]
	}
	fmt.Println("Player 1 total score:", player1Total)
	fmt.Println("Player 2 total score:", player2Total)
	fmt.Println("")
	if player1Total > player2Total {
		fmt.Println("PLAYER 1 WIN!!")
	} else if player1Total < player2Total {
		fmt.Println("PLAYER 2 WIN!!")
	} else {
		fmt.Println("IT'S A DRAW")
	}
}
