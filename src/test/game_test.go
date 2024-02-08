package test

import (
	"testing"

	"github.com/sinecode/prisoners-dilemma/src"
	"github.com/sinecode/prisoners-dilemma/src/strategies"
)

func TestGame(t *testing.T) {
	game := src.NewGame(&strategies.AlwaysDefect{}, &strategies.AlwaysCooperate{}, 10)
	result := game.Start()
	if result.Player1TotalScore < result.Player2TotalScore {
		t.Errorf("Wrong total scores. Player1 got %d, player 2 got %d", result.Player1TotalScore, result.Player2TotalScore)
	}
}
