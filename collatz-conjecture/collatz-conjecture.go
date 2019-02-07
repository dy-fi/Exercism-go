package collatzconjecture

import "errors"

// CollatzConjecture finds the steps required to get from a given integer n to 1 in the conjecture
func CollatzConjecture(n int) (int, error) {

	if n == 0 || n < 0 {
		return 0, errors.New("Invalid value of N")
	}

	cursor := n
	counter := 0

	for cursor != 1 {
		if cursor%2 == 0 {
			cursor = cursor / 2
		} else {
			if cursor == 0 {
				cursor++
			}
			cursor = (cursor * 3) + 1
		}
		counter++
	}
	return counter, nil
}
