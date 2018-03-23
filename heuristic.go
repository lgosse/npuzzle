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
	heuristic := ManhattanHeuristic(state)

	vert := verticalLinearConflict(state)

	hor := horizontalLinearConflict(state)

	heuristic += vert + hor

	return heuristic
}

func verticalLinearConflict(state nodeState) int {
	nb := 0
	for x := range state {
		for y := range state[x] {
			if state[y][x] != finalState.state[y][x] {
				for i := 0; i < len(finalState.state); i++ {
					if hasToPermuteVertical(x, y, i, state) {
						nb += 2
					}
				}
			}
		}
	}

	return nb
}

func horizontalLinearConflict(state nodeState) int {
	nb := 0
	for y := range state {
		for x := range state[y] {
			if state[y][x] != finalState.state[y][x] {
				for i := 0; i < len(finalState.state); i++ {
					if hasToPermuteHorizontal(y, x, i, state) {
						nb += 2
					}
				}
			}
		}
	}

	return nb
}

func hasToPermuteHorizontal(y, x, x2 int, state nodeState) bool {
	fstTileInRow := false
	idxFst := 0
	scdTileInRow := false
	idxScd := 0

	for i := range state[y] {
		if state[y][x] == state[y][i] {
			fstTileInRow = true
			idxFst = i
		}
		if state[y][x2] == state[y][i] {
			scdTileInRow = true
			idxScd = i
		}
	}

	if fstTileInRow && scdTileInRow {
		if idxFst > idxScd {
			return true
		}
	}

	return false
}

func hasToPermuteVertical(x, y, y2 int, state nodeState) bool {
	fstTileInCol := false
	idxFst := 0
	scdTileInCol := false
	idxScd := 0

	for i := range state[y] {
		if state[y][x] == state[i][x] {
			fstTileInCol = true
			idxFst = i
		}
		if state[y2][x] == state[i][x] {
			scdTileInCol = true
			idxScd = i
		}
	}

	if fstTileInCol && scdTileInCol {
		if idxFst > idxScd {
			return true
		}
	}

	return false
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

// PermutationHeuristic calculate the sum of permutations to execute in
// order to get every pieces in the right place
func PermutationHeuristic(state nodeState) int {
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
	return nb
}
