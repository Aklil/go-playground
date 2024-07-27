package main

import(
	"fmt"
)

// arrays contain one specific data type

// The range form of the for loop iterates over a slice or map.

func main() {

	// array size is not changing

	//declare
	var arr [3] int
	arr[0] = 5

	fmt.Println(arr)  //other elements are replaced by default values for eg. for int it is 0


	//declare and init
	nums := []int{6, 4, 1, 7, 8} 

	fmt.Println(nums)

	//range to loop through arrays, slices etc

	for _, num :=range nums{
		fmt.Println(num)
	}

	// multidimensional

	arr2 := [][]int {
		{3, 2},
		{5, 4},
	}

	fmt.Println(arr2)


	// string to rune

	strArr := "abcde"
	runeArr := []rune(strArr)  //cast to rune array
	for _, v := range runeArr {
		fmt.Printf("%c", v)
	}

	// byte array
	// byte represents- ASCII chars 
	// rune represents - Unicode UTF-8 chars

	byteArr := []byte{'a', 'b', 'c'}
	byteStr := string(byteArr)  // byte array to string
	//byteStr := string(byteArr[:])  // byte array to string with range; [:] all value or eg. [1:2] 

	fmt.Println(byteStr)

	// SLICES
	// Aslice is a reference to a contiguous segment of an array
	// Unlike an array, which is a value-type, slice is a reference type.
	// A slice can be a complete array or a part of an array, indicated by the start and end index.
	//A slice is a dynamic-sized data structure.
	// 	This means that the size of a slice can be changed after it is created. You can add or remove elements from a slice, and the slice will automatically resize itself to accommodate the changes.
	// A slice is backed by an array underneath the implementation in go

	sl1 := make([]string, 3) //The make built-in function allocates and initializes an object of type slice, map, chan
    //gave it initial size of 3
	sl1[0] = "initial"
	sl1[1] = "size"
	sl1[2] = "three"

	//now resize by appending
	sl1 = append(sl1,"Coding")
	sl1 = append(sl1, "Is")
	sl1 = append(sl1, "cool")
	sl1 = append(sl1, "right?")

	fmt.Println("Slice size :", len(sl1))


	sl2 := nums[0:2] //slice up to doesn't include the value at index

	fmt.Println(sl2)

	fmt.Println("1st 3 :", nums[:3])
	fmt.Println("last 3 :", nums[2:])  // slice from includes the value at the index

	// **** slice points to the underlining array
	// changing the array changes the value of the slice 
	// and changing the value of the slice changes the underlining array

	sl2 = append(sl2, 99)  //99 replaces the value 1 in the underlying array - nums
	sl2 = append(sl2, 88) //88 replaces 7 in the underlying array - nums

	fmt.Println(sl2)

	fmt.Println("nums :", nums)  






}