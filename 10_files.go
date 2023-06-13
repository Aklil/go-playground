package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

var println = fmt.Println


func main() {
	f, err := os.Create("data.txt")

	if err != nil {
		log.Fatal(err)
	}

	//close file, when main func returns

	defer f.Close()
	// A defer statement defers the execution of a function until the surrounding function returns.
	// The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.

	intArr := []int {2,3,5,7,11}
	var strArr []string

	for _, val := range intArr {
		strArr = append(strArr, strconv.Itoa(val)) //Itoa - Integer to string
	}

	for _, num := range strArr {
		_, err := f.WriteString(num + "\n")

		if err != nil {
			log.Fatal(err)
		}
	}

	f, err = os.Open("data.txt")

	if err != nil {
		log.Fatal(err)
	}

	scan := bufio.NewScanner(f)
	
	for scan.Scan(){ //scans line by line
		println("Number: ", scan.Text()) //reads the line
	}


	//demo, initializing and using a variable
	if err := scan.Err(); err != nil{
		log.Fatal(err)
	}


}