package main

import "fmt"

// FormatError allow defining a format error on specified line
type FormatError struct {
	l int
	s string
}

func (f FormatError) Error() string {
	return fmt.Sprintf("parse: format error on line %v\n    \"%v\"", f.l, f.s)
}

// UncaughtError allow defining an uncaught error
type UncaughtError struct {
	s string
}

func (f UncaughtError) Error() string {
	return fmt.Sprintf("parse: Uncaught error in %v\n", f.s)
}
