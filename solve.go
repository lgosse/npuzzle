package main

import (
	"container/heap"
	"crypto/sha256"
	"fmt"
	"log"
	"sync"
)

var finalState node

// Solve finds the solution (if it exists) for the provided puzzle
func Solve(puzzle *Puzzle) error {
	var successNode *node

	finalState = computeFinalState(puzzle.m)
	nmap := nodeMap{}
	initialState := nmap.get(node{
		hash:  hashNodeState(puzzle.m),
		state: puzzle.m,
	})
	initialState.open = true

	pq := &priorityQueue{c: sync.NewCond(new(sync.Mutex))}
	heap.Init(pq)

	quitChan = watchHeapOps()

	done := make(chan bool, 1)
	HeapPushSync(pq, initialState, done)
	<-done

	fmt.Printf("################# BEGIN ALGO ###############\n")
	winChan := make(chan *node, 10)
	errChan := make(chan error, 10)
	go astar(pq, nmap, winChan, errChan, 1)
	// go astar(pq, nmap, winChan, errChan, 2)

	select {
	case successNode := <-winChan:
		fmt.Printf("success: %v\nlen: %v\n", successNode, pq.Len())
	case err := <-errChan:
		return err
	}

	stopWatchHeapOps()

	fmt.Printf("success: %v\nlen: %v\n", successNode, pq.Len())
	fmt.Printf("################# END ALGO ###############\n")

	return nil
}

func astar(pq *priorityQueue, nmap nodeMap, winChan chan *node, errChan chan error, id int) {
	first := true
	for pq.Len() != 0 || first {
		first = false
		curState := HeapPop(pq).(*node)
		curState.open = false
		curState.closed = true

		if curState.hash == finalState.hash {
			winChan <- curState
			close(winChan)

			return
		}

		for _, v := range getPossibilities(curState.state) {
			cost := curState.cost + 1
			possibleState := nmap.get(v)

			if cost < possibleState.cost && possibleState.open {
				HeapRemove(pq, possibleState.index)
				possibleState.open = false
				possibleState.closed = false
			}

			if !possibleState.open && !possibleState.closed {
				possibleState.cost = cost
				possibleState.open = true
				possibleState.rank = cost + possibleState.Heuristic()
				possibleState.parent = curState

				done := make(chan bool)
				HeapPushSync(pq, possibleState, done)
				<-done
			}
		}
	}
}

func computeFinalState(m nodeState) node {
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
		hash:  hashNodeState(f),
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

	cpy := m.Copy()

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
		hash:  hashNodeState(cpy),
		state: cpy,
	}
}

func hashNodeState(state nodeState) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(fmt.Sprintf("%#v", state))))
}
