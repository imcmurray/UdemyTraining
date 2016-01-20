package main

import "fmt"

func main() {
	var x [58]string
	for i := 65; i <= 122; i++ {
		x[i-65] = string(i)
	}
	fmt.Println(x)
	fmt.Println(x[42])

	var y [256]int
	for i := 0; i < 256; i++ {
		y[i] = i
	}

	for i, v := range y {
		fmt.Printf("%v - %T - %b\n", v, v, v)
		if i > 50 {
			break
		}
	}

	var z [256]byte
	for i := 0; i < 256; i++ {
		z[i] = byte(i)
	}

	for i, v := range z {
		fmt.Printf("%v - %T - %b\n", v, v, v)
		if i > 50 {
			break
		}
	}

}
