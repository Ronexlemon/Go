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
	ff("struct1: %T\n",p) //%v type of the value
	ff("struct1: %t\n",true) //%v fomrating boolens
	
	fmt.Fprintf(os.Stderr, "io: an %s\n", "error")
}
