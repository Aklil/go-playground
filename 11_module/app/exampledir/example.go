//funcs and variables that start in Upper case are accessbile in another
package mypackage

import(
	// "errors"
	// "fmt"
	"strconv"
	// "tme"
)

var Name = "John"

func IntArrToStrArr(intArr []int) []string{
	var strArr []string

	for _, val := range intArr {
		strArr = append(strArr, strconv.Itoa(val))
	}

	return strArr
}

