package main

import (
	"fmt"
	"reflect"
	"strconv"
)

var println = fmt.Println

func main() {
	//types
	println("Exploring Go types")
	println(reflect.TypeOf(40))  //int
	println(reflect.TypeOf(3.14)) //float64
	println(reflect.TypeOf(true)) //bool
	println(reflect.TypeOf("Yollo!")) //string	

	//casting
	v1 := 1.5
	v2 := int(v1)
	println(v2)

	v3 := "50000000"
	v4, err := strconv.Atoi(v3)  // Ascii to int
	
	v31 := 50000000
	v41 := strconv.Itoa(v31) // Integer to Ascii string
	println(v41, err, reflect.TypeOf(v41))

	println(v4, err, reflect.TypeOf(v4))

	v5 := strconv.Quote("Hello, 世界")  //to double quoted string
	
	println(v5)

	//string to float

	v6 := "3.14"
	
	//one line err handling
	if v7, err := strconv.ParseFloat(v6, 64); err == nil {   //64 is for bitsize i.e float64
		println(v7)
	}

}