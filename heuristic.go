package main

import (
	"math"
)

// selectedHeuristic represents the heuristic selected by the end user
var selectedHeuristic func(state nodeState) float64

// ManhattanHeuristic calculates the distance between each piece
// and its final position
func ManhattanHeuristic(state nodeState) float64 {
	total := 0.0

	for y := range state {
		for x := range state[y] {

			toFind := state[y][x]

		Find:
			for by := range finalState.state {
				for bx := range finalState.state[by] {
					if finalState.state[by][bx] == toFind {
						total += math.Pow(float64(x-bx), 2) + math.Pow(float64(y-by), 2)
						break Find
					}
				}
			}

		}
	}

	return total / float64(len(state))
}

// MisplacedHeuristic calculates the number of misplaced pieces
func MisplacedHeuristic(state nodeState) float64 {
	misplacedElems := 0

	for y := range state {
		for x := range state[y] {
			if state[y][x] != finalState.state[y][x] {
				misplacedElems++
			}
		}
	}

	return float64(misplacedElems)
}

// LinearConflictHeuristic is the ManhattanHeuristic ponderated by
// the conflict between some pieces and their final destination
func LinearConflictHeuristic(state nodeState) float64 {
	return 1
}
