package main

import (
	"fmt"
	
)

const initialValue = 2

func LoopAndSum() int {
	sum := 0
	for i := 0; i < initialValue; i++ {
		sum += i
	}
	return sum
}
func Pattern(){
	for i :=0 ; i< 10 ; i++{
		for j :=0 ; j <i ; j++{
			fmt.Printf("*")
		}

	}
}
