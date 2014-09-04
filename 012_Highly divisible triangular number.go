package main

import "fmt"
	
func NumberOfDivisor(num int) int {	
	//see http://mathschallenge.net/library/number/number_of_divisors
	ret := 1
	for i := 2; i <= num; i++ {
		factors := 0
		for num % i == 0 {
			num /= i
			factors++
		}
		ret *= factors + 1
		
		if num == 1 {
			break
		}
	}
	return ret
}

func main() {

	const divisors = 500

	sum := 0
	for i := 1; NumberOfDivisor(sum)<= divisors; i++ {
		sum += i
	}
	fmt.Println(sum)
}
