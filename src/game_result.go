package src

type GameResult struct {
	Player1Scores     []int
	Player1TotalScore int
	Player2Scores     []int
	Player2TotalScore int
}

func NewGameResult() *GameResult {
	return &GameResult{
		Player1Scores: make([]int, 1),
		Player2Scores: make([]int, 1),
	}
}

func (gr *GameResult) AddPlayer1Score(score int) {
	gr.Player1Scores = append(gr.Player1Scores, score)
	gr.Player1TotalScore += score
}

func (gr *GameResult) AddPlayer2Score(score int) {
	gr.Player2Scores = append(gr.Player2Scores, score)
	gr.Player2TotalScore += score
}
