package strategies

import (
	"github.com/sinecode/prisoners-dilemma/src"
)

type FirmButFair struct {
	shouldDefetct bool
}

func (p *FirmButFair) Name() string {
	return "FBF"
}

func (p *FirmButFair) Description() string {
	return "Cooperates on the first move, and cooperates except after receiving a sucker payoff."
}

func (p *FirmButFair) Play(in chan src.Move, out chan src.Move, turns int) {
	p.shouldDefetct = false
	for i := 0; i < turns; i++ {
		if p.shouldDefetct {
			out <- src.Defect
			<-in
			continue
		}

		myMove := src.Cooperate
		out <- myMove
		opponentMove := <-in
		myScore, _ := src.ComputeScores(myMove, opponentMove)
		if myScore == src.SUCKERS_PAYOFF {
			p.shouldDefetct = true
		}
	}
}
