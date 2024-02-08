package strategies

import (
	"github.com/sinecode/prisoners-dilemma/src"
)

// Cooperates, until the opponent defects, and thereafter always defects.
type GrimTrigger struct {
	hasOpponentDefected bool
}

func (p *GrimTrigger) Name() string {
	return "GRIM"
}

func (p *GrimTrigger) Play(in chan src.Move, out chan src.Move, turns int) {
	for i := 0; i < turns; i++ {
		if p.hasOpponentDefected {
			out <- src.Defect
			<-in
		} else {
			out <- src.Cooperate
			p.hasOpponentDefected = src.Defect == <-in
		}
	}
}
