package main

import "fmt"
	 
func main() {

	const lim = 1000

	for c := 1; c < lim; c++ {
		for b := 1; b < c; b++ {
			a := lim - b - c
			if a >= 1 && a < b {
				if a + b + c == lim && a*a + b*b == c*c {
					fmt.Println(a*b*c)
					return
				}				
			}
		}
	}
}
