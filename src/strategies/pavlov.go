package strategies

import (
	"github.com/sinecode/prisoners-dilemma/src"
)

type Pavlov struct {
	nextMove src.Move
}

func (p *Pavlov) Name() string {
	return "Pavlov"
}

func (p *Pavlov) Description() string {
	return "Cooperates on the first move. If a reward or temptation payoff is received in the last round then repeats last choice, otherwise chooses the opposite choice."
}

func (p *Pavlov) Play(in chan src.Move, out chan src.Move, turns int) {
	for i := 0; i < turns; i++ {
		var myMove src.Move
		if i == 0 {
			myMove = src.Cooperate
		} else {
			myMove = p.nextMove
		}
		out <- myMove
		opponentMove := <-in
		myScore, _ := src.ComputeScores(myMove, opponentMove)
		if myScore == src.REWARD_PAYOFF || myScore == src.TEMPTATION_PAYOFF {
			p.nextMove = myMove
		} else {
			if myMove == src.Defect {
				p.nextMove = src.Cooperate
			} else {
				p.nextMove = src.Defect
			}
		}
	}
}
