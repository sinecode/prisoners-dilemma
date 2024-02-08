package strategies

import (
	"github.com/sinecode/prisoners-dilemma/src"
)

type AlwaysDefect struct {
	// Nothing. It always defects
}

func (p *AlwaysDefect) Name() string {
	return "AIID"
}

func (p *AlwaysDefect) Play(in chan src.Move, out chan src.Move, turns int) {
	for i := 0; i < turns; i++ {
		out <- src.Defect
		<-in
	}
}
