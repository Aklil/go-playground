package main

import(
	"fmt"
)

type KM float64 //user defined types
type Mile float64

//add associate methods to the userdefined types
func (km KM) toMile() Mile{
	return Mile(km * 0.621371) //casting
}

func (ml Mile) toKM() KM{
	return KM(ml * 1.60934)
}

func main() {
	km := KM(4.5)
	fmt.Printf("%.2f kms is %.2f miles", km , km.toMile())  //kind of like an extension method

	ml := Mile(6.78)
	fmt.Println()
	fmt.Printf("%.2f Miles is %.2f KMs", ml, ml.toKM())
}