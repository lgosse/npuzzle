package main

import (
	"fmt"
	"sync"
)

// Puzzle is the data structure where the puzzle to solve is stored
type Puzzle struct {
	m   [][]int
	s   int
	mut sync.Mutex
}

// String prints informations on Puzzle struct
func (p *Puzzle) String() string {
	return fmt.Sprintf("Size: %v\nMap:\n%s", p.s, FmtPuzzleState(p.m))
}

// Action describes a direction choose to solve the puzzle
type Action int

// Actions
const (
	UP = iota
	RIGHT
	DOWN
	LEFT
)

// Node describes a Puzzle's state
type Node struct {
	g      int
	h      int
	hash   string
	m      [][]int
	parent *Node
}
