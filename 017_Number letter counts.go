package main

import "fmt"

func main() {

	num_0_9 := []string{
		"", "one", "two", "three", "four",
		"five", "six", "seven", "eight", "nine"}

	num_0_9_tenth := []string{
		"", "ten", "twenty", "thirty", "forty",
		"fifty", "sixty", "seventy", "eighty", "ninety"}

	num_11_19 := []string{
		"eleven", "twelve", "thirteen", "fourteen", "fifteen",
		"sixteen", "seventeen", "eighteen", "nineteen"}

	total := len("onethousand")

	for i1, num1 := range num_0_9 {
		if i1 > 0 {
			num1 += "hundred"
		}
		for i2, num2 := range num_0_9_tenth {
			for i3, num3 := range num_0_9 {
				if i1 > 0 && i2 == 0 && i3 == 1 {
					num1 += "and"
				}
				if i2 == 1 && i3 > 0 {
					num2 = num_11_19[i3-1]
					num3 = ""
				}
				total += len(num1) + len(num2) + len(num3)
			}
		}
	}

	fmt.Println(total)
}
