package src

type Player interface {
	// in is the channel from which the player take the opponent's move
	// out is the channel to which the player send its move
	Play(in chan Signal, out chan Signal, turns int)
}
