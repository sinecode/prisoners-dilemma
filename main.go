package main

import (
	"github.com/sinecode/prisoners-dilemma/src"
	"github.com/sinecode/prisoners-dilemma/src/strategies"
)

func main() {
	player1 := &strategies.Random{}
	player2 := &strategies.Random{}
	src.StartGame(player1, player2, 100)
}
