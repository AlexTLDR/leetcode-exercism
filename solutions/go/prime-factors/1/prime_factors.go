package prime

func Factors(n int64) []int64 {
	var factors []int64
	var i int64 = 2
	for n > 1 {
		if n%i == 0 {
			factors = append(factors, i)
			n /= i
		} else {
			i++
		}
	}
	return factors
}
