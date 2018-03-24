package main

import (
	"fmt"
	"os"
	"strings"
)

func usage() string {
	return fmt.Sprintf(
		"usage: npuzzle [OPTION] HEURISTIC [FILE]\nAvailable heuristics:\n%s%s%s%s%s%s%s%s",
		" - manhattan\n",
		" - misplaced\n",
		" - linear-conflict\n",
		" - permutation\n",
		" - tiles-out\n",
		"\nOPTIONS:\n",
		"  -g --greedy  Use greedy search algorithm\n",
		"  -u --uniform Use uniform cost algorithm\n",
	)
}

var greedySearch = false
var uniformSearch = false

func handleArgs() (*Puzzle, error) {
	heuristicIdx := 1

	if len(os.Args) == 1 {
		return nil, fmt.Errorf(usage())
	}

	if strings.Compare(os.Args[1], "-g") == 0 || strings.Compare(os.Args[1], "--greedy") == 0 {
		greedySearch = true
		heuristicIdx++
	} else if os.Args[1] == "-u" || os.Args[1] == "--uniform" {
		uniformSearch = true
		heuristicIdx++
	}

	switch os.Args[heuristicIdx] {
	case MANHATTAN:
		selectedHeuristic = ManhattanHeuristic
		break
	case MISPLACED:
		selectedHeuristic = MisplacedHeuristic
		break
	case LINEAR:
		selectedHeuristic = LinearConflictHeuristic
		break
	case PERMUTATION:
		selectedHeuristic = PermutationHeuristic
		break
	case TILESOUT:
		selectedHeuristic = TilesOutHeuristic
		break
	default:
		return nil, fmt.Errorf(usage())
	}

	if len(os.Args) == heuristicIdx+1 {
		return GeneratePuzzle(), nil
	} else if len(os.Args) == heuristicIdx+2 {
		return Parse(os.Args[heuristicIdx+1])
	}

	return nil, nil
}

func main() {
	var puzzle *Puzzle
	var err error

	puzzle, err = handleArgs()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())

		return
	}

	if puzzle.IsValid() == false {
		fmt.Println("Error: le puzzle n'est pas solvable")
		return
	}

	fmt.Println(puzzle)

	err = Solve(puzzle)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())

		return
	}

	return
}
