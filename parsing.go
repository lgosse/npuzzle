package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/lgosse/npuzzle/stringutils"
)

// Parse Stdin and returns a Puzzle data structure
func Parse() (*Puzzle, error) {
	data, err := read()

	if err != nil {
		return &Puzzle{}, err
	}

	if data != nil {
		puzzle, err := createPuzzle(data)

		if err != nil {
			return &Puzzle{}, err
		}

		return puzzle, err
	}

	return &Puzzle{}, UncaughtError{"parsing.go, func parse()"}
}

func read() ([]byte, error) {
	var data []byte
	buf := make([]byte, 8)
	reader := bufio.NewReader(os.Stdin)

	for {
		len, err := reader.Read(buf)

		if err != nil {
			if err.Error() != "EOF" {
				return nil, err
			}

			break
		}

		data = append(data, buf[:len]...)
	}

	return data, nil
}

func createPuzzle(data []byte) (*Puzzle, error) {
	var tab [][]int
	s, m, err := clean(data)

	if err != nil {
		return &Puzzle{}, err
	}

	if len(m) != s {
		return &Puzzle{},
			fmt.Errorf(
				"parse: size mismatch error, declared %v, got %v",
				s,
				len(m),
			)
	}

	for i, v := range m {
		a, err := stringutils.ToIntArray(v)

		if len(a) != s {
			return &Puzzle{},
				fmt.Errorf(
					"parse: length mismatch error on line %v\n    \"%v\"\nDeclared %v, got %v",
					i,
					a,
					s,
					len(a),
				)
		}

		if err != nil {
			return &Puzzle{}, err
		}

		tab = append(tab, a)
	}

	return &Puzzle{m: tab, s: s}, nil
}

func clean(data []byte) (int, []string, error) {
	var err error
	var size int
	sizeInit := false
	newLines := make([]string, 0)

	lines := strings.Split(string(data), "\n")

	for i, v := range lines {
		if v == "" {
			break
		}

		line := strings.TrimSpace(v)
		if j := strings.Index(line, "#"); j != -1 {
			line = line[:j]
		}
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}

		if strings.Trim(line, "1234567890 ") != "" {
			return 0, nil, FormatError{i, v}
		}

		if sizeInit == false {
			sizeInit = true
			size, err = strconv.Atoi(line)

			if err != nil {
				return 0, nil, err
			}

			continue
		}

		newLines = append(newLines, line)
	}

	return size, newLines, nil
}
