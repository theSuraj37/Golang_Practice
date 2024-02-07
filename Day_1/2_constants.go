package main

import "fmt"

func main() {

	const a = 10
	const str = "suraj"
	const x, y, z = 11, 21.51, "demo"

	// there will be no error if we are only declaring it and not using it.

	fmt.Printf("%T%T%T", x, y, z)
}
