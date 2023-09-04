package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	fmt.Println(math.Pi)                 //returns the closest floating-point value of Pi.
	fmt.Println(math.Min(123, 33))       //returns the smaller of x or y.
	fmt.Println(math.Max(434, 2312))     //returns the larger of x or y.
	fmt.Println(math.MaxInt)             //returns the maximum value of type int.
	fmt.Println(math.MaxFloat64)         //returns the maximum finite floating-point value of type float64.
	fmt.Println(math.Ceil(123))          //returns the smallest integer value greater than or equal to x.
	fmt.Println(math.Floor(333))         //returns the greatest integer value less than or equal to x.
	fmt.Println(math.Round(123.5))       //returns the nearest integer, rounding half away from zero.
	fmt.Println(math.RoundToEven(123.5)) //returns the nearest integer, rounding ties to even.
	fmt.Println(math.Sqrt(123))          //returns the square root of x.
	fmt.Println(math.Pow(2, 3))          //returns x**y, the base-x exponential of y.
	fmt.Println(math.Log(123))           //returns the natural logarithm of x.

	fmt.Println(rand.Int())     //returns a non-negative pseudo-random int.
	fmt.Println(rand.Intn(100)) //returns a non-negative pseudo-random int in [0,n).
	fmt.Println(rand.Float64()) //returns a pseudo-random float64 in [0.0,1.0).
	fmt.Println(rand.Int())
	fmt.Println(rand.Int()) //From 1.20 version, the seed is not required to be set.
}
