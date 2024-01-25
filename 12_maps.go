package main

import(
	"fmt"
)

func main(){
	heroes := make(map [string]string)  //[key]value
	heroes["Batman"] = "Bruce Wayne"
	heroes["Superman"] = "Clark Kent"

	// or inline
	villains := map [string] string {"Lex Luther": "Lex Luther"}

	//access
	fmt.Println("1st element :", villains["Lex"]) //nothing will be printed , i.e nill character

	//check exists
	_, ok := heroes["Ironman"]   
	fmt.Println("Ironman exists :", ok) //boolean

	//loop 
	for k,v := range heroes{
		fmt.Println("hero", k, v)
	}

	//delete entry
	delete(heroes, "Batman")
	_, exists := heroes["Batman"]
	fmt.Println("is deleted :", !exists)
}