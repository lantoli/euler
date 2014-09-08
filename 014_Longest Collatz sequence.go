package main

import "fmt"

func LenCollatz(n int) int {
	chain := 0
	for ; n != 1; chain++ {
		if n % 2 == 0 {
			n = n / 2
		} else {
			n = n * 3 + 1
		}
	}
	return chain
}
  
func main() {
	const limit = 1000000
	max, n := 1, 1
	for i := 2; i < limit; i++ {	
		if cur := LenCollatz(i); cur > max {
			max = cur
			n = i
		}
	}
	fmt.Println(n) 
}
