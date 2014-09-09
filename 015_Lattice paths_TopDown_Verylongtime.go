package main

import "fmt"
	

func f(y, x int) int {
	if y == 1 {
		return x + 1
	}
	if x == 1 {
		return y + 1
	}
	return f(x-1, y) + f(x, y-1)
}


func main() {
	fmt.Println(f(20,20))
}
