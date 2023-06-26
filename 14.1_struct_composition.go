package main

import(
	"fmt"
)

type business struct{
	name string
	//contact contact or simply
	contact
}

type contact struct{
	firstName string
	lastName string
}

func main(){
	business1 := business{
		name: "sepa",
		contact: contact{
			firstName: "marvin",
			 lastName: "balvelski",
			},
	}

	contact2 := contact{
		firstName: "abdul",
		lastName: "abdu",
	}

	business2 := business{
		name: "Sepa",
		contact: contact2,
	}

	fmt.Println("Business name :", business1.name)
	fmt.Println("Contact name  :", business1.contact.firstName)

	fmt.Println(business2)
}