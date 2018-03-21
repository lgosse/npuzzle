package main

import (
	"fmt"
	"strings"
)

type nodeState [][]int

func (state nodeState) String() string {
	tab := make([]string, 0)

	for i, ln := range state {
		tab = append(tab, fmt.Sprintf("%v | ", i))
		for _, elem := range ln {
			tab = append(tab, fmt.Sprintf("%4v", elem))
		}
		tab = append(tab, "\n\n")
	}

	return strings.Join(tab, "")
}

type node struct {
	state  nodeState
	hash   string
	cost   float64
	rank   float64
	parent *node
	open   bool
	closed bool
	index  int
}

// Heuristic calls the selected heuristic
func (n node) Heuristic() {
	selectedHeuristic(n.state)
}
