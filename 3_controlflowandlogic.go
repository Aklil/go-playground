package main

import (
	"fmt"
)

var println = fmt.Println

func main(){
	iAge := 8
	if (iAge >= 1) && (iAge <= 18) {
		println("Important Birthday")
	}else {
		println("Not so Important")
	}
}