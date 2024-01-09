package main

import (
	"fmt"
	"cmp"
	"slices"
)

func SortByFunction(){
	fruits:=[]string{"peach","banana","kiwi"}

	lenCmp:= func(a,b string)int{
		return cmp.Compare(len(a),len(b))
	}
	slices.SortFunc(fruits,lenCmp)
	fmt.Println(fruits)
	
	//for built in types
	type Persons struct{
		name string
		age int
	} 
	persn:=[]Persons{{"rone",10},{"ron",5},{"ronex",3}}
	slices.SortFunc(persn,func(a,b Persons)int{
		return cmp.Compare(a.age,b.age)
	})
	fmt.Println(persn)
}