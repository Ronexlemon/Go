package main

import (
	"fmt"
	"time"
)

//basic switch condition

func basicSwitch(){
	i:=2
	fmt.Print("Write",i,"as")
	switch i{
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}
}

//use commas to separate multiple expressions in the same case statement

func multipleCase(){
	switch time.Now().Weekday(){
	case time.Saturday,time.Sunday:
		fmt.Println("Its the Weekend")
	default:
		fmt.Println("Its a weekday")
	}

}

//switch without an expression as an alternate way to express if/else

func withoutAnExpression(){
	t:= time.Now()
	switch{
	case t.Hour() < 12:
		fmt.Println("Its before noon")
	default:
		fmt.Println("its After noon")
	}

}
//compare types instead of values

func WhatAmI(i interface{}){
	switch i.(type){
	case bool:
		fmt.Println("boolean")
	case int:
		fmt.Println("integer")
	default:
		fmt.Println("don't know type")

	}
}

func SwitchCase(){
	basicSwitch()
	multipleCase()
	withoutAnExpression()
	WhatAmI(true)
	


}