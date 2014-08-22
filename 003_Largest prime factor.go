package main

import ("fmt" 
		"math")

func main() {
 
 	num := 600851475143
 	limit := int(math.Sqrt(float64(num)))

	for i := 2; i < limit; i++ {
		for num % i == 0 {
			num /= i
		}
		
		if num == 1 {
			fmt.Println(i)
			break
		}
	}
}
