package main

import "fmt"

func main() {
	sum, sumsqr := 0, 0
	for i := 1; i <= 100; i++ {
		sum += i
		sumsqr += i*i
	}
	fmt.Println(sum*sum - sumsqr)
}
