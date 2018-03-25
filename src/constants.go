package main

// PUZZLESZ represents the generated puzzle size
const PUZZLESZ = 3

// NBITERATIONS represents the number of iterations shaking the
// solved puzzle to obtain the final puzzle to solve
const NBITERATIONS = 20000

// Defines the different implemented heuristics available
const (
	MANHATTAN   = "manhattan"
	MISPLACED   = "misplaced"
	LINEAR      = "linear-conflict"
	PERMUTATION = "permutation"
	TILESOUT    = "tiles-out"
)

// NBGOROUTINES represents the number of goroutines to be launched to solve the puzzle
const NBGOROUTINES = 4

// Define options strings constants
const (
	OPTGREEDYLONG   = "--greedy"
	OPTGREEDYSHORT  = 'g'
	OPTUNIFORMLONG  = "--uniform"
	OPTUNIFORMSHORT = 'u'
	OPTMULTILONG    = "--multithread"
	OPTMULTISHORT   = 'm'
)
