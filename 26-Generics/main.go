package main

import (
	"fmt"
)

func main() {
	fmt.Println(miniumValueInt(2, 3))
	fmt.Println(miniumValueFloat(2.2, 3.3))

	//Type Parameter with interface{}
	miniumValue(2, 3)     //For Int
	miniumValue(2.2, 3.3) //For Float

	//Type Parameter with any
	miniumValueAny(2, 3)     //For Int
	miniumValueAny(2.2, 3.3) //For Float

	//Type Parameter with constraint
	miniumValueWithConstraint(1, 2)     //For Int
	miniumValueWithConstraint(2.2, 3.3) //For Float
	//miniumValueWithConstraint("a", "b") //For String (Error) because of constraint

	//Type interface
	var newFloat float64 = 2.5
	miniumValueWithConstraint(newFloat, 3.3)

	//Type Parameter with constraint only float
	miniumValueWithConstraintOnlyFloat(2.2, 3.3)
	//miniumValueWithConstraintOnlyFloat(2, 3) //Error because of constraint
	miniumValueWithConstraintOnlyFloat(newFloat, 3.3)

	//Type set
	fmt.Println(minValueUsingTypes(2, 3))
	fmt.Println(minValueUsingTypes(2.2, 3.3))
	//minValueUsingTypes("a", "b") //Error because of constraint

	//Generics New Deceleration
	fmt.Println(minValueNew(2, 3))
	fmt.Println(minValueNew(2.2, 3.3))
}

func minValueNew[T miniumValueTypes](a, b T) T {
	if a < b {
		return a
	}
	return b
}

type miniumValueTypes interface {
	int | float64
}

func minValueUsingTypes[T miniumValueTypes](a T, b T) T {
	if a < b {
		return a
	}
	return b
}

func miniumValueWithConstraintOnlyFloat[T ~float64](a T, b T) T {

	return a
}

func miniumValueWithConstraint[T float64](a T, b T) T {
	fmt.Println(a, b)
	return a
}

func miniumValue[T interface{}](a T, b T) T {
	fmt.Println(a, b)
	return a
}

func miniumValueAny[T any](a T, b T) T {
	fmt.Println(a, b)
	return a
}

func miniumValueInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func miniumValueFloat(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}
