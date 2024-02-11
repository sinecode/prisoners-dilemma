# Prisoner's dilemma

Go implementation of the famous [Prisoner's Dilemma](https://en.wikipedia.org/wiki/Prisoner's_dilemma).

## Usage

### Single game

To run a single game between two strategies:

```
go run main.go game <turns> <strategy1> <strategy2>
```

Where `<turns>` must be a positive integer. `<strategy1>` and `<strategy2`> must be chosen from (case insensitive):

* **AIID**   - Defects on every move.
* **AIIC**   - Cooperates on every move.
* **RAND**   - Makes a random move.
* **TFT**   - Cooperates on the first move, then copies the opponent’s last move.
* **GRIM**   - Cooperates, until the opponent defects, and thereafter always defects.
* **Pavlov** - Cooperates on the first move. If a reward or temptation payoff is received in the last round then repeats last choice, otherwise chooses the opposite choice.
* **FBF**   - Cooperates on the first move, and cooperates except after receiving a sucker payoff.

For example:

```go
$ go run main.go game 10 tft rand
TFT    moves:  [C C C D C C C D D C]
RAND   moves:  [C C D C C C D D C C]

TFT    scores: [3 3 0 5 3 3 0 1 5 3]
RAND   scores: [3 3 5 0 3 3 5 1 0 3]

TFT    total score: 26
RAND   total score: 26

IT'S A DRAW
```

### Tournament

You can also run a tournament. Which means that every strategy plays against each other strategy, including itself.

Run

```
go run main.go tournament <turns>
```

For example:

```
go run main.go tournament 200
STRATEGY  SCORE

Pavlov    3444
AIIC      3088
AIID      2808
TFT       2646
FBF       2325
RAND      2006
GRIM      1325
```

## Add a new strategy

To add a new strategy, you must follow these steps:

* Implement the `Strategy` interface:

```go
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
```

* Add the strategy to the `TestStrategy` in `src/test/strategies_test.go` function to test that the implementation is right.

* Add the strategy to the array `allStrategies` in `main.go` to have it available in games and tournaments.
