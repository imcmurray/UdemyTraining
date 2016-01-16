package main

import "fmt"

func main() {
	data := []float64{20, 40, 60, 80, 90}
	n := average(data...)
	//  n := average(20, 40, 60, 80, 90)

	fmt.Println("Average is: ", n)
}

func average(sf ...float64) float64 {
	fmt.Println(sf)
	fmt.Printf("%T \n", sf)
	var total float64
	for _, v := range sf {
		total += v
	}
	return total / float64(len(sf))
}
