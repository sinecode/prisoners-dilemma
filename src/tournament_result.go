package src

import (
	"sort"
)

type TournamentResult struct {
	scores map[Strategy]int
}

func NewTournamentResult() *TournamentResult {
	return &TournamentResult{
		scores: make(map[Strategy]int),
	}
}

func (tr *TournamentResult) AddScore(strategy Strategy, score int) {
	tr.scores[strategy] += score
}

func (tr *TournamentResult) ReadableRanking() ([]string, []int) {
	readableRanking := make(map[string]int)
	for strategy, score := range tr.scores {
		readableRanking[strategy.Name()] = score
	}
	return sortByValueDescending(readableRanking)
}

func sortByValueDescending(m map[string]int) ([]string, []int) {
	keys := make([]string, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}

	sort.Slice(keys, func(i, j int) bool {
		return m[keys[i]] > m[keys[j]]
	})

	values := make([]int, 0, len(m))
	for _, key := range keys {
		values = append(values, m[key])
	}

	return keys, values
}
