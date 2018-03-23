package main

// RNDMAXSIZE is the auto-generated puzzle's max size
const RNDMAXSIZE = 4

// Defines the different implemented heuristics available
const (
	MANHATTAN   = "manhattan"
	MISPLACED   = "misplaced"
	LINEAR      = "linear-conflict"
	PERMUTATION = "permutation"
)

// NBGOROUTINES represents the number of goroutines to be launched to solve the puzzle
const NBGOROUTINES = 4
