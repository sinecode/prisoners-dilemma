package test

import (
	"testing"

	"github.com/sinecode/prisoners-dilemma/src"
)

type dummyPlayer struct {
	PlayFunc func(in chan src.Signal, out chan src.Signal, turns int)
}

func (fp dummyPlayer) Play(in chan src.Signal, out chan src.Signal, turns int) {
	fp.PlayFunc(in, out, turns)
}

func NewDummyPlayer(playFunc func(in chan src.Signal, out chan src.Signal, turns int)) dummyPlayer {
	return dummyPlayer{playFunc}
}

func TestGame(t *testing.T) {
	alwaysDefectPlayer := NewDummyPlayer(
		func(in chan src.Signal, out chan src.Signal, turns int) {
			for turn := 0; turn < turns; turn++ {
				out <- src.Defect
				<-in
			}
		},
	)
	alwaysCooperatePlayer := NewDummyPlayer(
		func(in chan src.Signal, out chan src.Signal, turns int) {
			for turn := 0; turn < turns; turn++ {
				out <- src.Cooperate
				<-in
			}
		},
	)

	game := src.NewGame(alwaysDefectPlayer, alwaysCooperatePlayer, 10)
	result := game.Start()
	if result.Player1TotalScore < result.Player2TotalScore {
		t.Errorf("Wrong total scores. Player1 got %d, player 2 got %d", result.Player1TotalScore, result.Player2TotalScore)
	}
}
