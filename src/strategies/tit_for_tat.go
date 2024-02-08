package strategies

import (
	"github.com/sinecode/prisoners-dilemma/src"
)

type TitForTat struct {
	lastOpponentMove src.Move
}

func (p *TitForTat) Name() string {
	return "TFT"
}

func (p *TitForTat) Description() string {
	return "Cooperates on the first move, then copies the opponent’s last move."
}

func (p *TitForTat) Play(in chan src.Move, out chan src.Move, turns int) {
	for i := 0; i < turns; i++ {
		if i == 0 {
			out <- src.Cooperate
		} else {
			out <- p.lastOpponentMove
		}
		p.lastOpponentMove = <-in
	}
}
