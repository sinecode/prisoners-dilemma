package strategies

import (
	"math/rand"

	"github.com/sinecode/prisoners-dilemma/src"
)

type Random struct{}

func (p *Random) Name() string {
	return "RAND"
}

func (p *Random) Strategy() string {
	return "Makes a random move."
}

func (p *Random) Play(in chan src.Move, out chan src.Move, turns int) {
	for i := 0; i < turns; i++ {
		if rand.Intn(2) == 0 {
			out <- src.Defect
		} else {
			out <- src.Cooperate
		}
		<-in
	}
}
