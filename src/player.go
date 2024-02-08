package src

type Player interface {
	// Name of the player.
	Name() string

	// A description of the strategy.
	Strategy() string

	// in - the channel from which the player take the opponent's move.
	// out - the channel to which the player send its move
	// turns - the number of turns that the game will last
	Play(in chan Move, out chan Move, turns int)
}
