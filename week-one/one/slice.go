package main

import (
	"fmt"
)

func unInitialisedSlice(){
	var s []string
	
	fmt.Println("uninit",s,s ==nil,len(s)==0)
}
func emptysliceWithnonZero(){
	s:=make([]string,3)
	fmt.Println("emp",s,"len",len(s),"cap",cap(s))

}