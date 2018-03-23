package main

import (
	"math"
)

// selectedHeuristic represents the heuristic selected by the end user
var selectedHeuristic func(state nodeState) int

// ManhattanHeuristic calculates the distance between each piece
// and its final position
func ManhattanHeuristic(state nodeState) int {
	total := 0

	for y := range state {
		for x := range state[y] {

			toFind := state[y][x]

		Find:
			for by := range finalState.state {
				for bx := range finalState.state[by] {
					if finalState.state[by][bx] == toFind {
						total += int(math.Abs(float64(bx-x)) + math.Abs(float64(by-y)))
						break Find
					}
				}
			}

		}
	}

	return total
}

// MisplacedHeuristic calculates the number of misplaced pieces
func MisplacedHeuristic(state nodeState) int {
	misplacedElems := 0

	for y := range state {
		for x := range state[y] {
			if state[y][x] != finalState.state[y][x] {
				misplacedElems++
			}
		}
	}

	return misplacedElems
}

// LinearConflictHeuristic is the ManhattanHeuristic ponderated by
// the conflict between some pieces and their final destination
func LinearConflictHeuristic(state nodeState) int {
	return 1
}
