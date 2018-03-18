package main

import (
	"fmt"
	"os"
)

func main() {
	puzzle, err := Parse()

	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())

		return
	}
	fmt.Println(puzzle)

	puzzle.solve()
	return
}
