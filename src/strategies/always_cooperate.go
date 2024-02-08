package strategies

import (
	"github.com/sinecode/prisoners-dilemma/src"
)

type AlwaysCooperate struct {
	// Nothing. It always cooperates
}

func (p *AlwaysCooperate) Name() string {
	return "AIIC"
}

func (p *AlwaysCooperate) Play(in chan src.Signal, out chan src.Signal, turns int) {
	for i := 0; i < turns; i++ {
		out <- src.Cooperate
		<-in
	}
}
