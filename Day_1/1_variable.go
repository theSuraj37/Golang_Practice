package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a int = 10
	var b int = 20
	c := 30

	fmt.Println(a, b, c)
	fmt.Printf("%d", a)
	fmt.Printf("%T", a)            // int   // using format specifier %T we get the datatype of variable
	fmt.Println(reflect.TypeOf(a)) // int	// using reflect.TypeOf() we get the datatype of variable

	myString := `To write multiline string
	we can use this operator`

	fmt.Println(myString)

	x, y, z := 11, 21.51, "suraj"

	fmt.Println(reflect.TypeOf(x), reflect.TypeOf(y), reflect.TypeOf(z))
	fmt.Println(x, y, z)
}
