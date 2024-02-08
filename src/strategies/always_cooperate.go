package strategies

import (
	"github.com/sinecode/prisoners-dilemma/src"
)

type AlwaysCooperate struct{}

func (p *AlwaysCooperate) Name() string {
	return "AIIC"
}

func (p *AlwaysCooperate) Strategy() string {
	return "Cooperates on every move."
}

func (p *AlwaysCooperate) Play(in chan src.Move, out chan src.Move, turns int) {
	for i := 0; i < turns; i++ {
		out <- src.Cooperate
		<-in
	}
}
