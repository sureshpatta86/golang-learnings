package main

import "fmt"

func main() {

	var i int = 1
	var iPointer *int = &i

	//i value and & will give address of the i
	//and * represents the pointer which used to point the address

	fmt.Println(i, &i, iPointer)

	//how to fetch the address and value, using pointers
	fmt.Println(*iPointer) // it will defrence the pointer means get the address value

	a := 10
	fmt.Println(a)
	update(a)
	fmt.Println(a)
	updateAddressUsingPointer(&a)
	fmt.Println(a)
}

func update(a int) {
	a = 20
}

func updateAddressUsingPointer(aPointer *int) {
	*aPointer = 20
}
