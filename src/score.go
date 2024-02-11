package src

const (
	REWARD_PAYOFF     = 3
	TEMPTATION_PAYOFF = 5
	SUCKERS_PAYOFF    = 0
	PUNISHMENT_PAYOFF = 1
)

func ComputeScores(move1 Move, move2 Move) (int, int) {
	if move1 == Cooperate && move2 == Cooperate {
		return REWARD_PAYOFF, REWARD_PAYOFF
	} else if move1 == Cooperate && move2 == Defect {
		return SUCKERS_PAYOFF, TEMPTATION_PAYOFF
	} else if move1 == Defect && move2 == Cooperate {
		return TEMPTATION_PAYOFF, SUCKERS_PAYOFF
	} else if move1 == Defect && move2 == Defect {
		return PUNISHMENT_PAYOFF, PUNISHMENT_PAYOFF
	} else {
		panic("Invalid move")
	}
}
