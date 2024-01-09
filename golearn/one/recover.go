package main

import (
	"fmt"
)
//Go makes it possible to recover from a panic, by using the recover built-in function. A recover can stop a panic from aborting the program and let it continue with execution instead.
func myPanic(){
	panic("a problem")
}

func Recover(){
	defer func(){

		if r:= recover(); r !=nil{
			fmt.Println("Recovered. Error:\n",r)
		}
	}()
	myPanic()
	fmt.Println("After my Panic")
}