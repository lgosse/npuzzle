package main

import (
	"math/rand"
	"time"
)

// GeneratePuzzle returns a valid (not certainly solvable) instance of Puzzle
func GeneratePuzzle() *Puzzle {
	rand.Seed(time.Now().UnixNano())
	values, sz := generateValues()
	m := make([][]int, sz)

	for i := range m {
		m[i] = make([]int, sz)
		for j := range m[i] {
			idx := rand.Intn(len(values))
			n := values[idx]
			values = append(values[:idx], values[idx+1:]...)
			m[i][j] = n
		}
	}

	return &Puzzle{m: m, s: sz}
}

func generateValues() ([]int, int) {
	sz := rand.Intn(RNDMAXSIZE-3) + 3
	values := make([]int, sz*sz)

	for i := range values {
		values[i] = i
	}

	return values, sz
}
