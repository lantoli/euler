package main

import (
	"fmt" 
	"strconv"
	)
		
func IsPalindrome(i int) bool {
	s := strconv.Itoa(i)
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s) - 1 - i] {
			return false
		} 
	}
	return true
}

func main() {
	ret := 0
	for i := 0; i <= 999; i++ {
		for j := 0; j<= 999; j++ {
			if num := i*j; IsPalindrome(num) && num > ret {
				ret = num
			}
		}
	}
 	fmt.Println(ret)
}
