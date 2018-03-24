package main

import (
	"fmt"
	"os"
)

func usage() string {
	return fmt.Sprintf(
		"usage: npuzzle HEURISTIC [file]\nAvailable heuristics:\n%s%s%s%s%s",
		" - manhattan\n",
		" - misplaced\n",
		" - linear-conflict\n",
		" - permutation\n",
		" - tiles-out\n",
	)
}

func handleArgs() (*Puzzle, error) {
	if len(os.Args) == 1 {
		return nil, fmt.Errorf(usage())
	}

	switch os.Args[1] {
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

	if len(os.Args) == 2 {
		return GeneratePuzzle(), nil
	} else if len(os.Args) == 3 {
		return Parse(os.Args[2])
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
