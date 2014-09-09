package main

import "fmt"
	

func f(y, x int) int {

	if y == x {
		return 2 * f(y, x-1)
	}

	if y < x {
		y, x = x, y
	}

	if x == 1 {
		return y + 1
	}
	return f(x-1, y) + f(x, y-1)
}


func main() {
	fmt.Println(f(5,5))
}
