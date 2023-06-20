package main

import(
	"fmt"
	mypackage "com.aklil/example/exampledir"  
)

//mypackage "com.aklil/example/examplepackage"  
// mypackage is the name of the package in the file, 
// examplepackage is the folder
// com.aklil/example is the module name

func main(){
	fmt.Println("Hello", mypackage.Name)
	nums := []int{4,5,6}
	fmt.Println("StrArr :", mypackage.IntArrToStrArr(nums))
}