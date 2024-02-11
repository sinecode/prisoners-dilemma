package src

type Tournament struct {
	Strategies []Strategy
	Turns      int
}

func NewTournament(strategies []Strategy, turns int) *Tournament {
	return &Tournament{
		Strategies: strategies,
		Turns:      turns,
	}
}

func (t *Tournament) Start() *TournamentResult {
	result := NewTournamentResult()
	for _, pairedStrategies := range generatePairs(t.Strategies) {
		game := NewGame(pairedStrategies[0], pairedStrategies[1], t.Turns)
		gameResult := game.Start()
		result.AddScore(pairedStrategies[0], gameResult.Strategy1TotalScore)
		result.AddScore(pairedStrategies[1], gameResult.Strategy1TotalScore)
	}
	return result
}

func generatePairs(arr []Strategy) [][]Strategy {
	var pairs [][]Strategy
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			pair := []Strategy{arr[i], arr[j]}
			pairs = append(pairs, pair)
		}
	}
	return pairs
}
