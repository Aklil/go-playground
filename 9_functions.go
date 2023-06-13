package main

import(
	"fmt"
)

var println = fmt.Println

func getSum(x int, y int) int {
	return x + y
}

// multi val return

func getTwo(x int) (int, int){
	return x+1, x+2
}

// errors

func divide(x float64, y float64) (float64, error) {
	
	if y==0 {
		return 0, fmt.Errorf("Divide by Zero")
	}

	return x/y, nil
}

// varaidic function, a func that receives unknown number of values

func getDynamicSum(nums ...int) int {
	sum := 0

	for _, num := range nums {
		sum += num
	}

	return sum
}

// pass array by value, for an example sake here

func getArrSum(arr []int) int {
	sum := 0

	for _, val := range arr {
		sum += val
	}

	return sum
}


// pass the pointer
func changeVal(ptrVal *int) {
	*ptrVal = 10
}




func main() {
	println(getSum(5, 4))

	println(getTwo(5))

	println(divide(6.5, 0))

	println(divide(6.5, 2))

	println(getDynamicSum(4,5,6))

	println(getArrSum([]int{6, 7, 3}))

	// pass by reference

	num := 100

	println("before :", num)

	changeVal(&num) //pass by reference

	println("after :", num)

	var numPtr *int = &num
	//var numPtr = &num //go infers data type

	println("numPtr value is the memory address of num at :", numPtr) //the address
	println("value pointed by numPtr :", *numPtr)  // get value by pointing to the address
}