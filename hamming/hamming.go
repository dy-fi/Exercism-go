package hamming

import "errors"

// Distance returns the hamming distance of 2 DNA strings
func Distance(a, b string) (int, error) {
	count := 0

	if len(a) != len(b) {
		return 0, errors.New("Not same length")
	}

	for i := len(a) - 1; i >= 0; i-- {
		if a[i] != b[i] {
			count++
		}
	}

	return count, nil
}
