package main

import "fmt"

func main() {
	const n = 20
	for x := n; ; x++ {
		good := true
		for div := 2; div <= n; div++ {
			if x % div != 0 {
				good = false
			}
		}
		if good {
		 	fmt.Println(x)
		 	return
		}
	}

}
