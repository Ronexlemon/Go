package main

import (
	"fmt")

func simpleArraya(){
	var a[5]int
	fmt.Println("em",a)
	a[3]=100
	fmt.Println("new a 3",a)


	b:=[5]int{1,2,3,4,5}
	fmt.Println("dcl",b)
	b[3]=500
	fmt.Println("new b",b)
	fmt.Println("new b length",len(b))


}

func twoDArray(){
	var twoD[2][3]int
	b:=[2][3]int{{1,2,3},{1,2,3}}
	for i:=0;i <2; i++{
		for j:=0;j <3; j++{
			twoD[i][j]= i+j
		
		}
	}
	fmt.Println("twod array",twoD)
	fmt.Println("two interal array",b)
}

func Array(){
	simpleArraya()
	twoDArray()
}
