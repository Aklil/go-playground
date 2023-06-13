package main

import (
	"fmt"
	"math/rand"
	"time"
)

var println = fmt.Println

func main(){
	now := time.Now()

	println(now.Year(), now.Month(), now.Day())
	println(now.Hour(), now.Minute(), now.Second())

	//math
	seedSecs := time.Now().Unix()
	rand.Seed(seedSecs)
	randNum := rand.Intn(50) + 1

	println("Random :", randNum)
	
	

}