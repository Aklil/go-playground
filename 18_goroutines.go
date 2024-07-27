package main

import (
	"fmt"
	"time"
)


func print(index int) {
   fmt.Println("Index: ", index)
}

func main(){

	for i:=1; i <= 15; i++ {
		go print(i) //runs in separate go routine
	}

	//pause the main routine so that the other routines could run

	time.Sleep(2 * time.Second) //sleep for 2 seconds
}
