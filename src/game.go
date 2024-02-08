package src

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
		strategy1In <- strategy2Move
		strategy2In <- strategy1Move
		if strategy1Move == Cooperate && strategy2Move == Cooperate {
			result.AddStrategy1Score(3)
			result.AddStrategy2Score(3)
		} else if strategy1Move == Cooperate && strategy2Move == Defect {
			result.AddStrategy1Score(0)
			result.AddStrategy2Score(5)
		} else if strategy1Move == Defect && strategy2Move == Cooperate {
			result.AddStrategy1Score(5)
			result.AddStrategy2Score(0)
		} else if strategy1Move == Defect && strategy2Move == Defect {
			result.AddStrategy1Score(1)
			result.AddStrategy2Score(1)
		} else {
			panic("Invalid move")
		}
	}
	return result
}
