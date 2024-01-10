package main

import (
	"fmt"
	"os"
)

type point struct{
	x,y int
}
func StringFormating() {
	p:=point{1,2}

	fmt.Printf("struct1: %v\n",p)
	fmt.Fprintf(os.Stderr, "io: an %s\n", "error")
}
