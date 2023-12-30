package main

import (
	"fmt"
	
)

const initialValue = 2

func loopAndSum() int {
	sum := 0
	for i := 0; i < initialValue; i++ {
		sum += i
	}
	return sum
}
func pattern(){
	for i :=0 ; i< 10 ; i++{
		for j :=0 ; j <i ; j++{
			fmt.Printf("*")
		}

	}
}

func singleLoop(){
	i:= 2
	for i <=10 {
		fmt.Println(i)
		i +=1

	}
}

func noConditionLoop(){
	for{
		fmt.Println("loop")
		break
	}
	
}

func continueToNextIteration(){
	
	for i:=0 ; i<=10;i++{
 if i%3 ==0{
	continue
 }
 fmt.Println(i)
	}
}

func ForLoop(){
fmt.Println(loopAndSum())
singleLoop()
pattern()
continueToNextIteration()
noConditionLoop()



}