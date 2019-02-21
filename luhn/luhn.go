package luhn

import (
	"strconv"
	"strings"
)

func Valid(str string) bool {
	// Remove spaces
	str = strings.Replace(str, " ", "", -1)
	if len(str) < 2 {
		return false
	}
	var sum int64
	isOdd := len(str)%2 != 0

	for i := len(str) - 1; i >= 0; i-- {
		n, err := strconv.ParseInt(str[i:i+1], 0, 64)
		if err != nil {
			return false
		}
		if isOdd {
			if i%2 != 0 {
				n = n * 2
				if n > 9 {
					n = n - 9
				}
			}
		} else {
			if i%2 == 0 {
				n = n * 2
				if n > 9 {
					n = n - 9
				}
			}
		}
		sum = sum + n
	}
	return (sum%10 == 0)
}