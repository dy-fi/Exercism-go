package summultiples

//SumMultiples gets the sum of the multiples of divisors up to 
func SumMultiples(n int, multiples []int) int {
	// get the multiples and append them to mulList
	var mulList []int
	for i := 0; i < n; i++ {
		for _,v := range(multiples) {
			if i % v == 0 {
				mulList = append(mulList, i)
			}
		}
	}

	// filter multiples
	var sorted []int
	for _,v := range(mulList) {
		for _,h := range(sorted) {
			// no multiples
			if v == h {
				 break 
			} else {
				sorted = append(sorted, v)
			}
		}
	}

	// get the sum
	var sum int
	for _,v := range(sorted) {
		sum += v
	}
	return sum
}