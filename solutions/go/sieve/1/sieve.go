package sieve

func Sieve(limit int) []int {
	isPrime := make([]bool, limit+1)
	for i := 2; i <= limit; i++ {
		isPrime[i] = true
	}

	for i := 2; i <= limit; i++ {
		if isPrime[i] {
			for j := i * i; j <= limit; j += i {
				isPrime[j] = false
			}
		}
	}
	return getPrime(isPrime)
}

func getPrime(isPrime []bool) []int {
	primes := []int{}
	for i := 2; i < len(isPrime); i++ {
		if isPrime[i] {
			primes = append(primes, i)
		}
	}
	return primes
}

