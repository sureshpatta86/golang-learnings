package main

import "fmt"

func main() {

	var arr = [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr)
	fmt.Println(arr[3])

	arr1 := []int{1, 2, 3}
	fmt.Println(arr1, len(arr1))

	if len(arr) > 0 {

		for index, value := range arr1 {
			fmt.Println(index, "-", value)
		}
	}

	arr3 := [5]int{10}
	fmt.Println(arr3[1])

	arr4:= [3]string{"suresh","nani"}

	for _,value := range arr4{
		fmt.Println(value)
	}

	arr4[1]="abcd"
	fmt.Println(arr4[1])

}
