package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
	"log"
)

var finalState node

// Solve finds the solution (if it exists) for the provided puzzle
func Solve(puzzle *Puzzle) error {
	finalState = computeFinalState(puzzle.m)

	return nil
}

func computeFinalState(m [][]int) node {
	l := len(m)
	f := make([][]int, l)
	values := make([]int, l*l)
	for i := range values {
		values[i] = i + 1
	}
	values[l*l-1] = 0

	curX := 0
	curY := 0
	curDir := RIGHT

	for i := range f {
		f[i] = make([]int, l)
	}

	for _, v := range values {
		f[curY][curX] = v

		switch curDir {
		case RIGHT:
			{
				if curX+1 < l && f[curY][curX+1] == 0 {
					curX++
				} else {
					curY++
					curDir = DOWN
				}

				break
			}

		case DOWN:
			{
				if curY+1 < l && f[curY+1][curX] == 0 {
					curY++
				} else {
					curX--
					curDir = LEFT
				}

				break
			}

		case LEFT:
			{
				if curX-1 >= 0 && f[curY][curX-1] == 0 {
					curX--
				} else {
					curY--
					curDir = UP
				}

				break
			}

		case UP:
			{
				if curY-1 >= 0 && f[curY-1][curX] == 0 {
					curY--
				} else {
					curX++
					curDir = RIGHT
				}

				break
			}
		}
	}

	return node{
		state: f,
		hash:  fmt.Sprintf("%x", sha256.Sum256([]byte(fmt.Sprintf("%#v", f)))),
	}
}

func getPossibilities(m nodeState) []node {
	var x int
	var y int
	var nodes []node
	var l = len(m)

Loop:
	for i, ln := range m {
		for j, v := range ln {
			if v == 0 {
				y, x = i, j
				break Loop
			}
		}
	}

	if x != l-1 {
		nodes = append(nodes, getPossibility(m, RIGHT, x, y))
	}
	if x != 0 {
		nodes = append(nodes, getPossibility(m, LEFT, x, y))
	}
	if y != 0 {
		nodes = append(nodes, getPossibility(m, UP, x, y))
	}
	if y != l-1 {
		nodes = append(nodes, getPossibility(m, DOWN, x, y))
	}

	return nodes
}

func getPossibility(m nodeState, move Action, x int, y int) node {
	var tmp int

	var mod bytes.Buffer
	enc := gob.NewEncoder(&mod)
	dec := gob.NewDecoder(&mod)

	err := enc.Encode(m)
	if err != nil {
		log.Fatal("npuzzle: solve: encode error: ", err)
	}

	var cpy [][]int
	err = dec.Decode(&cpy)
	if err != nil {
		log.Fatal("npuzzle: solve: decode error: ", err)
	}

	switch move {
	case UP:
		{
			tmp = cpy[y][x]
			cpy[y][x] = cpy[y-1][x]
			cpy[y-1][x] = tmp
		}

	case RIGHT:
		{
			tmp = cpy[y][x]
			cpy[y][x] = cpy[y][x+1]
			cpy[y][x+1] = tmp
		}

	case DOWN:
		{
			tmp = cpy[y][x]
			cpy[y][x] = cpy[y+1][x]
			cpy[y+1][x] = tmp
		}

	case LEFT:
		{
			tmp = cpy[y][x]
			cpy[y][x] = cpy[y][x-1]
			cpy[y][x-1] = tmp
		}

	default:
		{
			log.Fatal(fmt.Sprintf(
				"npuzzle: solve: move \"%v\" does not exists",
				move,
			))
		}
	}

	return node{
		hash:  fmt.Sprintf("%x", sha256.Sum256([]byte(fmt.Sprintf("%#v", cpy)))),
		state: cpy,
	}
}
