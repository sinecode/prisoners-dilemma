package src

const (
	REWARD_PAYOFF     = 3
	TEMPTATION_PAYOFF = 5
	SUCKERS_PAYOFF    = 0
	PUNISHMENT_PAYOFF = 1
)

type Game struct {
	Strategy1 Strategy
	Strategy2 Strategy
	Turns     int
}

func NewGame(strategy1 Strategy, strategy2 Strategy, turns int) *Game {
	return &Game{
		Strategy1: strategy1,
		Strategy2: strategy2,
		Turns:     turns,
	}
}

func (game *Game) Start() *GameResult {
	result := NewGameResult()
	strategy1In := make(chan Move)
	strategy1Out := make(chan Move)
	go game.Strategy1.Play(strategy1In, strategy1Out, game.Turns)

	strategy2In := make(chan Move)
	strategy2Out := make(chan Move)
	go game.Strategy2.Play(strategy2In, strategy2Out, game.Turns)

	for turn := 0; turn < game.Turns; turn++ {
		strategy1Move := <-strategy1Out
		strategy2Move := <-strategy2Out
		result.AddStrategy1Move(strategy1Move)
		result.AddStrategy2Move(strategy2Move)
		strategy1In <- strategy2Move
		strategy2In <- strategy1Move
		// TODO: centralize score computation
		if strategy1Move == Cooperate && strategy2Move == Cooperate {
			result.AddStrategy1Score(REWARD_PAYOFF)
			result.AddStrategy2Score(REWARD_PAYOFF)
		} else if strategy1Move == Cooperate && strategy2Move == Defect {
			result.AddStrategy1Score(SUCKERS_PAYOFF)
			result.AddStrategy2Score(TEMPTATION_PAYOFF)
		} else if strategy1Move == Defect && strategy2Move == Cooperate {
			result.AddStrategy1Score(TEMPTATION_PAYOFF)
			result.AddStrategy2Score(SUCKERS_PAYOFF)
		} else if strategy1Move == Defect && strategy2Move == Defect {
			result.AddStrategy1Score(PUNISHMENT_PAYOFF)
			result.AddStrategy2Score(PUNISHMENT_PAYOFF)
		} else {
			panic("Invalid move")
		}
	}
	return result
}
