package main

import (
	"fmt"
)

//struct -> type collections of fields 

type Person struct{
	name string
	age  int
}
// create a new person

func newPerson(name string,age int) Person{
np:= Person{name:name,age:age}
return np
}

//return a pointer new person
func newPersonPointer(name string)*Person{
	np := Person{name:name}
	np.age = 23;
	return &np;
}

func Struct(){
	fmt.Println(newPerson("ron",20))
	fmt.Println(newPersonPointer(s));
}