package main

import "fmt" 
		
func main() {
 
 	num := 600851475143

	for i := 2; i <= num; i++ {
		for num % i == 0 {
			num /= i
		}
		
		if num == 1 {
			fmt.Println(i)
			break
		}
	}
}
