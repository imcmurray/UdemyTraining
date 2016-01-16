package main

import "fmt"

func visit(numbers []int, catdog func(int)) {
	for _, n := range numbers {
		catdog(n)
		//catdog is the callback to the function that was passed in.
	}
}

func main() {
	visit([]int{1, 2, 3, 4, 5, 6}, func(n int) {
		fmt.Println(n)
	})

	xs := filter([]int{1, 2, 3, 4, 5, 6}, func(n int) bool {
		return n > 1
	})
	fmt.Println(xs)
}

func filter(numbers []int, callback func(int) bool) []int {
	var xs []int
	for _, n := range numbers {
		if callback(n) {
			xs = append(xs, n)
		}
	}
	fmt.Printf("%T \n", callback)
	return xs
}
