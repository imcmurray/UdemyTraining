package main

import "fmt"

func main() {
	defer world() // this will be the last function executed before exit
	hello()
}

func hello() {
	fmt.Print("Hello ")
}

func world() {
	fmt.Println("World")
}
