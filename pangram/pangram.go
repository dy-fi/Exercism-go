package pangram

import (
	"strings"
)

// IsPangram returns whether a string is a pangram in a bool
func IsPangram(s string) bool {

	if len(s) < 26 {
		return false
	}

	alphabet := make(map[string]bool)
	s = strings.ToLower(s)

	for _, v := range s {
		for i := 'a'; i >= 'a' && i <= 'z'; i++ {

			if v == i {
				alphabet[string(i)] = true
			}
		}
	}

	for _, v := range alphabet {
		if v == false || len(alphabet) < 26 {
			return false
		}
	}

	return true
}
