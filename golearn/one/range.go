package main

import (
	"fmt"
)
/*
range iterates over elements in a variety of data structures.
*/

func returnSlice()[]int{
	
	s:=[]int{1,2,3,4,5,6,7,8,9}
	return s
}

func returnMap()map[string]int{
	m:=make(map[string]int)
	m["big"] =10
	m["medium"]=5
	m["small"] =1
	m["bigInit"] =20
	m["mediumiit"]=15
	m["smallinit"] =11
	
	return m
}
func sumAllNumbersInASlice(){
	nums:= returnSlice()
	sum:=0

	for _,num:= range nums{ // ignore the first index
		sum +=num
	}
	fmt.Println("the sum is",sum)
}
//range on arrays and slices provides both the index and value for each entry

func usingTheIndex(){
	s:=returnSlice()
	for i,num:=range s{
		if num == 5{
			fmt.Println("index",i)
		}
	}
}

//range on maps iterates over key/values
func iterateOverMap(){
	m:= returnMap()

	for k,v:= range m{
		fmt.Printf(k,v)
	}
}
//iterate just over only the keys
func iterateOverMapKeys(){
	m:= returnMap()

	for k:= range m{
		fmt.Printf("keys: \n ",k)
	}
}
//range on strings iterates over Unicode code points
func iterateString(){
	for i, c:= range "ronex"{
		fmt.Println(i,c)
	}
}
func Range(){
	sumAllNumbersInASlice()
	usingTheIndex()
	iterateOverMap()
	iterateOverMapKeys()
	iterateString()

}