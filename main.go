package main

import (
	"fmt"
	"os"
)

func usage() string {
	return fmt.Sprintf(
		"usage: npuzzle HEURISTIC [file]\nAvailable heuristics:\n%s%s%s",
		" - manhattan\n",
		" - misplaced\n",
		" - manhattan+\n",
	)
}

func handleArgs() (*Puzzle, error) {
	if len(os.Args) == 1 {
		return nil, fmt.Errorf(usage())
	}

	switch os.Args[1] {
	case MANNHATAN:
		selectedHeuristic = ManhattanHeuristic
		break
	case MISPLACED:
		selectedHeuristic = MisplacedHeuristic
		break
	case LINEAR:
		selectedHeuristic = LinearConflictHeuristic
		break
	default:
		return nil, fmt.Errorf(usage())
	}

	if len(os.Args) == 2 {
		return GeneratePuzzle(), nil
	} else if len(os.Args) == 3 {
		return Parse(os.Args[2])
	}

	fmt.Fprintf(os.Stderr, "npuzzle: Too many arguments\n")
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

	fmt.Println(puzzle)

	err = Solve(puzzle)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())

		return
	}

	return
}
