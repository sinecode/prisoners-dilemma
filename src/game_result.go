package src

type GameResult struct {
	Strategy1Scores     []int
	Strategy1TotalScore int
	Strategy2Scores     []int
	Strategy2TotalScore int
}

func NewGameResult() *GameResult {
	return &GameResult{
		Strategy1Scores: make([]int, 0),
		Strategy2Scores: make([]int, 0),
	}
}

func (gr *GameResult) AddStrategy1Score(score int) {
	gr.Strategy1Scores = append(gr.Strategy1Scores, score)
	gr.Strategy1TotalScore += score
}

func (gr *GameResult) AddStrategy2Score(score int) {
	gr.Strategy2Scores = append(gr.Strategy2Scores, score)
	gr.Strategy2TotalScore += score
}
