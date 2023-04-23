package main

import "fmt"

func main() {
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, number := range numbers {
		numType := "odd"   // default number type
		if number%2 == 0 { // e.g. 2 % 2 -> 0
			numType = "even" // change to even if modulus equals 0
		}
		fmt.Printf("%d is %v \n", number, numType) // e.g. 2 is even, 3 is odd
	}
}
