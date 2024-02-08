package test

import (
	"testing"

	"github.com/sinecode/prisoners-dilemma/src"
	"github.com/sinecode/prisoners-dilemma/src/strategies"
)

func TestStrategies(t *testing.T) {
	var tests = []struct {
		player src.Player
	}{
		{&strategies.Random{}},
		{&strategies.AlwaysCooperate{}},
		{&strategies.AlwaysDefect{}},
		{&strategies.TitForTat{}},
		{&strategies.GrimTrigger{}},
	}
	for _, test := range tests {
		testPlayer(t, test.player)
	}
}