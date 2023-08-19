package main

import (
	"fmt"
)

func main() {

	fmt.Println(true, false, true || false, true && false) //Logical

	fmt.Println(100+100, 100-100, 100*100, 100/100, 100%100) //Mathematical

	fmt.Println(1 == 1, 1 != 1, 1 > 1, 1 < 1, 1 <= 1, 1 >= 1) //Relational

	//Assignment
	i := 0
	fmt.Println(i)
	i = 1
	fmt.Println(i)
	i += 1
	fmt.Println(i)
	i -= 1
	fmt.Println(i)
	i *= 1
	fmt.Println(i)
	i /= 1
	fmt.Println(i)
	i %= 1
	fmt.Println(i)
	i++
	fmt.Println(i)
	i--
	fmt.Println(i)

	//Bitwise
	j := 1
	fmt.Println(j)

	j = j << 1
	fmt.Println(j)

	j = j >> 1
	fmt.Println(j)

	j += j << 1
	fmt.Println(j)

	j += j >> 1
	fmt.Println(j)
}
