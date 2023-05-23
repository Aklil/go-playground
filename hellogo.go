package main // this is the main package

import (
	"bufio"
	"fmt" // format:
	"log"
	"os" // common abstraction for os specifics
)

var println = fmt.Println // shortcut for the funtion

// the main function for start of execution

func main(){
	println("What is your name?")
	reader := bufio.NewReader(os.Stdin)  //buffered io reader
	name, err := reader.ReadString('\n') // read until new line or enter key

	//only funcs with Upper case initials can be exported and used from another package

	// err stores any error that has occurred

	//handling error
	if err == nil {
		println("hello", name)
	} else {
		log.Fatal(err)
	}



}

