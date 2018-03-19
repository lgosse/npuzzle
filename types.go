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

// String prints informations on Puzzle struct
func (p *Puzzle) String() string {
	s := fmt.Sprintf("Size: %v", p.s)
	tab := make([]string, 0)

	for i, ln := range p.m {
		tab = append(tab, fmt.Sprintf("%v | ", i))
		for _, elem := range ln {
			tab = append(tab, fmt.Sprintf("%4v", elem))
		}
		tab = append(tab, "\n\n")
	}

	return fmt.Sprintf("%s\nMap:\n%s", s, strings.Join(tab, ""))
}

// Solution is the data structure containing solution to the puzzle
type Solution struct {
}
