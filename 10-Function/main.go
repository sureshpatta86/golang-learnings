package main

import (
	"fmt"
)

func main() {
	greeting("suresh")
	g := greeting
	g("nani")
	multipleNamesGreetings("suresh", "aarush")
	multipleGreeting("suresh", "aarush", "nani")
	multipleOrders(1, 2, 3, 4, 5)
	functionAsAParam(greeting, "nani")
	fmt.Println(addWithReturnType(1, 2))
	fmt.Println(addWithMultiReturnType(1, 2))

	function := func() { fmt.Println("Direct assign") }
	fmt.Println("In-between")
	function()
}

func greeting(name string) {
	fmt.Println("Hello", name)
}

func multipleNamesGreetings(name1 string, name2 string) {
	fmt.Println("Hello", name1, name2)
}

func multipleGreeting(names ...string) {
	for _, value := range names {
		fmt.Println("Hello", value)
	}
}

func multipleOrders(orders ...int64) {
	for _, order := range orders {
		fmt.Println("Order Id is", order)
	}
}

func functionAsAParam(f func(string), name string) {
	f(name)
}

func addWithReturnType(a int, b int) int {
	return (a + b)
}

func addWithMultiReturnType(a int, b int) (int, bool) {
	if a < 0 || b < 0 {
		return 0, false
	} else {
		return a + b, true
	}
}
