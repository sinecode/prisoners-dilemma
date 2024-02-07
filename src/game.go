package src

type Game struct {
	Player1 Player
	Player2 Player
	Turns   int
}

func NewGame(player1 Player, player2 Player, turns int) *Game {
	return &Game{
		Player1: player1,
		Player2: player2,
		Turns:   turns,
	}
}

func (game *Game) Start() *GameResult {
	result := NewGameResult()
	player1In := make(chan Signal)
	player1Out := make(chan Signal)
	go game.Player1.Play(player1In, player1Out, game.Turns)

	player2In := make(chan Signal)
	player2Out := make(chan Signal)
	go game.Player2.Play(player2In, player2Out, game.Turns)

	for turn := 0; turn < game.Turns; turn++ {
		player1Move := <-player1Out
		player2Move := <-player2Out
		player1In <- player2Move
		player2In <- player1Move
		if player1Move == Cooperate && player2Move == Cooperate {
			result.AddPlayer1Score(3)
			result.AddPlayer2Score(3)
		} else if player1Move == Cooperate && player2Move == Defect {
			result.AddPlayer1Score(0)
			result.AddPlayer2Score(5)
		} else if player1Move == Defect && player2Move == Cooperate {
			result.AddPlayer1Score(5)
			result.AddPlayer2Score(0)
		} else if player1Move == Defect && player2Move == Defect {
			result.AddPlayer1Score(1)
			result.AddPlayer2Score(1)
		} else {
			panic("Invalid move")
		}
	}
	return result
}
