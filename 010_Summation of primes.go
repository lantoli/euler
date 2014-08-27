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

	const total = 2000000

	primes := make([]int, 0)
	sum := 0
	for num := 2 ; num < total ; num++ {
		if IsPrime(primes, num) {
			primes = append(primes, num)
			sum += num
		}
	}
	fmt.Println(sum)
}
