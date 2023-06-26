package main

import(
	"fmt"
)

type rectangle struct{
	length, width float64
	perimeter func() float64 //regarded as just a field, IT IS NOT a method set of the struct
}

//struct function
func (r rectangle) getArea() float64{
	return r.length * r.width
}

//pointer is used inorder to change the value of the struct
func (r *rectangle) scale(scale float32){
	r.length *= float64(scale)
	r.width *= float64(scale)
}

func main(){

	rect:= rectangle{
		length: 5.6,
		width: 8.99,
		//perimeter: func(){return 2*rect.length + 2*rect.width},
	}
	fmt.Println("rect :", rect)

	fmt.Println("area :", rect.getArea())


	rect.perimeter = func() float64{ //assigning just a field of func type
		return 2*rect.length + 2*rect.width
	}

	fmt.Println("perimeter:", rect.perimeter())

	scale := 0.5
	rect.scale(float32(scale))

	fmt.Printf("Scale :%f , Area : %f", scale, rect.getArea() )
}