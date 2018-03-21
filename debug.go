package main

import (
	"fmt"
	"strings"
)

// FmtPuzzleState formats a puzzle state represented by [][]int
// to a human readable format with break-lines
func FmtPuzzleState(m [][]int) string {
	tab := make([]string, 0)

	for i, ln := range m {
		tab = append(tab, fmt.Sprintf("%v | ", i))
		for _, elem := range ln {
			tab = append(tab, fmt.Sprintf("%4v", elem))
		}
		tab = append(tab, "\n\n")
	}

	return strings.Join(tab, "")
}
