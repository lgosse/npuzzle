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

func serpentard(state nodeState) []int {
	l := len(state)

	f := make([][]int, l)
	for i := range f {
		f[i] = make([]int, l)
		copy(f[i], state[i])
	}
	snake := make([]int, l*l)
	i, j := 0, -1

	for value := 0; value < len(snake); i++ {
		for j++; j < l && f[i][j] != -1; j++ {
			snake[value] = f[i][j]
			f[i][j] = -1
			value++
		}
		j--
		for i++; i < l && f[i][j] != -1; i++ {
			snake[value] = f[i][j]
			f[i][j] = -1
			value++
		}
		i--
		for j--; j >= 0 && f[i][j] != -1; j-- {
			snake[value] = f[i][j]
			f[i][j] = -1
			value++
		}
		j++
		for i--; i >= 0 && f[i][j] != -1; i-- {
			snake[value] = f[i][j]
			f[i][j] = -1
			value++
		}
	}
	return snake
}

// PermutationHeuristic pourquoi cette merde oblige a faire des commentaires de merde inutiles ????!!!!
func PermutationHeuristic(state nodeState) float64 {
	sort := serpentard(state)
	len := len(sort)
	nb := 0
	i := 0
	for i < len {
		if sort[i] != 0 {
			j := i + 1
			for j < len {
				if sort[j] != 0 && sort[i] > sort[j] {
					nb++
				}
				j++
			}
		}
		i++
	}
	return float64(nb)
}
