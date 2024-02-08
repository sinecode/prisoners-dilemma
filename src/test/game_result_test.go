package test

import (
	"testing"

	"github.com/sinecode/prisoners-dilemma/src"
)

func TestGameResult(t *testing.T) {
	result := src.NewGameResult()
	result.AddStrategy1Score(3)
	result.AddStrategy2Score(1)
	result.AddStrategy1Score(7)
	result.AddStrategy2Score(5)
	got := result.Strategy1TotalScore
	want := 10
	if got != want {
		t.Errorf("Unexpected strategy 1 total score. Got %d, want %d", got, want)
	}
	got = result.Strategy2TotalScore
	want = 6
	if got != want {
		t.Errorf("Unexpected strategy 2 total score. Got %d, want %d", got, want)
	}
}
