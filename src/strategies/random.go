package strategies

import (
	"math/rand"

	"github.com/sinecode/prisoners-dilemma/src"
)

type Random struct {
	// Nothing. It really plays random
}

func (p *Random) Name() string {
	return "RAND"
}

func (p *Random) Play(in chan src.Signal, out chan src.Signal, turns int) {
	for i := 0; i < turns; i++ {
		if rand.Intn(2) == 0 {
			out <- src.Defect
		} else {
			out <- src.Cooperate
		}
		// Don't do nothing. He plays random
		<-in
	}
}
