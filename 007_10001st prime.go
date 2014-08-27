package main

import "fmt"
	 
func IsPrime(prevPrimes []int, candidate int) bool {
	for _, prime := range prevPrimes {
		if prime * prime > candidate {
			return true
		}
		if candidate % prime == 0  {
			return false
		}
	}
	return true
}

func main() {

	const total = 10001

	primes := make([]int, 0, total)

	for num := 2 ; len(primes) < total ; num++ {
		if IsPrime(primes, num) {
			primes = append(primes, num)
		}
	}
	fmt.Println(primes[len(primes)-1])
}
