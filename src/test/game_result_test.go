package test

import (
	"testing"

	"github.com/sinecode/prisoners-dilemma/src"
)

func TestGameResult(t *testing.T) {
	result := src.NewGameResult()
	result.AddPlayer1Score(3)
	result.AddPlayer2Score(1)
	result.AddPlayer1Score(7)
	result.AddPlayer2Score(5)
	got := result.Player1TotalScore
	want := 10
	if got != want {
		t.Errorf("Unexpected player 1 total score. Got %d, want %d", got, want)
	}
	got = result.Player2TotalScore
	want = 6
	if got != want {
		t.Errorf("Unexpected player 2 total score. Got %d, want %d", got, want)
	}
}
