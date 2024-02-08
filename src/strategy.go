package src

type Strategy interface {
	// Name of the strategy.
	Name() string

	// A description of the strategy.
	Description() string

	// in - the channel from which the strategy take the opponent's move.
	// out - the channel to which the strategy send its move
	// turns - the number of turns that the game will last
	Play(in chan Move, out chan Move, turns int)
}
