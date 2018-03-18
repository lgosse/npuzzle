package main

import (
	"fmt"
	"strings"
	"sync"
)

// Puzzle is the data structure where the puzzle to solve is stored
type Puzzle struct {
	m   [][]int
	s   int
	mut sync.Mutex
}

// Solve solves the puzzle
func (p *Puzzle) solve() {

}

// String prints informations on Puzzle struct
func (p *Puzzle) String() string {
	s := fmt.Sprintf("Puzzle size: %v", p.s)
	tab := make([]string, 0)

	for i, v := range p.m {
		tab = append(tab, fmt.Sprintf("%-4v %v\n", i, v))
	}

	return fmt.Sprintf("%s\nMap:\n%s", s, strings.Join(tab, "\n"))
}
