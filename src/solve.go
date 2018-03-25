package main

import (
	"container/heap"
	"crypto/sha256"
	"fmt"
	"log"
	"sync"
)

var finalState node

// Solution represents the data structure used to represent the found puzzle solution
type Solution struct {
	sync.Mutex
	maxNodesExplored int
	node             *node
}

var solution = &Solution{}

// Solve finds the solution (if it exists) for the provided puzzle
func Solve(puzzle *Puzzle) error {
	nmap := nodeMap{nodes: map[string]*node{}}
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

	winChan := make(chan *node, 10)
	for i := 0; i < nbGoRoutines; i++ {
		go astar(pq, &nmap, winChan, 1)
	}

	solution.node = <-winChan

	stopWatchHeapOps()

	solution.Lock()
	curNode := solution.node
	totalCost := curNode.cost
	defer fmt.Printf("Total nodes to path: %v\nTotal nodes explored: %v\nMaximum nodes in memory: %v\n", totalCost, len(nmap.nodes), solution.maxNodesExplored)
	solution.Unlock()

	for i := totalCost; curNode != nil; i-- {
		if curNode != nil && curNode.state != nil {
			defer fmt.Printf("STEP %v:\n%v\n", i, curNode.state)
		}
		curNode = curNode.parent
	}

	return nil
}

func astar(pq *priorityQueue, nmap *nodeMap, winChan chan *node, id int) {
	var curState *node
	first := true

	for pq.Len() != 0 || first {
		first = false

		pq.c.L.Lock()
		for {
			if pq.Len() != 0 {
				curState = HeapPop(pq).(*node)
				break
			}
			pq.c.Wait()
		}
		pq.c.L.Unlock()

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
				if greedySearch {
					possibleState.rank = possibleState.Heuristic()
				} else if uniformSearch {
					possibleState.rank = cost
				} else {
					possibleState.rank = cost + possibleState.Heuristic()
				}
				possibleState.parent = curState

				done := make(chan bool)
				HeapPushSync(pq, possibleState, done)
				<-done

				solution.Lock()
				if solution.maxNodesExplored < pq.Len() {
					solution.maxNodesExplored = pq.Len()
				}
				solution.Unlock()
			}
		}
	}
}

func computeFinalState(l int) node {
	f := make([][]int, l)
	values := make([]int, l*l)
	for i := range values {
		values[i] = i + 1
	}
	values[len(values)-1] = 0
	for i := range f {
		f[i] = make([]int, l)
	}
	i, j := 0, -1

	for value := 0; value < len(values); i++ {
		for j++; j < l && f[i][j] == 0; j++ {
			f[i][j] = values[value]
			value++
		}
		j--
		for i++; i < l && f[i][j] == 0; i++ {
			f[i][j] = values[value]
			value++
		}
		i--
		for j--; j >= 0 && f[i][j] == 0; j-- {
			f[i][j] = values[value]
			value++
		}
		j++
		for i--; i >= 0 && f[i][j] == 0; i-- {
			f[i][j] = values[value]
			value++
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

// IsValid return true if the npuzzle is solvable
func (puzzle *Puzzle) IsValid() bool {
	var x1, y1, x2, y2 int
	sz := len(finalState.state)
	board := puzzle.m
	boardInv := getInversions(board)
	finalInv := getInversions(finalState.state)

	if sz%2 == 0 {
	BoardFindEmptyLoop:
		for i := range board {
			for j := range board[i] {
				if board[i][j] == 0 {
					x1, y1 = j, i

					break BoardFindEmptyLoop
				}
			}
		}

	FinalFindEmptyLoop:
		for i := range finalState.state {
			for j := range finalState.state[i] {
				if finalState.state[i][j] == 0 {
					x2, y2 = j, i

					break FinalFindEmptyLoop
				}
			}
		}

		boardInv += sz*sz - (y1*sz + x1)
		finalInv += sz*sz - (y2*sz + x2)
	}

	return boardInv%2 == finalInv%2
}

func getInversions(state [][]int) (inv int) {
	flat := flatten(state)

	for i := range flat {
		if flat[i] == 0 {
			continue
		}

		invPart := flat[i+1:]
		for j := range invPart {
			if flat[i] > invPart[j] && invPart[j] != 0 {
				inv++
			}
		}
	}

	return
}

func flatten(state [][]int) []int {
	var flat []int

	for i := range state {
		flat = append(flat, state[i]...)
	}

	return flat
}

func hashNodeState(state nodeState) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(fmt.Sprintf("%#v", state))))
}
