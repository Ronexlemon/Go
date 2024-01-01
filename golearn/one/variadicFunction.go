package main

import (
	"fmt"
)
/*
Variadic functions can be called with any number of trailing arguments. For example, fmt.Println is a common variadic function.
*/

func sum(nums ... int){
	total:=0
	for _,num:= range nums{
		total +=num
	}
	fmt.Println("the sum of nums:",nums,total)
}

func Variadic(){
	s:=[]int{1,2,3,4,5}
	sum(8,9)
	sum(s...)
}