package main

import (
	"fmt"
	"math/big"
)

func main() {

	// Central binomial coefficients: C(2*n,n) = (2*n)!/(n!n!) 
	const n = 20
	res, num, den := new(big.Int), new(big.Int), new(big.Int)
	num.MulRange(n+1, 2*n)
	den.MulRange(2, n)
	res.Div(num, den)
	fmt.Println(res)
}
