package main

import "fmt"

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}


func main() {
	ret := 1
	for x := 2; x<= 20 ; x++ {
		ret = lcm(ret, x)
	}
	fmt.Println(ret)
}
