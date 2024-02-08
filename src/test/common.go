package test

import (
	"testing"

	"github.com/sinecode/prisoners-dilemma/src"
)

func testPlayer(t *testing.T, player src.Player) {
	numTurns := []int{1, 10, 100, 1000, 5000, 10000}
	for _, turns := range numTurns {
		testGame(t, player, turns)
	}
}

func testGame(t *testing.T, player src.Player, turns int) {
	in := make(chan src.Move)
	out := make(chan src.Move)

	go player.Play(in, out, turns)

	for turn := 0; turn < turns; turn++ {
		playerOut := <-out
		var dummy src.Move
		if turn%2 == 0 {
			dummy = src.Defect
		} else {
			dummy = src.Cooperate
		}
		in <- dummy
		if playerOut != src.Cooperate && playerOut != src.Defect {
			t.Errorf("Expected a %d or %d as output, got %d", src.Defect, src.Cooperate, playerOut)
		}
	}
}
