package test

import (
	"testing"

	"github.com/sinecode/prisoners-dilemma/src/strategies"
)

func TestRandomStrategy(t *testing.T) {
	testPlayer(t, &strategies.Random{})
}
