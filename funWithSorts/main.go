package main

import (
	"fmt"
	"sort"
)

//type people []string
//studyGroup := people{"Zeno", "John", "Al", "Jenny"}

type people []string

func main() {

	studyGroup := people{"Zeno", "Ian", "John", "Al", "Jenny"}

	fmt.Println("1: = sorting slice of strings of type people by name:")
	fmt.Println("\to:", studyGroup)
	sort.Strings(studyGroup)
	fmt.Println("\t\t [Strings method]:", studyGroup)
	sort.Sort(sort.StringSlice(studyGroup))
	fmt.Println("\t\t [StringSlice interface conversion]", studyGroup)
	sort.Sort(sort.Reverse(sort.StringSlice(studyGroup)))
	fmt.Println("\t\t [Reverse]", studyGroup)

	fmt.Println("2: = by literal slice of strings")
	s := []string{"Zeno", "John", "Al", "Jenny", "Ben"}
	fmt.Println("\to:", s)
	sort.Strings(s)
	fmt.Println("\t\t [Strings method]:", s)
	sort.Sort(sort.StringSlice(s))
	fmt.Println("\t\t [StringSlice interface conversion]", s)
	sort.Sort(sort.Reverse(sort.StringSlice(s)))
	fmt.Println("\t\t [Reverse]", s)

	fmt.Println("3: = slice of ints")
	n := []int{7, 4, 8, 2, 9, 19, 12, 32, 3}
	fmt.Println("\to:", n)
	sort.Ints(n)
	fmt.Println("\t\t [Ints method]", n)
	sort.Sort(sort.IntSlice(n))
	fmt.Println("\t\t [IntSlice interface conversion]", n)
	sort.Sort(sort.Reverse(sort.IntSlice(n)))
	fmt.Println("\t\t [Reverse]", n)

}
