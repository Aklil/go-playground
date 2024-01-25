package main

import (
	"fmt"
	"reflect"
	mypackage "aklil.com/example/exampledir" //referring a module
)

//mypackage "com.aklil/example/exampledir"
// mypackage is the name of the package in the file,
// exampledir is the folder
// com.aklil/example is the module name

func main(){
	fmt.Println("Hello", mypackage.Name)
	nums := []int{4,5,6}
	fmt.Println("StrArr :", mypackage.IntArrToStrArr(nums))
	//check type using reflect
	fmt.Println("StrArr type check :", reflect.TypeOf(mypackage.IntArrToStrArr(nums)))
}