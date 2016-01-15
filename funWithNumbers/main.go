package main

import "fmt"

func main() {
	fmt.Println("Numeral: ", 42)
	fmt.Printf("with Binary: %d - %b\n", 42, 42)
	fmt.Printf("with Hex: %d - %b - %x\n", 42, 42, 42)

	fmt.Printf("with Hex type (lower) & UPPER: %d - %b - %#x - %#X\n", 42, 42, 42, 42)

	fmt.Println("Decimal, Binary, Hex, Char, Base8, Unicode")
	for i := 60; i <= 70; i++ {
		fmt.Printf("%d \t %b \t %x \t %q \t %o \t %U \n", i, i, i, i, i, i)

	}

}
