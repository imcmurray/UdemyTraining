// Write a function with one variadic parameter that finds the greatest number in a list of numbers.
package main

import "fmt"

func main() {
	numbers := []int{1, 43, 12, 48, 69}
	largestNumber := findBig(numbers...)
	fmt.Println("Largest number is", largestNumber)
}

// findBig func can take one or an unlimited number of arguments passed in
// to the numbers slice

func findBig(ns ...int) int {
	var biggest int
	for _, number := range ns {
		if number > biggest {
			biggest = number
		}
	}
	return biggest

}
