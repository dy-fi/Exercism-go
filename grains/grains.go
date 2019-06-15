package grains

import (
	"errors"
)

// Square returns 2^n
func Square(n int) (uint64, error) {

	if n < 1 || n > 64 {
		return 0, errors.New("Invalid size")
	}

	sum := 1 << uint64(n - 1)

	return uint64(sum), nil
}

// Total returns the total after 64 days 
func Total() uint64 {
	var expected uint64 = 18446744073709551615
	return expected
}

