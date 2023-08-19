package main

import (
	"fmt"
	"strconv"
)

func main() {
	//int to float

	var a int = 123

	var b float64 = float64(a)

	fmt.Println(a, b)

	//float to int

	price := 200.50

	convtPrice := int64(price)

	fmt.Println(price, convtPrice) //it will round off

	//string to int

	currentPrice := "234"

	//cPriceInt:= int64(currentPrice)//cant do it like this
	cPriceInt, err := strconv.Atoi(currentPrice)

	fmt.Println(currentPrice, cPriceInt, err)

	//string to byte

	//name:= "suresh"
	name := 's'

	char := byte(name)

	fmt.Println(name, char)

}
