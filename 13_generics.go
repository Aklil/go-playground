package main

import (
	"fmt"

)

// there are built in constraints
// for more at: https://pkg.go.dev/golang.org/x/exp/constraints
func sum[T complex128](a T, b T) T{
	return a + b
}

// user defined constraint

type MyConstraint interface{
	int | float64
}

func twoSum[T MyConstraint](a T, b T) T{
	return a + b
}

func main(){
	fmt.Println("sum 5.6, 7.8 :", twoSum(5.6, 7.8))
}