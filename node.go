package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
	"strings"
	"sync"
)

type nodeMap struct {
	sync.Mutex
	nodes map[string]*node
}

// get retrieves a node
func (nm *nodeMap) get(n node) *node {
	nm.Lock()
	m, ok := nm.nodes[n.hash]
	nm.Unlock()

	if !ok {
		m = &node{
			hash:  n.hash,
			state: n.state,
		}

		nm.Lock()
		nm.nodes[n.hash] = m
		nm.Unlock()
	}

	return m
}

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

func (state nodeState) Copy() nodeState {
	var mod bytes.Buffer
	enc := gob.NewEncoder(&mod)
	dec := gob.NewDecoder(&mod)

	err := enc.Encode(state)
	if err != nil {
		log.Fatal("nodeState: copy: encode error: ", err)
	}

	var cpy [][]int
	err = dec.Decode(&cpy)
	if err != nil {
		log.Fatal("nodeState: copy: decode error: ", err)
	}

	return cpy
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

func (n node) String() string {
	return fmt.Sprintf(
		"state:\n%s\nhash: %v\ncost: %v\nrank: %v\nopen: %v\nclosed: %v\nindex: %v\n",
		n.state,
		n.hash,
		n.cost,
		n.rank,
		n.open,
		n.closed,
		n.index,
	)
}

// Heuristic calls the selected heuristic
func (n node) Heuristic() float64 {
	return selectedHeuristic(n.state)
}
