// Sieve of Atkin algorithm
//
// A variation of prime sieves using binary quadratic forms.
// Reference: https://en.wikipedia.org/wiki/Sieve_of_Atkin

package main

import (
	"math"
    "fmt"
)

func Sieve(N int) (primes []int) {
	var x, y, n int
	nsqrt := math.Sqrt(float64(N))

	isPrime := make([]bool, N)

	for x = 1; float64(x) <= nsqrt; x++ {
		for y = 1; float64(y) <= nsqrt; y++ {
			n = 4*(x*x) + y*y
			if n <= N && (n%12 == 1 || n%12 == 5) {
				isPrime[n] = !isPrime[n]
			}

			n = 3*(x*x) + y*y
			if n <= N && n%12 == 7 {
				isPrime[n] = !isPrime[n]
			}

			n = 3*(x*x) - y*y
			if x > y && n <= N && n%12 == 11 {
				isPrime[n] = !isPrime[n]
			}
		}
	}

	for n = 5; float64(n) <= nsqrt; n++ {
		if isPrime[n] {
			for y = n * n; y < N; y += n * n {
				isPrime[y] = false
			}
		}
	}

	isPrime[2] = true
	isPrime[3] = true

	for x = 0; x < len(isPrime)-1; x++ {
		if isPrime[x] {
			primes = append(primes, x)
		}
	}

	return
}

func main() {
	// Only primes less than or equal to N will be generated
	const N = 100
	primes := Sieve(N)

	// Primes is now a slice that contains all primes numbers up to N.
	// So let's print them
	for _, x := range primes {
		fmt.Println(x)
	}
}
