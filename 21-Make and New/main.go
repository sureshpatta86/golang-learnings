package main

import (
	"fmt"
)

func main() {
	// make
	a := make([]int, 5)
	fmt.Println(a)

	// new
	b := new([]int)
	fmt.Println(b)
	fmt.Println(*b)

	// Make is used for slices, maps, and channels only
	// New is used for types that require dynamic memory allocation like struct, int, etc.

	// make
	c := make(map[string]int)
	fmt.Println(c)

	// new
	d := new(map[string]int)
	fmt.Println(d)
	fmt.Println(*d)

	// make
	e := make(chan int)
	fmt.Println(e)

	// new
	f := new(chan int)
	fmt.Println(f)
	fmt.Println(*f)

	// return types of make and new

	// make
	g := make([]int, 5)
	fmt.Printf("%T\n", g)

	// new
	h := new([]int)
	fmt.Printf("%T\n", h)

	var i *[]int = new([]int)
	fmt.Printf("%T\n", i)

	//Memory Initialization
	// make
	j := make([]int, 5, 10)
	fmt.Println(j)

	// new
	k := new([]int)
	fmt.Println(k)

	j[0] = 1
	fmt.Println(j)

	var new *map[int]int = new(map[int]int)
	(*new)[0] = 1    //compile fine but will panic at runtime
	fmt.Println(new) // panic: assignment to entry in nil map
	//need to consider the warning yellow line too
}
