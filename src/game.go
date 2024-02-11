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
		result.AddStrategy1Move(strategy1Move)
		result.AddStrategy2Move(strategy2Move)
		strategy1In <- strategy2Move
		strategy2In <- strategy1Move
		strategy1Score, strategy2Score := ComputeScores(strategy1Move, strategy2Move)
		result.AddStrategy1Score(strategy1Score)
		result.AddStrategy2Score(strategy2Score)
	}
	return result
}
