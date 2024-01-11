package main

import (
	"fmt"
	"os"
)

var ff = fmt.Printf
type point struct{
	x,y int
}
func StringFormating() {
	p:=point{1,2}

	ff("struct1: %v\n",p) //%v prints the struct values
	ff("struct1: %+v\n",p) //+%includes struct values
	ff("struct1: %#v\n",p) //%v prints the syNTAX
	ff("type: %T\n",p) //%v type of the value
	ff("bool: %t\n",true) //%v fomrating boolens
	ff("int: %d\n",123)
	ff("bin %b\n",14)//binary representation
	
	fmt.Fprintf(os.Stderr, "io: an %s\n", "error")
}
