package main

import (
	"fmt"
	"log"

	dateencapsulator "aklil.com/dateutils/datefuncs" //referring a module
)

func main(){

	 date := dateencapsulator.Date{}  // instantiate
   
	err := date.SetDay(1)

	if err != nil {
		log.Fatal(err)  //displays error and exits
	} 


	err = date.SetMonth(6)

	if err != nil {
		log.Fatal(err)
	}

	err = date.SetYear(1990)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Date %d/%d/%d", date.GetDay(), date.GetMonth(), date.GetYear()) 
}