package main

import (
	"fmt"

)

// there are built in constraints
// for more at: https://pkg.go.dev/golang.org/x/exp/constraints
func getSum[T complex128](a T, b T) T{
	return a + b
}

// user defined constraint

type MyContraint interface{
	int | float64
}

func getTwoSum[T MyContraint](a T, b T) T{
	return a + b
}

func main(){
	fmt.Println("sum 5.6, 7.8 :", getTwoSum(5.6, 7.8))
}