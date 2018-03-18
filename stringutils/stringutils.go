package stringutils

import (
	"strconv"
	"strings"
)

// ToIntArray converts string into int array
func ToIntArray(s string) ([]int, error) {
	a := strings.Fields(s)
	b := make([]int, len(a))
	for i, v := range a {
		var err error
		b[i], err = strconv.Atoi(v)

		if err != nil {
			return make([]int, 0), err
		}
	}

	return b, nil
}
