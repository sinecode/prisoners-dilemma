package src

type GameResult struct {
	strategy1Moves      []Move
	Strategy1Scores     []int
	Strategy1TotalScore int

	strategy2Moves      []Move
	Strategy2Scores     []int
	Strategy2TotalScore int
}

func NewGameResult() *GameResult {
	return &GameResult{
		strategy1Moves:  make([]Move, 0),
		Strategy1Scores: make([]int, 0),
		strategy2Moves:  make([]Move, 0),
		Strategy2Scores: make([]int, 0),
	}
}

func (gr *GameResult) AddStrategy1Move(move Move) {
	gr.strategy1Moves = append(gr.strategy1Moves, move)
}

func (gr *GameResult) AddStrategy1Score(score int) {
	gr.Strategy1Scores = append(gr.Strategy1Scores, score)
	gr.Strategy1TotalScore += score
}

func (gr *GameResult) AddStrategy2Move(move Move) {
	gr.strategy2Moves = append(gr.strategy2Moves, move)
}

func (gr *GameResult) AddStrategy2Score(score int) {
	gr.Strategy2Scores = append(gr.Strategy2Scores, score)
	gr.Strategy2TotalScore += score
}

func (gr *GameResult) ReadableStrategy1Moves() []string {
	return readableMoves(gr.strategy1Moves)
}

func (gr *GameResult) ReadableStrategy2Moves() []string {
	return readableMoves(gr.strategy2Moves)
}

func readableMoves(strategyMoves []Move) []string {
	moves := make([]string, len(strategyMoves))
	for i, move := range strategyMoves {
		var readableMove string
		if move == Cooperate {
			readableMove = "C"
		} else if move == Defect {
			readableMove = "D"
		} else {
			panic("Unknown move")
		}
		moves[i] = readableMove
	}
	return moves
}
