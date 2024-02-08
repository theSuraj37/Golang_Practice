package main

import (
	"fmt"
	"time"
)

func checkPrime(a int) bool {
	defer func(t time.Time) {
		fmt.Println("It took", time.Since(t))
	}(time.Now())
	i := 0
	for i = 2; i <= a/2; i++ {
		if a%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println("Enter Number: ")
	input := 0
	fmt.Scan(&input)
	ans := checkPrime(input)
	if ans {
		fmt.Println("Its a Prime Number.")
	} else {
		fmt.Println("Its not a Prime Number")
	}
}
