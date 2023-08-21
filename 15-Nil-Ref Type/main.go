package main

import "fmt"

func main() {
	var i int

	var p *int

	if i == 0 { //default value for i is zero
		fmt.Println(i)
	}

	if p == nil { // default value for pointer is nil
		fmt.Println(p)
	}
	//nil also a default value 0 for
	//pointers
	//maps
	//slices
	//channels
	//functions
	//interfaces
}
