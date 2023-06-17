https://exercism.org/tracks/go/exercises/sieve

package main

func main(){
 fmt.Println(Sieve(100))
}

func Sieve(limit int) []int {
	isPrime := make([]bool, limit+1)
	for i := 2; i <= limit; i++ {
		isPrime[i] = true
	}

	
}