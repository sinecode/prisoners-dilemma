package test

import (
	"testing"

	"github.com/sinecode/prisoners-dilemma/src"
	"github.com/sinecode/prisoners-dilemma/src/strategies"
)

func TestGame(t *testing.T) {
	game := src.NewGame(&strategies.AlwaysDefect{}, &strategies.AlwaysCooperate{}, 10)
	result := game.Start()
	if result.Strategy1TotalScore < result.Strategy2TotalScore {
		t.Errorf("Wrong total scores. Strategy1 got %d, strategy 2 got %d", result.Strategy1TotalScore, result.Strategy2TotalScore)
	}
}
