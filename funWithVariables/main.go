package main

import "fmt"

var g string = "Buffy"

func main() {

	a := 10
	b := "golang"
	c := 4.17
	d := true

	fmt.Printf("%T %v \n", a, a)
	fmt.Printf("%T %v \n", b, b)
	fmt.Printf("%T %v \n", c, c)
	fmt.Printf("%T %v \n", d, d)

}
