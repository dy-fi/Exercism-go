package reverse

// String haskjf;asdfkjl
func String(s string) string {
	var result string

	for _, v := range s {
		result = string(v) + result
	}
	return result
}
