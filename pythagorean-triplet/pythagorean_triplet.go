package pythagorean

import "math"

// Triplet is a 3 length array that holds pythagorean triples
type Triplet [3]int

// Range returns all Pythagorean triplets with sides in a given range min-max
func Range(min int, max int) (result []Triplet) {

	for i := min; i < max; i++ {
		for j := max; i > min; i-- {
			sum := math.Sqrt(float64(i * i + j * j))

			if sum == float64(int64(sum)) {
				result = append(result, Triplet{i, j, int(sum)})
			}
		}
	}
	return
}

// Sum returns a list of all Pythagorean triplets where their sum == p
func Sum(p int) (result []Triplet) {
	triplets := Range(0, p)

	for i := len(triplets); i > 0; i-- {
		if triplets[i][0] + triplets[i][1] + triplets[i][2] == p {
			result = append(result, triplets[i])
		}
	}
	return
}

