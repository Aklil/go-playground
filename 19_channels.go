package main


import(
	"fmt"
)

//pass values from go routine to another go routine via channel

func generateInts(channel chan int){
	channel <- 1
	channel <- 2
	channel <- 3

	close(channel)  // close channel for 'range' use, unless 'deadlock' is thrown
}

func generateRunes(channel chan rune) {
	channel <- 'A'
	channel <- 'B'
	channel <- 'C'

	close(channel)
}



func main() {
   channel1 := make(chan int)
   channel2 := make(chan rune)

   go generateInts(channel1)
   go generateRunes(channel2)


//    fmt.Println("channel1 val: ", <-channel1)
   //iterate channel values using range

   for element := range channel1 {
	fmt.Println("channel 1 element: ", element)
   }

   for element2 := range channel2 {
	fmt.Println("channel 2 element: ", element2)
   }
}