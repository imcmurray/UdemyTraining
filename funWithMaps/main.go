package main

import "fmt"

func main() {

	// all of the following work to create a map that we can add things to
	//	var ages = make(map[string]int)
	//ages := make(map[string]int)
	//var ages = map[string]int{}
	var ages = map[string]int{"Ian": 42, "Kristy": 42}
	//	ages["Ian"] = 42
	//	ages["Kristy"] = 41
	ages["Scott"] = 17
	ages["Toby"] = 14
	ages["Alyson"] = 12
	ages["Fiona"] = 10

	fmt.Println("Ages: ", ages)

}
