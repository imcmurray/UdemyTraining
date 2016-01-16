// Write a function which takes an integer. The function will have two returns. The first return should be the argument divided by 2. The second return should be a bool that let’s us know whether or not the argument was even. For example:
// half(1) returns (0, false)
// half(2) returns (1, true)

package main

import "fmt"

func main() {
	var number int
	//	var evenValue bool
	fmt.Print("Please enter your number: ")
	fmt.Scan(&number)
	halfValue, isEven := check(number)
	fmt.Println("returns: ", halfValue, " and is even: ", isEven)

}

func check(n int) (float64, bool) {
	fmt.Println("Got to check with: ", n)
	return float64(n) / 2, n%2 == 0
}
