package test

import (
	"testing"

	"github.com/sinecode/prisoners-dilemma/src"
)

func testStrategy(t *testing.T, strategy src.Strategy) {
	numTurns := []int{1, 10, 100, 1000, 5000, 10000}
	for _, turns := range numTurns {
		testGame(t, strategy, turns)
	}
}

func testGame(t *testing.T, strategy src.Strategy, turns int) {
	in := make(chan src.Move)
	out := make(chan src.Move)

	go strategy.Play(in, out, turns)

	for turn := 0; turn < turns; turn++ {
		strategyOut := <-out
		var dummy src.Move
		if turn%2 == 0 {
			dummy = src.Defect
		} else {
			dummy = src.Cooperate
		}
		in <- dummy
		if strategyOut != src.Cooperate && strategyOut != src.Defect {
			t.Errorf("Expected a %d or %d as output, got %d", src.Defect, src.Cooperate, strategyOut)
		}
	}
}
