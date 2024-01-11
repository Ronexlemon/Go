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
	ff("char: %c\n",33) //character representation of a number
	ff("hex: %x\n",433) //hex encoding
	ff("float1: %f\n",70.6) //floating point
	ff("float2: %e\n",124000000.0) //scientific notation
	ff("float2: %E\n",124000000.0)
	
	fmt.Fprintf(os.Stderr, "io: an %s\n", "error")
}
