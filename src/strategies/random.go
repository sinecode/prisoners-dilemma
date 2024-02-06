package strategies

import (
	"fmt"
	"math/rand"

	"github.com/sinecode/prisoners-dilemma/src"
)

type Random struct {
	// Nothing. It really plays random
}

func (p *Random) Play(in chan src.Signal, out chan src.Signal, turns int) {
	for i := 0; i < turns; i++ {
		if rand.Intn(2) == 0 {
			out <- src.Defect
			fmt.Println("Random player plays Defect")
		} else {
			out <- src.Cooperate
			fmt.Println("Random player plays Cooperate")
		}
		// Don't do nothing. He plays random
		<-in
	}
}
