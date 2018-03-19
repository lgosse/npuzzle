package main

import (
	"fmt"
	"os"
)

func handleArgs() (*Puzzle, error) {
	if len(os.Args) < 2 {
		return GeneratePuzzle(), nil
	} else if len(os.Args) > 2 {
		fmt.Fprintf(os.Stderr, "npuzzle: Too many arguments\n")
		return nil, nil
	} else {
		return Parse(os.Args[1])
	}
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

	solution, err := Solve(puzzle)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())

		return
	}

	fmt.Println(solution)

	return
}
