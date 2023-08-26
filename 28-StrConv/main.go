package main

import (
	"fmt"
	"strconv"
)

func main() {

	//String to int
	str := "123"
	i, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(i)
	str = "a123"
	i, err = strconv.Atoi(str)
	if err != nil {
		fmt.Println(err) //strconv.Atoi: parsing "a123": invalid syntax
	}
	fmt.Println(i) //gives o

	//int to string

	a := 123
	str = strconv.Itoa(a)
	fmt.Println(str)
	b := 123.45
	str = strconv.Itoa(int(b))
	fmt.Println(str)

	//string to float
	str = "123.45"
	f, err := strconv.ParseFloat(str, 64)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(f)

	//Bool to string
	b1 := true
	str = strconv.FormatBool(b1)
	fmt.Println(str)

	//String to bool
	str = "true"
	b1, err = strconv.ParseBool(str)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(b1)

	str = "Truee"
	b1, err = strconv.ParseBool(str)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(b1)

	//string to int64
	str = "123"
	i64, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(i64)

	//string literals

	str = `Hello
	World`
	fmt.Println(str)

	str = "Hello\nWorld"
	fmt.Println(str)

	fmt.Println(strconv.Quote("Hello \"World\"")) // "Hello \"World\""

}
