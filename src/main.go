package main

import (
	"fmt"
	"os"
	"strings"
)

func usage() string {
	return fmt.Sprintf(
		"usage: npuzzle [OPTION] HEURISTIC [FILE]\nAvailable heuristics:\n%s%s%s%s%s%s%s%s%s",
		" - manhattan\n",
		" - misplaced\n",
		" - linear-conflict\n",
		" - permutation\n",
		" - tiles-out\n",
		"\nOPTIONS:\n",
		"  -g --greedy      Use greedy search algorithm. Cannot be used with --uniform.\n",
		"  -u --uniform     Use uniform cost algorithm. Cannot be used with --greedy.\n",
		"  -m --multithread Multithread the search algorithm.",
	)
}

var greedySearch = false
var uniformSearch = false
var nbGoRoutines = 1

func handleOptions(arg string) bool {
	if strings.HasPrefix(arg, "--") {
		if arg == OPTGREEDYLONG && !uniformSearch {
			greedySearch = true

			return true
		} else if arg == OPTUNIFORMLONG && !greedySearch {
			uniformSearch = true

			return true
		} else if arg == OPTMULTILONG {
			nbGoRoutines = NBGOROUTINES

			return true
		}
	} else if strings.HasPrefix(arg, "-") {
		valid := true
		for _, v := range arg {
			if v == OPTGREEDYSHORT && !uniformSearch {
				greedySearch = true
			} else if v == OPTUNIFORMSHORT && !greedySearch {
				uniformSearch = true
			} else if v == OPTMULTISHORT {
				nbGoRoutines = NBGOROUTINES
			} else if v != '-' {
				valid = false
			}
		}

		return valid
	}

	return false
}

func handleArgs() (*Puzzle, error) {
	argIdx := 1

	if len(os.Args) == 1 {
		return nil, fmt.Errorf(usage())
	}

	for len(os.Args) > argIdx && handleOptions(os.Args[argIdx]) {
		argIdx++
	}

	if len(os.Args) == argIdx {
		return nil, fmt.Errorf(usage())
	}

	switch os.Args[argIdx] {
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

	if len(os.Args) == argIdx+1 {
		return GeneratePuzzle(), nil
	} else if len(os.Args) == argIdx+2 {
		return Parse(os.Args[argIdx+1])
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

	finalState = computeFinalState(len(puzzle.m))

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
