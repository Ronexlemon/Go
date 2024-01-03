package main

import (
	"fmt"
)

//supports methods on struct types
type rect struct{
	width, height int
}
type test int

func  (r *rect) area()int{

	return r.width * r.height

}

func (r rect) perim()int{
	return 2*r.width + 2*r.height
}

func (t test)  testit()int{
	return 6;
}

func Method(){
	instance:= rect{10,12}
	testa:=test(0)

	fmt.Println(testa.testit())
	fmt.Println(instance.area())
	fmt.Println(instance.perim())
}
