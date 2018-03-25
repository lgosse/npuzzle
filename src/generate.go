package main

import (
	"math/rand"
	"time"
)

func shake(state nodeState) nodeState {
	var x, y, tmp int
	var dirs [4]bool
	var dir int

	rand.Seed(time.Now().UTC().UnixNano())
Loop:
	for i, ln := range state {
		for j, v := range ln {
			if v == 0 {
				y, x = i, j

				break Loop
			}
		}
	}

	if x != PUZZLESZ-1 {
		dirs[RIGHT] = true
	}
	if x != 0 {
		dirs[LEFT] = true
	}
	if y != 0 {
		dirs[UP] = true
	}
	if y != PUZZLESZ-1 {
		dirs[DOWN] = true
	}

	dir = rand.Intn(4)
	for dirs[dir] != true {
		dir = rand.Intn(4)
	}

	switch dir {
	case UP:
		{
			tmp = state[y][x]
			state[y][x] = state[y-1][x]
			state[y-1][x] = tmp
		}

	case RIGHT:
		{
			tmp = state[y][x]
			state[y][x] = state[y][x+1]
			state[y][x+1] = tmp
		}

	case DOWN:
		{
			tmp = state[y][x]
			state[y][x] = state[y+1][x]
			state[y+1][x] = tmp
		}

	case LEFT:
		{
			tmp = state[y][x]
			state[y][x] = state[y][x-1]
			state[y][x-1] = tmp
		}
	}

	return state
}

// GeneratePuzzle returns a valid (not certainly solvable) instance of Puzzle
func GeneratePuzzle() *Puzzle {
	state := computeFinalState(PUZZLESZ).state

	for i := 0; i < NBITERATIONS; i++ {
		state = shake(state)
	}

	return &Puzzle{m: state, s: PUZZLESZ}
}
