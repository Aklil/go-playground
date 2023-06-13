package main

import (
	"fmt"
)

var println = fmt.Println

func main(){

	for i:=1; i<5; i++ {
		println("i :", i)
	}

	// or initialize is outside

	j := 1

	for j<5 {
		println("j :", j)
		j++
	}

	// The range form of the for loop iterates over a slice or map.

	
}