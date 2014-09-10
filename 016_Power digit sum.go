package main

import (
	"fmt"
	"math/big"
)

func main() {
	num := new(big.Int)
	num.Exp(big.NewInt(2), big.NewInt(1000), nil)

	total := 0
	for _, digit := range num.String() {
		total += int(digit - '0')
	}
	fmt.Println(total)
}
