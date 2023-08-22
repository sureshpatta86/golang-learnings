package main

import "fmt"

func main() {

	var s = []int{1, 2, 3}
	fmt.Println(s, len(s), cap(s))

	//Append to existing one
	s = append(s, 4, 5)
	fmt.Println(s, len(s), cap(s)) //capacity doubled

	//Append to existing one
	s = append(s, 6, 7)
	fmt.Println(s, len(s), cap(s)) //capacity doubled

	s = make([]int, 5) //initiate new capacity
	fmt.Println(s, len(s), cap(s))

	s = []int{1, 2, 3, 4, 5}
	fmt.Println(s, len(s), cap(s))

	//slicing
	s1 := s[0:2] //get first two
	fmt.Println(s1, len(s1), cap(s1))

	s1 = s[0:] //get all the values from starting
	fmt.Println(s1, len(s1), cap(s1))

	s1 = s[:2] //get first two values
	fmt.Println(s1, len(s1), cap(s1))
}
