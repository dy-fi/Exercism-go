package leap

//IsLeapYear calculates whether the given year is a leap year in the Gregorian calendar 
func IsLeapYear(n int) bool {
	if n % 4 == 0 {
		if n % 100 == 0 {
			if n % 400 == 0 {
				return true 
			}
			return false 
		}
		return true
	}
	return false
}
