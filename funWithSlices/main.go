package main

import "fmt"

func main() {
	mySlice := []string{"a", "b", "c", "g", "m", "z"}
	fmt.Printf("%T\n", mySlice)
	fmt.Println(mySlice)
	fmt.Println(mySlice[2:4])
	fmt.Println(mySlice[2])
	fmt.Println("myString"[2])

	mySlice2 := []int{1, 3, 5, 7, 9, 11}
	for i, value := range mySlice2 {
		fmt.Println(i, " - ", value)
	}

	mySlice3 := make([]int, 0, 5)
	fmt.Println("----------")
	fmt.Println(mySlice3)
	fmt.Println(len(mySlice3))
	fmt.Println(cap(mySlice3))
	fmt.Println("----------")

	for i := 0; i <= 80; i++ {
		mySlice3 = append(mySlice3, i)
		fmt.Println("Len: ", len(mySlice3), "Cap: ", cap(mySlice3), "Val: ", mySlice3[i])
	}

	transactions := make([][]int, 0, 3) // Preferred method to make slice
	for i := 0; i < 3; i++ {
		//var transaction []int // empty nil slice - not preferred but capable
		// slice is pointing to an uninitialized array
		// (slices made without make) must be appened to.
		transaction := []int{}
		for j := 0; j < 4; j++ {
			transaction = append(transaction, j)
		}
		transactions = append(transactions, transaction)
	}
	fmt.Println(transactions)
}
