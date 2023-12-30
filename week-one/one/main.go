package main

import (
	"fmt"
)

func main() {
	fmt.Println("go" + "lang") // strings
	fmt.Println(true && false) // boolean
	fmt.Println(7.0 + 7.33)    // float
	fmt.Println(5 + 5)         // integer

	// Call the Variable function from the same package
	 Variable()
     Constants()
     
	ForLoop()
	IfElse()
}
