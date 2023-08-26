package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Hello World", "Hello"))
	fmt.Println(strings.Compare("Hello", "Hello"))        //return 0 if equal
	fmt.Println(strings.Compare("Hello", "Helloo"))       //return -1 if not equal
	fmt.Println(strings.Compare("Hello", "Hell"))         //return 1 if not equal
	fmt.Println(strings.ContainsAny("Hello World", "H"))  //return true if any char is present
	fmt.Println(strings.ContainsAny("Hello World", "Hh")) //return true if any char is present
	fmt.Println(strings.Count("Hello World", "l"))        //return 3
	fmt.Println(strings.Index("Hello World", "l"))        //return 2

	//String Builder
	sb := strings.Builder{}
	sb.WriteString("Hello")
	sb.WriteString(" World")
	fmt.Println(sb.String())
	fmt.Println(sb.Len())
	fmt.Println(sb.Cap())

	sb.Reset()
	fmt.Println(sb.String())
	fmt.Println(sb.Len())
	fmt.Println(sb.Cap())

	sb.Grow(10)
	fmt.Println(sb.String())
	fmt.Println(sb.Len())
	fmt.Println(sb.Cap())

	sb.WriteString("Hello")

	fmt.Println(sb.String())
	fmt.Println(sb.Len())
	fmt.Println(sb.Cap())

}
