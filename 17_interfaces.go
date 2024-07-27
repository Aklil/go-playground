package main

import(
	"fmt"
)

type Vehicle interface {
	move()
}

type Car string

func (c Car) move(){   //make car work with Vehicle interface
    fmt.Println("car is driving")
}

func (c Car) Name() string {  //car can have its own associative method
	return string(c) //cast from car to string, car is defined as string after all
}



func main() {
   var car Vehicle
   car = Car("Toyota")
   car.move()

   car2 := car.(Car)  //cast from Vehicle to Car

   fmt.Println("car name", car2.Name())
}