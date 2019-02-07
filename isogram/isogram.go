package isogram

import (
	"strings"
)

// IsIsogram returns true if the word is an isogram, and false if it doesn't
func IsIsogram(word string) bool {
	mem := strings.ToLower(word)

	for i := 0; i < len(mem); i++ {
		for j := 0; j < len(mem); j++ {
			test := string(mem[i])

			if mem[i] == mem[j] && i != j {
				if test != "-" && test != " " {
					return false
				}
			}
		}
	}
	return true
}
