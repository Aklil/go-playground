package main

import (
	"fmt"
	"strings"
)

var println = fmt.Println

func main() {
	str1 := "Hello"
	println("length:", len(str1))
	println("Contains another : ", strings.Contains(str1, "llo"))
	println("0 index:", strings.Index(str1, "l"))

	println("replace :", strings.Replace(str1, "l", "1", -1))  //replaces "l" with "1". -1 tells to search the whole string

	
}