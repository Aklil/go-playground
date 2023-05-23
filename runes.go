package main

import (
	"fmt"
	"unicode/utf8"
)

var println = fmt.Println

func main(){

	str := "Runes are chars in another language"

	// rune count
	println("rune count :", utf8.RuneCountInString(str))

	for i, runeVal := range str {
		println("i :", i, runeVal)
		//formatted printing
		fmt.Printf("\n%d : %#U : %c \n", i, runeVal, runeVal)  // %#U is a unicode representation of the character
	}
}
