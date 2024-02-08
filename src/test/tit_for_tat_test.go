package test

import (
	"testing"

	"github.com/sinecode/prisoners-dilemma/src/strategies"
)

func TestTitForTatStrategy(t *testing.T) {
	testPlayer(t, &strategies.TitForTat{})
}
